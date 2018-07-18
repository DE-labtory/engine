package api_gateway

import (
	"github.com/it-chain/it-chain-Engine/blockchain"
	"github.com/it-chain/yggdrasill"

	"sync"
)

type BlockQueryApi struct {
	CommitedBlockRepository
}

type CommitedBlockRepository interface {
	Save(block blockchain.Block) error
	GetLastBlock() (blockchain.Block, error)
	GetBlockByHeight(height blockchain.BlockHeight) (blockchain.Block, error)
}

type CommitedBlockRepositoryImpl struct {
	mux *sync.RWMutex
	yggdrasill.BlockStorageManager
}
