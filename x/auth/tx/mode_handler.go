package tx

import (
	"fmt"

	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// DefaultSignModes are the default sign modes enabled for protobuf transactions.
var DefaultSignModes = []signingtypes.SignMode{
	signingtypes.SignMode_SIGN_MODE_DIRECT,
	signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
	signingtypes.SignMode_SIGN_MODE_EIP_191,
}

// makeSignModeHandler returns the default protobuf SignModeHandler supporting
// SIGN_MODE_DIRECT, SIGN_MODE_LEGACY_AMINO_JSON, and SIGN_MODE_SIGN_MODE_EIP_191.
func makeSignModeHandler(modes []signingtypes.SignMode) signing.SignModeHandler {
	if len(modes) < 1 {
		panic(fmt.Errorf("no sign modes enabled"))
	}

	handlers := make([]signing.SignModeHandler, len(modes))

	for i, mode := range modes {
		switch mode {
		case signingtypes.SignMode_SIGN_MODE_DIRECT:
			handlers[i] = signModeDirectHandler{}
		case signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
			handlers[i] = signModeLegacyAminoJSONHandler{}
		case signingtypes.SignMode_SIGN_MODE_EIP_191:
			handlers[i] = signModeEIP191Handler{}
		default:
			panic(fmt.Errorf("unsupported sign mode %+v", mode))
		}
	}

	return signing.NewSignModeHandlerMap(
		modes[0],
		handlers,
	)
}
