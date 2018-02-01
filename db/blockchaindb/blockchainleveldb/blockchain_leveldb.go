package blockchainleveldb

import (
	"it-chain/db/leveldbhelper"
	"it-chain/service/blockchain"
	"it-chain/common"
	"fmt"
	"github.com/spf13/viper"
)

var logger = common.GetLogger("blockchain_leveldb.go")

const (
	BLOCK_HASH_DB = "block_hash"
	BLOCK_NUMBER_DB = "block_number"
	TRANSACTION_DB = "transaction"
	UTIL_DB = "util"

	LAST_BLOCK_KEY = "last_block"
	UNCONFIRMED_BLOCK_KEY = "unconfirmed_block"
)

type BlockchainLevelDB struct {
	DBProvider *leveldbhelper.DBProvider
}

func CreateNewBlockchainLevelDB(levelDBPath string) *BlockchainLevelDB {
	if levelDBPath == "" {
		levelDBPath = viper.GetString("database.leveldb.default_path")
	}
	levelDBProvider := leveldbhelper.CreateNewDBProvider(levelDBPath)
	return &BlockchainLevelDB{levelDBProvider}
}

func (l *BlockchainLevelDB) Close() {
	l.DBProvider.Close()
}

func (l *BlockchainLevelDB) AddBlock(block *blockchain.Block) error {
	blockHashDB := l.DBProvider.GetDBHandle(BLOCK_HASH_DB)
	blockNumberDB := l.DBProvider.GetDBHandle(BLOCK_NUMBER_DB)
	transactionDB := l.DBProvider.GetDBHandle(TRANSACTION_DB)
	utilDB := l.DBProvider.GetDBHandle(UTIL_DB)

	serializedBlock, err := common.Serialize(block)
	if err != nil {
		return err
	}

	err = blockNumberDB.Put([]byte(fmt.Sprint(block.Header.Number)), serializedBlock, true)
	if err != nil {
		return err
	}

	err = blockHashDB.Put([]byte(block.Header.BlockHash), serializedBlock, true)
	if err != nil {
		return err
	}

	err = utilDB.Put([]byte(LAST_BLOCK_KEY), serializedBlock, true)
	if err != nil {
		return err
	}

	for _, tx := range block.Transactions {
		serializedTx, err := common.Serialize(tx)
		if err != nil {
			return err
		}

		err = transactionDB.Put([]byte(tx.TransactionID), serializedTx, true)
		if err != nil {
			return err
		}

		err = utilDB.Put([]byte(tx.TransactionID), []byte(block.Header.BlockHash), true)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *BlockchainLevelDB) AddUnconfirmedBlock(block *blockchain.Block) error {
	utilDB := l.DBProvider.GetDBHandle(UTIL_DB)

	serializedBlock, err := common.Serialize(block)
	if err != nil {
		return err
	}

	err = utilDB.Put([]byte(block.Header.BlockHash), serializedBlock, true)
	if err != nil {
		return err
	}

	return nil
}

func (l *BlockchainLevelDB) GetBlockByNumber(blockNumber uint64) (*blockchain.Block, error) {
	blockNumberDB := l.DBProvider.GetDBHandle(BLOCK_NUMBER_DB)

	serializedBlock, err := blockNumberDB.Get([]byte(fmt.Sprint(blockNumber)))
	if err != nil {
		return nil, err
	}

	block := &blockchain.Block{}
	err = common.Deserialize(serializedBlock, block)
	if err != nil {
		return nil, err
	}

	return block, err
}

func (l *BlockchainLevelDB) GetBlockByHash(hash string) (*blockchain.Block, error) {
	blockHashDB := l.DBProvider.GetDBHandle(BLOCK_HASH_DB)

	serializedBlock, err := blockHashDB.Get([]byte(hash))
	if err != nil {
		return nil, err
	}

	block := &blockchain.Block{}
	err = common.Deserialize(serializedBlock, block)
	if err != nil {
		return nil, err
	}

	return block, err
}

func (l *BlockchainLevelDB) GetLastBlock() (*blockchain.Block, error) {
	utilDB := l.DBProvider.GetDBHandle(UTIL_DB)

	serializedBlock, err := utilDB.Get([]byte(LAST_BLOCK_KEY))
	if err != nil {
		return nil, err
	}

	block := &blockchain.Block{}
	err = common.Deserialize(serializedBlock, block)
	if err != nil {
		return nil, err
	}

	return block, err
}

func (l *BlockchainLevelDB) GetTransactionByTxID(txid string) (*blockchain.Transaction, error) {
	transactionDB := l.DBProvider.GetDBHandle(TRANSACTION_DB)

	serializedTX, err := transactionDB.Get([]byte(txid))
	if err != nil {
		return nil, err
	}

	transaction := &blockchain.Transaction{}
	err = common.Deserialize(serializedTX, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, err
}

func (l *BlockchainLevelDB) GetBlockByTxID(txid string) (*blockchain.Block, error) {
	utilDB := l.DBProvider.GetDBHandle(UTIL_DB)

	blockHash, err := utilDB.Get([]byte(txid))
	if err != nil {
		return nil, err
	}

	return l.GetBlockByHash(string(blockHash))
}