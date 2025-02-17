package etrog

import (
	"context"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/model"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/actions"
)

type L1InfoTreeLeaf = entities.L1InfoTreeLeaf
type stateTxType = entities.Tx
type ForkIdType = actions.ForkIdType
type VirtualBatch = entities.VirtualBatch
type SequencedBatches = entities.SequencedBatches
type SequenceOfBatches = model.SequenceOfBatches

type stateOnSequencedBatchesInterface interface {
	OnSequencedBatchesOnL1(ctx context.Context, seq SequenceOfBatches, dbTx stateTxType) error
}
