package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cmwaters/rook/x/matchmaker/types"
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers matchmaker-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
    r.HandleFunc("custom/matchmaker/" + types.QueryListFindGame, listFindGameHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
    r.HandleFunc("/matchmaker/findGame", createFindGameHandler(clientCtx)).Methods("POST")

}

