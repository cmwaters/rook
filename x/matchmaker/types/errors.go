package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rook module sentinel errors
var (
	ErrGameDoesNotExist = sdkerrors.Register(ModuleName, 1100, "game does not exist")
)
