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
	"github.com/it-chain/engine/txpool"
	"github.com/it-chain/midgard"
)

type SyncUpdateCommand struct {
	midgard.EventModel
	sync bool
}

type NodeUpdateCommand struct {
	midgard.EventModel
	Peer
}

type ProposeBlockCommand struct {
	midgard.CommandModel
	// TODO: Transaction이 너무 다름.
	Transactions []txpool.Transaction
}

// consensus에서 합의된 블록이 넘어오면 block pool에 저장한다.
type ConfirmBlockCommand struct {
	midgard.CommandModel
	Block Block
}

type BlockValidateCommand struct {
	midgard.CommandModel
	Block Block
}

type GrpcDeliverCommand struct {
	midgard.CommandModel
	Recipients []string
	Body       []byte
	Protocol   string
}

type GrpcReceiveCommand struct {
	midgard.CommandModel
	Body         []byte
	ConnectionID string
	Protocol     string
	FromPeer     Peer
}
