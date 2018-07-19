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

package blockchain

import (
	"time"

	"github.com/it-chain/midgard"
)

type SyncStartEvent struct {
	midgard.EventModel
}

type SyncDoneEvent struct {
	midgard.EventModel
}

// event when block is committed to event store
type BlockCommittedEvent struct {
	midgard.EventModel
	State string
}

// event when block is staged to event store
type BlockStagedEvent struct {
	midgard.EventModel
	State string
}

//event when block is created in event store
type BlockCreatedEvent struct {
	midgard.EventModel
	Seal      []byte
	PrevSeal  []byte
	Height    uint64
	TxList    []byte
	TxSeal    [][]byte
	Timestamp time.Time
	Creator   []byte
	State     string
}
