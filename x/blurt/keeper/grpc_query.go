package keeper

import (
	"github.com/faddat/blurt/x/blurt/types"
)

var _ types.QueryServer = Keeper{}
