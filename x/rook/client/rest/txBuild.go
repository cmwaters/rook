package rest

import (
	"net/http"
	"strconv"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cmwaters/rook/x/rook/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createBuildRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Settlement string `json:"settlement"`
	Position string `json:"position"`
	
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

		
		parsedSettlement := req.Settlement
		
		parsedPosition := req.Position
		

		msg := types.NewMsgBuild(
			creator,
			parsedSettlement,
			parsedPosition,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
