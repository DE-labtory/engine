package service

import (
	"it-chain/network/comm"
)

type LeaderElectionService interface{

	// 시작
	Run()

	// 종료
	Stop()

	// HeartBeat, Request Vote Massage, Vote 메세지 Receive
	ReceiveMessage(message comm.OutterMessage)

	// 리더 선출 서비스에 피어를 추가
	AddPeerId()

	// DB로 부터 마지막 블록의 해시값을 가져옴
	GetLastBlockHash() string
}