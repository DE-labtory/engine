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

package api

import (
	"github.com/it-chain/engine/consensus"
	"github.com/it-chain/engine/consensus/infra/adapter"
)

//TODO Api 는 interface 선언 infra에서 구현
type ConsensusApi struct {
	parliamentService adapter.ParliamentService
	consensus         consensus.Consensus
	propagateService  adapter.PropagateService
	confirmService    adapter.ConfirmService
}

// todo : Event Sourcing 첨가
func (cApi ConsensusApi) StartConsensus(userId consensus.MemberId, block consensus.ProposedBlock) error {

	// 합의 시작!! 리더에 의해 시작 만약 블록이 생성되면 Consensus가 필요한지 따져야함
	// consensus를 시작한 멤버 아이디와, 제안된 블록으로 consensus를 만든다.
	peerList, _ := cApi.parliamentService.RequestPeerList()
	if cApi.parliamentService.IsNeedConsensus() {
		createdConsensus, err := consensus.CreateConsensus(peerList, block)
		if err != nil {
			return consensus.CreateConsensusError
		}
		createdConsensus.Start()
		createdPrePrepareMsg := consensus.NewPrePrepareMsg(createdConsensus)
		cApi.propagateService.BroadcastPrePrepareMsg(*createdPrePrepareMsg)

		return nil

	} else {
		cApi.confirmService.ConfirmBlock(block)
	}

	return nil
}

func (cApi ConsensusApi) ReceivePrePrepareMsg(msg consensus.PrePrepareMsg) error {
	// 검증하는 함수 if -> 검증.false == 수용 x
	// msg가 leader에게 온 것인지 검증
	// TODO message Service에 옮김 추가 검증 필요?
	lid, _ := cApi.parliamentService.RequestLeader()
	if lid.ToString() == msg.SenderId {
		// 검증 후 consensus Construct
		createdConsensus, err := consensus.ConstructConsensus(msg)

		if err != nil {
			return consensus.CreateConsensusError
		}
		createdConsensus.ToPrepareState()
		prepareMsg := consensus.NewPrepareMsg(createdConsensus)
		cApi.propagateService.BroadcastPrepareMsg(*prepareMsg)

		return nil
	}

	return consensus.InvalidLeaderIdError
}

func (cApi ConsensusApi) ReceivePrepareMsg(msg consensus.PrepareMsg) error {

	// Prepare Msg 받으면 개수가 2f개 이상인지
	err := cApi.consensus.SavePrepareMsg(&msg)

	if err != nil {
		return consensus.SavePrepareMsgError
	}
	// 2f 조건 체크
	if cApi.parliamentService.CheckPrepareCondition(cApi.consensus.PrepareMsgPool) {
		//조건 만족
		/*tempConsensus := &consensus.Consensus{}

		newConsensus := eventstore.Load(tempConsensus, msg.ConsensusId.Id)

		newCommitMsg := consensus.NewCommitMsg(newConsensus)
		// Commit state로 전환 후 Broadcast
		cApi.consensusService.ToCommitState()
		cApi.propagateService.BroadcastCommitMsg(*newCommitMsg)*/
	} else {
		return nil
	}

	return nil
}

func (cApi ConsensusApi) ReceiveCommitMsg(msg consensus.CommitMsg) error {
	// 커밋 메시지 저
	err := cApi.consensus.SaveCommitMsg(&msg)
	if err != nil {
		return consensus.SaveCommitMsgError
	}

	if cApi.parliamentService.CheckCommitCondition(cApi.consensus.CommitMsgPool) {
		// 조건 만족
		// TODO Client 한테 Response

	} else {
		return nil
	}
	return nil
}
