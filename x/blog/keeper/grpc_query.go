package keeper

import (
	"github.com/slandymani/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
