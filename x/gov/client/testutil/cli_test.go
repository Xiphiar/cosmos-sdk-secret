//go:build norace
// +build norace

package testutil

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewIntegrationTestSuite(cfg))

	genesisState := types.DefaultGenesisState()
	genesisState.DepositParams = types.NewDepositParams(sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, types.DefaultMinDepositTokens)), time.Duration(15)*time.Second, sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, types.DefaultMinExpeditedDepositTokens)))
	genesisState.VotingParams = types.NewVotingParams(time.Duration(5)*time.Second, time.Duration(2)*time.Second, []types.ProposalVotingPeriod{})
	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState["gov"] = bz
	suite.Run(t, NewDepositTestSuite(cfg))
}
