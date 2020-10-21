package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cmwaters/rook/x/rook/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllMove(c context.Context, req *types.QueryAllMoveRequest) (*types.QueryAllMoveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var moves []*types.MsgMove
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	moveStore := prefix.NewStore(store, types.KeyPrefix(types.MoveKey))

	pageRes, err := query.Paginate(moveStore, req.Pagination, func(key []byte, value []byte) error {
		var move types.MsgMove
		if err := k.cdc.UnmarshalBinaryBare(value, &move); err != nil {
			return err
		}

		moves = append(moves, &move)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMoveResponse{Move: moves, Pagination: pageRes}, nil
}
