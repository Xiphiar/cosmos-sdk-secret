package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// GetParams returns the total set of distribution parameters.
func (k Keeper) GetParams(clientCtx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(clientCtx, &params)
	return params
}

// SetParams sets the distribution parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetCommunityTax returns the current distribution community tax.
func (k Keeper) GetCommunityTax(ctx sdk.Context) (percent sdk.Dec) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyCommunityTax, &percent)
	return percent
}

// GetBaseProposerReward returns the current distribution base proposer rate.
func (k Keeper) GetBaseProposerReward(ctx sdk.Context) (percent sdk.Dec) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyBaseProposerReward, &percent)
	return percent
}

// GetBonusProposerReward returns the current distribution bonus proposer reward
// rate.
func (k Keeper) GetBonusProposerReward(ctx sdk.Context) (percent sdk.Dec) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyBonusProposerReward, &percent)
	return percent
}

// GetWithdrawAddrEnabled returns the current distribution withdraw address
// enabled parameter.
func (k Keeper) GetWithdrawAddrEnabled(ctx sdk.Context) (enabled bool) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyWithdrawAddrEnabled, &enabled)
	return enabled
}

// GetSecretFoundationTax returns the current secret foundation tax.
func (k Keeper) GetSecretFoundationTax(ctx sdk.Context) (tax sdk.Dec) {
	k.paramSpace.Get(ctx, types.ParamSecretFoundationTax, &tax)
	return tax
}

// GetSecretFoundationAddr returns the current secret foundation address.
func (k Keeper) GetSecretFoundationAddr(ctx sdk.Context) (addr string) {
	k.paramSpace.Get(ctx, types.ParamSecretFoundationAddress, &addr)
	return addr
}

// GetSecretFoundationTax returns the current secret foundation tax.
func (k Keeper) GetMinimumRestakeThreshold(ctx sdk.Context) (amount sdk.Dec) {
	k.paramSpace.Get(ctx, types.ParamMinimumRestakeThreshold, &amount)
	return amount
}

// GetSecretFoundationTax returns the current secret foundation tax.
func (k Keeper) GetRestakePeriod(ctx sdk.Context) (amount sdk.Int) {
	k.paramSpace.Get(ctx, types.ParamRestakePeriod, &amount)
	return amount
}
