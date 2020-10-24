package rest

import (
	"github.com/gorilla/mux"

	"github.com/cmwaters/rook/x/rook/types"
	"github.com/cosmos/cosmos-sdk/client"
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers rook-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/custom/rook/"+types.QueryListBuild, listBuildHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/custom/rook/"+types.QueryListMove, listMoveHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/custom/rook/build", createBuildHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/custom/rook/move", createMoveHandler(clientCtx)).Methods("POST")
}
