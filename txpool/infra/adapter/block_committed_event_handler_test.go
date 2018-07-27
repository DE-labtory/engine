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

package adapter_test

//
//import (
//	"testing"
//
//	"github.com/it-chain/engine/txpool"
//	"github.com/it-chain/engine/txpool/infra/adapter"
//	"github.com/it-chain/midgard"
//	"github.com/magiconair/properties/assert"
//)
//
//type MockTransactionRepository struct {
//	SaveFunc     func(transaction txpool.Transaction) error
//	RemoveFunc   func(id txpool.TransactionId) error
//	FindByIdFunc func(id txpool.TransactionId) (txpool.Transaction, error)
//	FindAllFunc  func() ([]txpool.Transaction, error)
//}
//
//func (m MockTransactionRepository) Save(transaction txpool.Transaction) error {
//	return m.SaveFunc(transaction)
//}
//
//func (m MockTransactionRepository) Remove(id txpool.TransactionId) error {
//	return m.RemoveFunc(id)
//}
//
//func (m MockTransactionRepository) FindBlockById(id txpool.TransactionId) (txpool.Transaction, error) {
//	return m.FindByIdFunc(id)
//}
//
//func (m MockTransactionRepository) FindAllBlock() ([]txpool.Transaction, error) {
//	return m.FindAllFunc()
//}
//
//type MockLeaderRepository struct {
//	GetLeaderFunc func() txpool.Leader
//	SetLeaderFunc func(leader txpool.Leader)
//}
//
//func (m MockLeaderRepository) GetLeader() txpool.Leader {
//	return m.GetLeaderFunc()
//}
//
//func (m MockLeaderRepository) SetLeader(leader txpool.Leader) {
//	m.SetLeaderFunc(leader)
//}
//
//func TestTxEventHandler_HandleTxCreatedEvent(t *testing.T) {
//
//	//given
//	tests := map[string]struct {
//		input txpool.TxCreatedEvent
//		err   error
//	}{
//		"handle success": {
//			input: txpool.TxCreatedEvent{
//				PublishPeerId: "1",
//				ID:            "123",
//				Jsonrpc:       "13",
//				ICodeID:       "icode1",
//				EventModel: midgard.EventModel{
//					ID: "txID",
//				},
//			},
//			err: nil,
//		},
//		"handle fail": {
//			input: txpool.TxCreatedEvent{},
//			err:   adapter.ErrNoEventID,
//		},
//	}
//
//	mockTxRepo := MockTransactionRepository{}
//	mockTxRepo.SaveFunc = func(transaction txpool.Transaction) error {
//		assert.Equal(t, transaction.GetID(), "txID")
//		assert.Equal(t, transaction.TxData.ICodeID, "icode1")
//		return nil
//	}
//
//	mockLeaderRepo := MockLeaderRepository{}
//	event_handler := adapter.NewRepositoryProjector(mockTxRepo, mockLeaderRepo)
//
//	for testName, test := range tests {
//		t.Logf("Running test case %s", testName)
//
//		err := event_handler.HandleTxCreatedEvent(test.input)
//		assert.Equal(t, err, test.err)
//	}
//}
//
//func TestTxEventHandler_HandleTxDeletedEvent(t *testing.T) {
//
//	//given
//	tests := map[string]struct {
//		input txpool.TxDeletedEvent
//		err   error
//	}{
//		"handle success": {
//			input: txpool.TxDeletedEvent{
//				EventModel: midgard.EventModel{
//					ID: "txID",
//				},
//			},
//			err: nil,
//		},
//		"handle fail": {
//			input: txpool.TxDeletedEvent{},
//			err:   adapter.ErrNoEventID,
//		},
//	}
//
//	for testName, test := range tests {
//		t.Logf("Running test case %s", testName)
//
//		mockTxRepo := MockTransactionRepository{}
//		mockTxRepo.RemoveFunc = func(id txpool.TransactionId) error {
//
//			assert.Equal(t, id, "txID")
//			return nil
//		}
//
//		mockLeaderRepo := MockLeaderRepository{}
//		event_handler := adapter.NewRepositoryProjector(mockTxRepo, mockLeaderRepo)
//
//		err := event_handler.HandleTxDeletedEvent(test.input)
//		assert.Equal(t, err, test.err)
//	}
//}
//
//func TestTxEventHandler_HandleLeaderChangedEvent(t *testing.T) {
//
//	//given
//	tests := map[string]struct {
//		input txpool.LeaderChangedEvent
//		err   error
//	}{
//		"handle success": {
//			input: txpool.LeaderChangedEvent{
//				EventModel: midgard.EventModel{
//					ID: "leaderID",
//				},
//			},
//			err: nil,
//		},
//		"handle fail": {
//			input: txpool.LeaderChangedEvent{},
//			err:   adapter.ErrNoEventID,
//		},
//	}
//
//	for testName, test := range tests {
//		t.Logf("Running test case %s", testName)
//
//		mockTxRepo := MockTransactionRepository{}
//		mockLeaderRepo := MockLeaderRepository{}
//
//		mockLeaderRepo.SetLeaderFunc = func(leader txpool.Leader) {
//
//			assert.Equal(t, leader.LeaderId.Id, "leaderID")
//		}
//
//		projector := adapter.NewRepositoryProjector(mockTxRepo, mockLeaderRepo)
//
//		err := projector.HandleLeaderChangedEvent(test.input)
//		assert.Equal(t, err, test.err)
//	}
//}
