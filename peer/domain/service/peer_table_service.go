package service

import (
	"github.com/it-chain/it-chain-Engine/peer/domain/model"
	"github.com/it-chain/it-chain-Engine/peer/domain/repository"
)

// 피어 테이블 구조체는
// 자신이 포함된 chain에서 자신의 peer 정보와 리더 정보를, 그리고 리더를 변경하고 가져오는 메소드를 가진다.
type PeerTable struct {
	leader     *model.Peer
	myInfo     *model.Peer
	repository repository.Peer
}

func NewPeerTableService(peerRepo repository.Peer, myinfo *model.Peer) *PeerTable {
	peerRepo.Save(*myinfo)
	return &PeerTable{
		leader:     nil,
		myInfo:     myinfo,
		repository: peerRepo,
	}
}


func (pts *PeerTable) SetLeader(peer *model.Peer) error {
	find, err := pts.repository.FindById(peer.Id)
	if err != nil {
		return err
	}
	pts.leader = find
	return nil
}

func (pts *PeerTable) GetLeader() *model.Peer {
	return pts.leader
}
