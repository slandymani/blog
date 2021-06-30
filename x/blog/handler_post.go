package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/slandymani/blog/x/blog/keeper"
	"github.com/slandymani/blog/x/blog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
