package rest

import (
	"net/http"
	"strconv"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createBuildRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Creator    string       `json:"creator"`
	Settlement uint32       `json:"settlement"`
	X          uint32       `json:"x"`
	Y          uint32       `json:"y"`
}

func createBuildHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBuildRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var parsedSettlement types.Settlement
		switch req.Settlement {
		case 1:
			parsedSettlement = types.Settlement_TOWN
		case 2:
			parsedSettlement = types.Settlement_CITY
		case 3:
			parsedSettlement = types.Settlement_CAPITAL
		case 4:
			parsedSettlement = types.Settlement_LUMBERMILL
		case 5:
			parsedSettlement = types.Settlement_QUARRY
		case 6:
			parsedSettlement = types.Settlement_FARM
		case 7:
			parsedSettlement = types.Settlement_ROOK
		}

		parsedPosition := types.Position{X: req.X, Y: req.Y}

		msg := types.NewMsgBuild(
			creator,
			parsedSettlement,
			parsedPosition,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
