package syncinterfaces

import (
	"context"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/etherman"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	"github.com/ethereum/go-ethereum/common"
)

type ProcessBlockRangeL1BlocksMode bool

const (
	StoreL1Blocks   ProcessBlockRangeL1BlocksMode = true
	NoStoreL1Blocks ProcessBlockRangeL1BlocksMode = false
)

type BlockRangeProcessor interface {
	ProcessBlockRange(ctx context.Context, blocks []etherman.Block, order map[common.Hash][]etherman.Order, finalizedBlockNumber uint64) error
	ProcessBlockRangeSingleDbTx(ctx context.Context, blocks []etherman.Block, order map[common.Hash][]etherman.Order, finalizedBlockNumber uint64, storeBlocks ProcessBlockRangeL1BlocksMode, dbTx entities.Tx) error
}
