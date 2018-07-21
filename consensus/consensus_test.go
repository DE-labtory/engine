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

package consensus

import (
	"testing"

	"errors"

	"github.com/it-chain/engine/common/event"
	"github.com/it-chain/engine/consensus/test/mock"
	"github.com/it-chain/engine/core/eventstore"
	"github.com/it-chain/midgard"
	"github.com/stretchr/testify/assert"
)

func TestPrepareMsgPool_Save(t *testing.T) {
	// given
	pPool := NewPrepareMsgPool()

	// case 1 : save
	pMsg := PrepareMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s1",
		BlockHash:   make([]byte, 0),
	}

	// when
	pPool.Save(&pMsg)

	// then
	assert.Equal(t, 1, len(pPool.messages))

	// case 2 : save
	pMsg = PrepareMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s2",
		BlockHash:   make([]byte, 0),
	}

	// when
	pPool.Save(&pMsg)

	// then
	assert.Equal(t, 2, len(pPool.messages))

	// case 3 : same sender
	pMsg = PrepareMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s2",
		BlockHash:   make([]byte, 0),
	}

	// when
	pPool.Save(&pMsg)

	// then
	assert.Equal(t, 2, len(pPool.messages))

	// case 4 : block hash is is nil
	pMsg = PrepareMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s3",
		BlockHash:   nil,
	}

	// when
	pPool.Save(&pMsg)

	// then
	assert.Equal(t, 2, len(pPool.messages))
}

func TestPrepareMsgPool_RemoveAllMsgs(t *testing.T) {
	// given
	pPool := NewPrepareMsgPool()

	pMsg := PrepareMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s1",
		BlockHash:   make([]byte, 0),
	}

	pPool.Save(&pMsg)

	// when
	pPool.RemoveAllMsgs()

	// then
	assert.Equal(t, 0, len(pPool.messages))
}

func TestCommitMsgPool_Save(t *testing.T) {
	// given
	cPool := NewCommitMsgPool()

	// case 1 : save
	cMsg := CommitMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s1",
	}

	// when
	cPool.Save(&cMsg)

	// then
	assert.Equal(t, 1, len(cPool.messages))

	// case 2 : save
	cMsg = CommitMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s2",
	}

	// when
	cPool.Save(&cMsg)

	// then
	assert.Equal(t, 2, len(cPool.messages))

	// case 3 : same sender
	cMsg = CommitMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s2",
	}

	// when
	cPool.Save(&cMsg)

	// then
	assert.Equal(t, 2, len(cPool.messages))
}

func TestCommitMsgPool_RemoveAllMsgs(t *testing.T) {
	// given
	cPool := NewCommitMsgPool()

	cMsg := CommitMsg{
		ConsensusId: ConsensusId{"c1"},
		SenderId:    "s1",
	}

	cPool.Save(&cMsg)

	// when
	cPool.RemoveAllMsgs()

	// then
	assert.Equal(t, 0, len(cPool.messages))
}

func TestConsensus_SavePrepareMsg(t *testing.T) {
	// given
	c := Consensus{
		ConsensusID:     NewConsensusId("c1"),
		Representatives: nil,
		Block: ProposedBlock{
			Seal: make([]byte, 0),
		},
		CurrentState:   IDLE_STATE,
		PrepareMsgPool: NewPrepareMsgPool(),
		CommitMsgPool:  NewCommitMsgPool(),
	}

	eventRepository := mock.MockEventRepository{}

	// case 1 : save
	pMsg := PrepareMsg{
		ConsensusId: NewConsensusId("c1"),
		SenderId:    "s1",
		BlockHash:   make([]byte, 0),
	}

	eventRepository.SaveFunc = func(aggregateID string, events ...midgard.Event) error {
		assert.Equal(t, "c1", events[0].(*event.PrepareMsgAdded).GetID())
		assert.Equal(t, "s1", events[0].(*event.PrepareMsgAdded).SenderId)
		return nil
	}
	eventstore.InitForMock(eventRepository)

	// when
	err := c.SavePrepareMsg(&pMsg)

	// then
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.PrepareMsgPool.messages))

	// case 2 : incorrect consensus ID
	pMsg = PrepareMsg{
		ConsensusId: NewConsensusId("c2"),
		SenderId:    "s1",
		BlockHash:   make([]byte, 0),
	}

	eventRepository.SaveFunc = func(aggregateID string, events ...midgard.Event) error {
		assert.Equal(t, "c2", events[0].(*event.PrepareMsgAdded).GetID())
		assert.Equal(t, "s1", events[0].(*event.PrepareMsgAdded).SenderId)
		return errors.New("Consensus ID is not same")
	}
	eventstore.InitForMock(eventRepository)

	// when
	err = c.SavePrepareMsg(&pMsg)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(c.PrepareMsgPool.messages))
}

func TestConsensus_SaveCommitMsg(t *testing.T) {
	// given
	c := Consensus{
		ConsensusID:     NewConsensusId("c1"),
		Representatives: nil,
		Block: ProposedBlock{
			Seal: make([]byte, 0),
		},
		CurrentState:   IDLE_STATE,
		PrepareMsgPool: NewPrepareMsgPool(),
		CommitMsgPool:  NewCommitMsgPool(),
	}

	eventRepository := mock.MockEventRepository{}

	// case 1 : save
	cMsg := CommitMsg{
		ConsensusId: NewConsensusId("c1"),
		SenderId:    "s1",
	}

	eventRepository.SaveFunc = func(aggregateID string, events ...midgard.Event) error {
		assert.Equal(t, "c1", events[0].(*event.CommitMsgAdded).GetID())
		assert.Equal(t, "s1", events[0].(*event.CommitMsgAdded).SenderId)
		return nil
	}
	eventstore.InitForMock(eventRepository)

	// when
	err := c.SaveCommitMsg(&cMsg)

	// then
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.CommitMsgPool.messages))

	// case 2 : incorrect consensus ID
	cMsg = CommitMsg{
		ConsensusId: NewConsensusId("c2"),
		SenderId:    "s1",
	}

	eventRepository.SaveFunc = func(aggregateID string, events ...midgard.Event) error {
		assert.Equal(t, "c2", events[0].(*event.CommitMsgAdded).GetID())
		assert.Equal(t, "s1", events[0].(*event.CommitMsgAdded).SenderId)
		return errors.New("Consensus ID is not same")
	}
	eventstore.InitForMock(eventRepository)

	// when
	err = c.SaveCommitMsg(&cMsg)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(c.CommitMsgPool.messages))
}
