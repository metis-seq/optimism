package rollup

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PosClient interface {
	ChainID(context.Context) (*big.Int, error)
	GetSequencerByHeight(context.Context, int64) (common.Address, error)
}

// ValidatePosConfig checks L1 config variables for errors.
func (cfg *Config) ValidatePosConfig(ctx context.Context, client PosClient) error {
	// Validate the L1 Client Chain ID
	if err := cfg.CheckPosChainID(ctx, client); err != nil {
		return err
	}

	// Validate the Rollup L1 Genesis Blockhash
	// if err := cfg.CheckL1GenesisBlockHash(ctx, client); err != nil {
	// 	return err
	// }

	return nil
}

// CheckPosChainID checks that the configured L1 chain ID matches the client's chain ID.
func (cfg *Config) CheckPosChainID(ctx context.Context, client PosClient) error {
	id, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	if cfg.L1ChainID.Cmp(id) != 0 {
		return fmt.Errorf("incorrect Pos RPC chain id %d, expected %d", cfg.PosChainID, id)
	}
	return nil
}

func (cfg *Config) GetSequencerByHeight(ctx context.Context, client PosClient, height int64) (common.Address, error) {
	return client.GetSequencerByHeight(ctx, height)
}
