package keeper

import (
	"context"

	"github.com/cmwaters/rook/x/matchmaker/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllFindGame(c context.Context, req *types.QueryAllFindGameRequest) (*types.QueryAllFindGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var findGames []*types.MsgFindGame
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	findGameStore := prefix.NewStore(store, types.KeyPrefix(types.FindGameKey))

	pageRes, err := query.Paginate(findGameStore, req.Pagination, func(key []byte, value []byte) error {
		var findGame types.MsgFindGame
		if err := k.cdc.UnmarshalBinaryBare(value, &findGame); err != nil {
			return err
		}

		findGames = append(findGames, &findGame)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFindGameResponse{FindGame: findGames, Pagination: pageRes}, nil
}
