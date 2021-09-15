package types

import storetypes "github.com/cosmos/cosmos-sdk/store/types"

// NullGasMeter doesn't care about anything, doesn't check anything
// (unlike Cosmos' InfiniteGasMeter), just always reports "positive" values.
type NullGasMeter struct{}

func (NullGasMeter) GasConsumed() storetypes.Gas {
	return 0
}

func (NullGasMeter) GasConsumedToLimit() storetypes.Gas {
	return 0
}

func (NullGasMeter) Limit() storetypes.Gas {
	return 2 ^ 64
}

func (NullGasMeter) ConsumeGas(amount storetypes.Gas, descriptor string) {
	return
}

func (NullGasMeter) RefundGas(amount storetypes.Gas, descriptor string) {
	return
}

func (NullGasMeter) IsPastLimit() bool {
	return false
}

func (NullGasMeter) IsOutOfGas() bool {
	return false
}

func (NullGasMeter) String() string {
	return "I'm instance of NullGasMeter... so no stats here :)"
}
