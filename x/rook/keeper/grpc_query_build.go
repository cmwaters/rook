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

func (k Keeper) AllBuild(c context.Context, req *types.QueryAllBuildRequest) (*types.QueryAllBuildResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var builds []*types.MsgBuild
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	buildStore := prefix.NewStore(store, types.KeyPrefix(types.BuildKey))

	pageRes, err := query.Paginate(buildStore, req.Pagination, func(key []byte, value []byte) error {
		var build types.MsgBuild
		if err := k.cdc.UnmarshalBinaryBare(value, &build); err != nil {
			return err
		}

		builds = append(builds, &build)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBuildResponse{Build: builds, Pagination: pageRes}, nil
}
