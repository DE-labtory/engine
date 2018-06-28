package leveldb

import (
	"errors"

	"github.com/it-chain/it-chain-Engine/p2p"
)
//types of errors
var ErrExistPeer = errors.New("proposed peer already exists")
var ErrNoMatchingPeer = errors.New("no matching peer exists")
var ErrEmptyPeerId = errors.New("empty peer id proposed")
var ErrEmptyPeerTable = errors.New("peer table is empty")

type PeerRepository struct {}

//in memory peer list
var peerTable map[string]p2p.Peer

// 새로운 p2p repo 생성
func NewPeerRepository() (PeerRepository, error) {

	once.Do(func() {
		peerTable = make(map[string]p2p.Peer)
	})
	return PeerRepository{}, nil
}


// 새로운 p2p 를 leveldb에 저장
func (pr *PeerRepository) Save(data p2p.Peer) error {

	// return empty peerID error if peerID is null
	if data.PeerId.Id == "" {
		return ErrEmptyPeerId
	}
	_, exist := peerTable[data.PeerId.Id]
	if exist {
		return ErrExistPeer
	}

	peerTable[data.PeerId.Id] = data

	return nil
}

// p2p 삭제
func (pr *PeerRepository) Remove(id p2p.PeerId) error {
	if len(id.Id) == 0 {
		return ErrEmptyPeerId
	}

	_, exist := peerTable[id.Id]

	if !exist{
		return ErrNoMatchingPeer
	}

	delete(peerTable, id.Id)
	return nil
}

// p2p 읽어옴
func (pr *PeerRepository) FindById(id p2p.PeerId) (p2p.Peer, error) {
	v, exist := peerTable[id.Id]

	if id.Id == "" {
		return v, ErrEmptyPeerId
	}
	//no matching id
	if !exist {
		return v, ErrNoMatchingPeer
	}

	return v, nil
}

// 모든 피어 검색
func (pr *PeerRepository) FindAll() ([]p2p.Peer, error) {
	peers := make([]p2p.Peer, 0)

	if len(peerTable) == 0{
		return peers, ErrEmptyPeerTable
	}

	for _, value := range peerTable{
		peers = append(peers, value)
	}

	return peers, nil
}

func ClearPeerTable(){
	for key := range peerTable{
		delete(peerTable, key)
	}
	peerTable = make(map[string]p2p.Peer)
}