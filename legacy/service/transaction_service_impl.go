package service

import (
	"strconv"
	"time"

	"github.com/it-chain/it-chain-Engine/common"
	"github.com/it-chain/it-chain-Engine/legacy/db/leveldbhelper"
	"github.com/it-chain/it-chain-Engine/legacy/domain"
	"github.com/it-chain/it-chain-Engine/legacy/network/comm"
	"github.com/it-chain/it-chain-Engine/legacy/network/comm/msg"
	pb "github.com/it-chain/it-chain-Engine/legacy/network/protos"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/spf13/viper"
)

const (
	WAITING_TRANSACTION = "waiting_transaction"
)

type TransactionServiceImpl struct {
	DB          *leveldbhelper.DBProvider
	Comm        comm.ConnectionManager
	PeerService PeerService
}

func NewTransactionService(path string, comm comm.ConnectionManager, ps PeerService) *TransactionServiceImpl {

	transactionService := &TransactionServiceImpl{DB: leveldbhelper.CreateNewDBProvider(path), Comm: comm, PeerService: ps}

	i, _ := strconv.Atoi(viper.GetString("batchTimer.pushPeerTable"))

	broadCastPeerTableBatcher := NewBatchService(time.Duration(i)*time.Second, transactionService.SendToLeader, false)
	broadCastPeerTableBatcher.Add("Send tx to leader")
	broadCastPeerTableBatcher.Start()

	comm.Subscribe("receive transactions", transactionService.handleTransaction)

	return transactionService
}

func (t *TransactionServiceImpl) Close() {
	t.DB.Close()
}

func (t *TransactionServiceImpl) AddTransaction(tx *domain.Transaction) error {
	db := t.DB.GetDBHandle(WAITING_TRANSACTION)
	serializedTX, err := common.Serialize(tx)
	if err != nil {
		return err
	}

	err = db.Put([]byte(tx.TransactionID), serializedTX, true)
	return err
}

func (t *TransactionServiceImpl) DeleteTransactions(txs []*domain.Transaction) error {
	db := t.DB.GetDBHandle(WAITING_TRANSACTION)
	batch := make(map[string][]byte)

	for _, tx := range txs {
		batch[tx.TransactionID] = nil
	}

	return db.WriteBatch(batch, true)
}

func (t *TransactionServiceImpl) GetTransactions(limit int) ([]*domain.Transaction, error) {

	db := t.DB.GetDBHandle(WAITING_TRANSACTION)
	iter := db.GetIteratorWithPrefix()
	ret := make([]*domain.Transaction, 0)
	cnt := 0

	for iter.Next() {
		val := iter.Value()
		tx := &domain.Transaction{}
		err := common.Deserialize(val, tx)

		if err != nil {
			return nil, err
		}

		ret = append(ret, tx)
		cnt++
		if cnt == limit {
			break
		}
	}

	iter.Release()

	return ret, nil
}

func (t *TransactionServiceImpl) handleTransaction(msg msg.OutterMessage) {

	common.Log.Println("Received Transaction1")

	if txMsg := msg.Message.GetTransaction(); txMsg != nil {
		common.Log.Println("Received Transaction")
		transaction := domain.FromProtoTransaction(*txMsg)
		t.AddTransaction(transaction)
	}

	return
}

func (t *TransactionServiceImpl) SendToLeader(interface{}) {

	//todo max 몇개까지 보낼것인지
	txs, err := t.GetTransactions(100)

	if err != nil {
		common.Log.Println("Error on GetTransactions")
	}

	if len(txs) == 0 {
		common.Log.Println("No transactions to send")
		return
	}

	for _, tx := range txs {

		message := &pb.StreamMessage{}
		message.Content = &pb.StreamMessage_Transaction{
			Transaction: domain.ToProtoTransaction(*tx),
		}

		if err != nil {
			common.Log.Println("fail to serialize message")
		}

		errorCallBack := func(onError error) {
			common.Log.Println("fail to send message error:", onError.Error())
		}

		successCallBack := func(interface{}) {
			common.Log.Println("success to send tx")
			t.DeleteTransactions(txs)
		}

		//todo need leader selection alg
		//내가 리더가 아니고, 리더가 nil아니면 보낸다.
		if t.PeerService.GetLeader() != nil && t.PeerService.GetLeader().PeerID != t.PeerService.GetPeerTable().MyID {
			common.Log.Println("Sending:", domain.ToProtoTransaction(*tx))
			t.Comm.SendStream(message, successCallBack, errorCallBack, t.PeerService.GetLeader().PeerID)
		}
	}
}

func (t *TransactionServiceImpl) CreateTransaction(txData *domain.TxData) (*domain.Transaction, error) {

	transaction := domain.CreateNewTransaction(
		t.PeerService.GetPeerTable().MyID,
		xid.New().String(),
		domain.General,
		time.Now(),
		txData)

	err := t.AddTransaction(transaction)

	if err != nil {
		return nil, errors.New("faild to add transaction")
	}

	return transaction, nil
}
