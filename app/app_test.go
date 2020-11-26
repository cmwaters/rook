package app

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/CosmWasm/wasmd/x/wasm"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestWasmdExport(t *testing.T) {
	db := db.NewMemDB()
	gapp := NewRookApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, "", 0, wasm.EnableAllProposals, EmptyAppOptions{})
	err := setGenesis(gapp)
	require.NoError(t, err)

	// Making a new app object with the db, so that initchain hasn't been called
	newGapp := NewRookApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, "", 0, wasm.EnableAllProposals, EmptyAppOptions{})
	_, err = newGapp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlackListedAddrs(t *testing.T) {
	db := db.NewMemDB()
	gapp := NewRookApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, "", 0, wasm.EnableAllProposals, EmptyAppOptions{})

	for acc := range maccPerms {
		require.Equal(t, !allowedReceivingModAcc[acc], gapp.bankKeeper.BlockedAddr(gapp.accountKeeper.GetModuleAddress(acc)))
	}
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}

func TestGetEnabledProposals(t *testing.T) {
	cases := map[string]struct {
		proposalsEnabled string
		specificEnabled  string
		expected         []wasm.ProposalType
	}{
		"all disabled": {
			proposalsEnabled: "false",
			expected:         wasm.DisableAllProposals,
		},
		"all enabled": {
			proposalsEnabled: "true",
			expected:         wasm.EnableAllProposals,
		},
		"some enabled": {
			proposalsEnabled: "okay",
			specificEnabled:  "StoreCode,InstantiateContract",
			expected:         []wasm.ProposalType{wasm.ProposalTypeStoreCode, wasm.ProposalTypeInstantiateContract},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ProposalsEnabled = tc.proposalsEnabled
			EnableSpecificProposals = tc.specificEnabled
			proposals := GetEnabledProposals()
			assert.Equal(t, tc.expected, proposals)
		})
	}
}

func setGenesis(gapp *RookApp) error {
	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		return err
	}

	// Initialize the chain
	gapp.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	gapp.Commit()
	return nil
}