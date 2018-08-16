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

	"github.com/it-chain/engine/common"
	"github.com/it-chain/engine/common/command"
	"github.com/it-chain/engine/consensus/pbft"
	"github.com/it-chain/midgard"
	"github.com/rs/xid"
)

type Publish func(topic string, data interface{}) (err error)

type PropagateService struct {
	publish Publish
}

func NewPropagateService(publish Publish) PropagateService {
	return PropagateService{
		publish: publish,
	}
}

func (ps PropagateService) BroadcastPrePrepareMsg(msg pbft.PrePrepareMsg, representatives []*pbft.Representative) error {
	if msg.StateID.ID == "" {
		return errors.New("State ID is empty")
	}

	if msg.ProposedBlock.Body == nil {
		return errors.New("Block is empty")
	}

	SerializedMsg, err := common.Serialize(msg)

	if err != nil {
		return err
	}

	if err = ps.broadcastMsg(SerializedMsg, "PrePrepareMsgProtocol", representatives); err != nil {
		return err
	}

	return nil
}

func (ps PropagateService) BroadcastPrepareMsg(msg pbft.PrepareMsg, representatives []*pbft.Representative) error {
	if msg.StateID.ID == "" {
		return errors.New("State ID is empty")
	}

	if msg.BlockHash == nil {
		return errors.New("Block hash is empty")
	}

	SerializedMsg, err := common.Serialize(msg)

	if err != nil {
		return err
	}

	if err = ps.broadcastMsg(SerializedMsg, "PrepareMsgProtocol", representatives); err != nil {
		return err
	}

	return nil
}

func (ps PropagateService) BroadcastCommitMsg(msg pbft.CommitMsg, representatives []*pbft.Representative) error {
	if msg.StateID.ID == "" {
		return errors.New("State ID is empty")
	}

	SerializedMsg, err := common.Serialize(msg)

	if err != nil {
		return err
	}

	if err = ps.broadcastMsg(SerializedMsg, "CommitMsgProtocol", representatives); err != nil {
		return err
	}

	return nil
}

func (ps PropagateService) broadcastMsg(SerializedMsg []byte, protocol string, representatives []*pbft.Representative) error {
	if SerializedMsg == nil {
		return errors.New("Message is empty")
	}

	command, err := createDeliverGrpcCommand(protocol, SerializedMsg)

	if err != nil {
		return err
	}

	for _, r := range representatives {
		command.RecipientList = append(command.RecipientList, r.GetID())
	}

	return ps.publish("message.deliver", command)
}

func createDeliverGrpcCommand(protocol string, body interface{}) (command.DeliverGrpc, error) {
	data, err := common.Serialize(body)

	if err != nil {
		return command.DeliverGrpc{}, err
	}

	return command.DeliverGrpc{
		CommandModel: midgard.CommandModel{
			ID: xid.New().String(),
		},
		RecipientList: make([]string, 0),
		Body:          data,
		Protocol:      protocol,
	}, nil
}
