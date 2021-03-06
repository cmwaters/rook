package rest

import (
	"net/http"
	"strconv"

	"github.com/cmwaters/rook/x/matchmaker/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createMoveRequest struct {
	BaseReq  rest.BaseReq    `json:"base_req"`
	Creator  string          `json:"creator"`
	Quantity uint32          `json:"quantity"`
	From     types.Position  `json:"from"`
	To       types.Direction `json:"to"`
}

func createMoveHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createMoveRequest
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

		parsedQuantity := req.Quantity

		parsedFrom := req.From

		parsedTo := req.To

		msg := types.NewMsgMove(
			creator,
			parsedQuantity,
			parsedFrom,
			parsedTo,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
