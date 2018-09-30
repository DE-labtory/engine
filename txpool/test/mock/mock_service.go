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

package mock

import (
	"github.com/it-chain/engine/p2p"
)

type PeerQueryService struct {
	GetPeerListFunc func() ([]p2p.Peer, error)
	GetLeaderFunc   func() (p2p.Leader, error)
}

func (m PeerQueryService) GetPeerList() ([]p2p.Peer, error) {
	return m.GetPeerListFunc()
}

func (m PeerQueryService) GetLeader() (p2p.Leader, error) {
	return m.GetLeaderFunc()
}

type EventService struct {
	PublishFunc func(topic string, event interface{}) error
}

func (s EventService) Publish(topic string, event interface{}) error {
	return s.PublishFunc(topic, event)
}
