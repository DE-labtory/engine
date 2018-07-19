/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package txpool

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/it-chain/engine/common"
	"github.com/it-chain/engine/core/eventstore"
	"github.com/it-chain/midgard"
	"github.com/rs/xid"
)

const (
	VALID   TransactionStatus = 0
	INVALID TransactionStatus = 1

	General TransactionType = 0 + iota
)

type TransactionId = string

type TransactionStatus int
type TransactionType int

//TxData Declaration
const (
	Invoke TxDataType = "invoke"
	Query  TxDataType = "query"
)

type TxDataType string

type TxData struct {
	Jsonrpc string
	Method  TxDataType
	Params  Param
	ICodeID string
}

type Param struct {
	Function string
	Args     []string
}

//Aggregate root must implement aggregate interface
type Transaction struct {
	TxId          TransactionId
	PublishPeerId string
	TxStatus      TransactionStatus
	TxHash        string
	TimeStamp     time.Time
	TxData        TxData
}

// must implement id method
func (t Transaction) GetID() string {
	return string(t.TxId)
}

// must implement on method
func (t *Transaction) On(event midgard.Event) error {

	switch v := event.(type) {

	case *TxCreatedEvent:

		t.TxId = TransactionId(v.ID)
		t.PublishPeerId = v.PublishPeerId
		t.TxStatus = TransactionStatus(v.TxStatus)
		t.TxHash = v.TxHash
		t.TimeStamp = v.TimeStamp
		t.TxData = TxData{
			Params:  v.Params,
			Method:  TxDataType(v.Method),
			Jsonrpc: v.Jsonrpc,
			ICodeID: v.ICodeID,
		}

	case *TxDeletedEvent:
		t.TxId = ""

	default:
		return errors.New(fmt.Sprintf("unhandled event [%s]", v))
	}

	return nil
}

func (t Transaction) Serialize() ([]byte, error) {
	return common.Serialize(t)
}

func Deserialize(b []byte, transaction *Transaction) error {

	err := json.Unmarshal(b, transaction)

	if err != nil {
		return err
	}

	return nil
}

func CalTxHash(txData TxData, publishPeerId string, txId TransactionId, timeStamp time.Time) string {

	hashArgs := []string{
		txData.Jsonrpc,
		string(txData.Method),
		string(txData.Params.Function),
		txData.ICodeID,
		publishPeerId,
		timeStamp.String(),
		string(txId),
	}

	for _, str := range txData.Params.Args {
		hashArgs = append(hashArgs, str)
	}

	return common.ComputeSHA256(hashArgs)
}

func CreateTransaction(publisherId string, txData TxData) (Transaction, error) {

	id := xid.New().String()
	timeStamp := time.Now()
	hash := CalTxHash(txData, publisherId, TransactionId(id), timeStamp)

	event := &TxCreatedEvent{
		EventModel: midgard.EventModel{
			ID:   id,
			Type: "transaction.created",
		},
		PublishPeerId: publisherId,
		TxStatus:      int(VALID),
		TxHash:        hash,
		TimeStamp:     timeStamp,
		ICodeID:       txData.ICodeID,
		Jsonrpc:       txData.Jsonrpc,
		Method:        string(txData.Method),
		Params:        txData.Params,
	}

	tx := &Transaction{}

	if err := saveAndOn(tx, event); err != nil {
		return *tx, err
	}

	return *tx, nil
}

func DeleteTransaction(transaction Transaction) error {

	event := &TxDeletedEvent{
		EventModel: midgard.EventModel{
			ID:   transaction.TxId,
			Type: "transaction.deleted",
		},
	}

	if err := saveAndOn(&transaction, event); err != nil {
		return err
	}

	return nil
}

//apply on aggrgate and publish to eventstore
func saveAndOn(aggregate midgard.Aggregate, event midgard.Event) error {

	//must do call on func first!!!
	//after save events if aggregate.On failed then data inconsistency will be occurred
	if err := aggregate.On(event); err != nil {
		return err
	}

	if err := eventstore.Save(event.GetID(), event); err != nil {
		return err
	}

	return nil
}
