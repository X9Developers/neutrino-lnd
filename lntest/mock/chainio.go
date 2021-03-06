package mock

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// ChainIO is a mock implementation of the BlockChainIO interface.
type ChainIO struct {
	BestHeight int32
}

// GetBestBlock currently returns dummy values.
func (c *ChainIO) GetBestBlock() (*chainhash.Hash, int32, error) {
	return chaincfg.TestNet3Params.GenesisHash, c.BestHeight, nil
}

// GetUtxo currently returns dummy values.
func (c *ChainIO) GetUtxo(op *wire.OutPoint, _ []byte,
	heightHint uint32, _ <-chan struct{}) (*wire.TxOut, error) {

	return nil, nil
}

// GetBlockHash currently returns dummy values.
func (c *ChainIO) GetBlockHash(blockHeight int64) (*chainhash.Hash, error) {
	return nil, nil
}

// GetBlock currently returns dummy values.
func (c *ChainIO) GetBlock(blockHash *chainhash.Hash) (*wire.MsgBlock, error) {
	return nil, nil
}

// no cache for mock
func (b *ChainIO) FreeCache() error {
	return nil
}

// no cache for mock
func (b *ChainIO) GetRawTxByIndex(blockHeight int64, txIndex uint32) (*wire.MsgTx, error) {
	return nil, nil
}

// no cache for mock
func (b *ChainIO) LoadCache(startHeight uint32) (bool, error) {
	return false, nil
}

func (b *ChainIO) OutputSpent(op *wire.OutPoint) (bool, error) {
	return false, nil
}