package rest

import (
	"fmt"
	"net/http"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func listBuildHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("/custom/%s/list-build", types.QuerierRoute), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
