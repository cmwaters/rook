module github.com/cmwaters/rook

go 1.14

require (
	github.com/CosmWasm/wasmd v0.11.0
	github.com/cosmos/cosmos-sdk v0.40.0-rc0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/regen-network/cosmos-proto v0.3.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0
	github.com/tendermint/tm-db v0.6.3
	google.golang.org/grpc v1.33.2
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
