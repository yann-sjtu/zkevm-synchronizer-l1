package l1sync_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/etherman"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock_l1_check_block "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/l1_check_block/mocks"
	l1sync "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/l1_sync"
	mock_l1sync "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/l1_sync/mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestGetL1BlockPointsOk(t *testing.T) {
	mockBlockProtection := mock_l1_check_block.NewL1BlockNumberFetcher(t)
	mockFinalized := mock_l1_check_block.NewL1BlockNumberFetcher(t)
	ctx := context.TODO()
	mockBlockProtection.EXPECT().Description().Return("mock mockBlockProtection").Maybe()
	mockFinalized.EXPECT().Description().Return("mock mockFinalized").Maybe()

	mockBlockProtection.EXPECT().BlockNumber(ctx, nil).Return(uint64(1), nil)
	mockFinalized.EXPECT().BlockNumber(ctx, nil).Return(uint64(2), nil)
	sut := l1sync.NewBlockPointsRetriever(mockBlockProtection, mockFinalized, nil)

	points, err := sut.GetL1BlockPoints(ctx)
	require.NoError(t, err)
	require.Equal(t, uint64(2), points.L1FinalizedBlockNumber)
	require.Equal(t, uint64(1), points.L1LastBlockToSync)
}

func TestGetL1BlockPointsErr1(t *testing.T) {
	mockBlockProtection := mock_l1_check_block.NewL1BlockNumberFetcher(t)
	mockFinalized := mock_l1_check_block.NewL1BlockNumberFetcher(t)
	ctx := context.TODO()
	mockBlockProtection.EXPECT().BlockNumber(ctx, nil).Return(uint64(1), fmt.Errorf("error"))
	//mockFinalized.EXPECT().BlockNumber(ctx, nil).Return(uint64(2), nil)
	sut := l1sync.NewBlockPointsRetriever(mockBlockProtection, mockFinalized, nil)

	_, err := sut.GetL1BlockPoints(ctx)
	require.Error(t, err)
}

func TestGetL1BlockPointsErr2(t *testing.T) {
	mockBlockProtection := mock_l1_check_block.NewL1BlockNumberFetcher(t)
	mockFinalized := mock_l1_check_block.NewL1BlockNumberFetcher(t)
	ctx := context.TODO()
	mockBlockProtection.EXPECT().BlockNumber(ctx, nil).Return(uint64(1), nil)
	mockFinalized.EXPECT().BlockNumber(ctx, nil).Return(uint64(2), fmt.Errorf("error"))
	sut := l1sync.NewBlockPointsRetriever(mockBlockProtection, mockFinalized, nil)

	_, err := sut.GetL1BlockPoints(ctx)
	require.Error(t, err)
}

func TestSyncBlocksSequentialNothingToDo(t *testing.T) {
	testData := newL1SyncData(t)
	testData.mockBlockRetriever.EXPECT().GetL1BlockPoints(testData.ctx).Return(l1sync.BlockPoints{
		L1LastBlockToSync:      99,
		L1FinalizedBlockNumber: 100,
	}, nil)
	resBlock, synced, err := testData.sut.SyncBlocks(testData.ctx, testData.lastEthBlock)
	require.NoError(t, err)
	require.Equal(t, synced, true)
	require.Equal(t, resBlock, testData.lastEthBlock)
}

func TestSyncBlocksSequentialReorgMissingFirstBlockOnRollupResponse(t *testing.T) {
	testData := newL1SyncData(t)
	testData.mockBlockRetriever.EXPECT().GetL1BlockPoints(testData.ctx).Return(l1sync.BlockPoints{
		L1LastBlockToSync:      100,
		L1FinalizedBlockNumber: 100,
	}, nil)
	toBlock := uint64(100)
	testData.mockEth.EXPECT().GetRollupInfoByBlockRange(testData.ctx, uint64(100), &toBlock).Return(nil, nil, nil)
	_, _, err := testData.sut.SyncBlocks(testData.ctx, testData.lastEthBlock)
	require.Error(t, err)
}

func TestSyncBlocksSequentialReorgDetected(t *testing.T) {
	testData := newL1SyncData(t)
	testData.mockBlockRetriever.EXPECT().GetL1BlockPoints(testData.ctx).Return(l1sync.BlockPoints{
		L1LastBlockToSync:      100,
		L1FinalizedBlockNumber: 100,
	}, nil)
	toBlock := uint64(100)
	testData.mockEth.EXPECT().GetRollupInfoByBlockRange(testData.ctx, uint64(100), &toBlock).Return(nil, nil, nil)
	_, _, err := testData.sut.SyncBlocks(testData.ctx, testData.lastEthBlock)
	require.Error(t, err)
}

// First block to request is the genesis
func TestSyncBlocksSequentialStartingNoBlock(t *testing.T) {
	testData := newL1SyncData(t)
	toBlock := uint64(140)
	testData.mockBlockRetriever.EXPECT().GetL1BlockPoints(testData.ctx).Return(l1sync.BlockPoints{
		L1LastBlockToSync:      toBlock,
		L1FinalizedBlockNumber: toBlock,
	}, nil)

	testData.mockEth.EXPECT().GetRollupInfoByBlockRange(testData.ctx, uint64(testData.cfg.GenesisBlockNumber), &toBlock).Return(nil, nil, nil)
	testData.mockBlockProcessor.EXPECT().ProcessBlockRange(testData.ctx, []etherman.Block(nil), map[common.Hash][]etherman.Order(nil), toBlock).Return(nil)

	_, _, err := testData.sut.SyncBlocks(testData.ctx, nil)
	require.NoError(t, err)
}

// --- HELPER FUNCTIONS ----------------------------------------------
type testL1SyncData struct {
	mockBlockRetriever *mock_l1sync.BlockPointsRetriever
	mockEth            *mock_l1sync.EthermanInterface
	mockState          *mock_l1sync.StateL1SeqInterface
	mockBlockProcessor *mock_l1sync.BlockRangeProcessor
	mockReorg          *mock_l1sync.ReorgManager
	sut                *l1sync.L1SequentialSync
	ctx                context.Context
	lastEthBlock       *entities.L1Block
	cfg                l1sync.L1SequentialSyncConfig
}

func newL1SyncData(t *testing.T) *testL1SyncData {
	mockBlock := mock_l1sync.NewBlockPointsRetriever(t)
	mockEth := mock_l1sync.NewEthermanInterface(t)
	mockState := mock_l1sync.NewStateL1SeqInterface(t)
	mockBlockProcessor := mock_l1sync.NewBlockRangeProcessor(t)
	mockReorg := mock_l1sync.NewReorgManager(t)
	cfg := l1sync.L1SequentialSyncConfig{
		SyncChunkSize:      100,
		GenesisBlockNumber: 123,
	}
	sut := l1sync.NewL1SequentialSync(mockBlock, mockEth, mockState, mockBlockProcessor, mockReorg, nil, cfg)
	ctx := context.TODO()
	lastEthBlock := &entities.L1Block{
		BlockNumber: 100,
		HasEvents:   true,
	}
	return &testL1SyncData{
		mockBlockRetriever: mockBlock,
		mockEth:            mockEth,
		mockState:          mockState,
		mockBlockProcessor: mockBlockProcessor,
		mockReorg:          mockReorg,
		sut:                sut,
		ctx:                ctx,
		lastEthBlock:       lastEthBlock,
		cfg:                cfg,
	}
}
