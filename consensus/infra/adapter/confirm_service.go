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

package adapter

import (
	"errors"

	"github.com/it-chain/engine/common/command"
	"github.com/it-chain/engine/consensus"
	"github.com/it-chain/midgard"
)

type ConfirmService struct {
	publish Publish
}

func NewConfirmService(publish Publish) *ConfirmService {
	return &ConfirmService{
		publish: publish,
	}
}

// todo : command가 아닌 event를 날려야함
func (cs *ConfirmService) ConfirmBlock(block consensus.ProposedBlock) error {
	if block.Seal == nil {
		return errors.New("Block hash is nil")
	}

	if block.Body == nil {
		return errors.New("There is no block")
	}

	cmd := command.ConfirmBlock{
		CommandModel: midgard.CommandModel{},
		Seal:         nil,
		Body:         nil,
	}

	return cs.publish("Command", "block.create", cmd)
}
