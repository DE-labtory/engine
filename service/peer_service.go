package service

import "it-chain/domain"

//peer 최상위 service
type PeerService interface{

	//peer table 조회
	GetPeerTable() domain.PeerTable

	//peer info 찾기
	GetPeerInfoByPeerID(peerID string) *domain.PeerInfo

	//peer info
	PushPeerTable(peerIDs []string)

	//update peerTable
	UpdatePeerTable(peerTable domain.PeerTable)

	//Add peer
	AddPeerInfo(peerInfo *domain.PeerInfo)

	//Request Peer Info
	RequestPeerInfo(host string, port string) *domain.PeerInfo

	BroadCastPeerTable(interface{})
}