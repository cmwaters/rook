package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PendingGameState struct {
	playerVotes map[string]bool
	config      *GameConfig
	gameID      string

	total uint32
	count uint32
	timer uint32 // countdown where the pending game is available
}

func NewPendingGame(players []sdk.AccAddress, gameId string, config *GameConfig) *PendingGameState {
	playerMap := make(map[string]bool, len(players))
	for _, p := range players {
		playerMap[p.String()] = false
	}
	return &PendingGameState{
		playerVotes: playerMap,
		config:      config,
		gameID:      gameId,
		total:       uint32(len(players)),
		count:       0,
		timer:       defaultCountdownBlocks,
	}
}

func (p *PendingGameState) Tick() (expired bool, ready bool, players []*sdk.AccAddress) {
	p.timer--
	if p.timer == 0 {
		ready, players := p.Quorum()
		return true, ready, players
	}
	return false, false, nil
}

func (p *PendingGameState) AddVote(player *sdk.AccAddress) bool {
	if _, ok := p.playerVotes[player.String()]; !ok {
		p.playerVotes[player.String()] = true
		p.count++
		if p.count == p.total {
			return true
		}
	}
	return false
}

func (p *PendingGameState) Participants() []*sdk.AccAddress {
	players := make([]*sdk.AccAddress, 0, p.total)
	for playerString, vote := range p.playerVotes {
		if vote {
			player, _ := sdk.AccAddressFromBech32(playerString)
			players = append(players, &player)
		}
	}
	return players
}

func (p *PendingGameState) Quorum() (bool, []*sdk.AccAddress) {
	counter := uint32(0)
	players := make([]*sdk.AccAddress, 0, p.total)
	for playerString, vote := range p.playerVotes {
		if vote {
			counter++
			player, _ := sdk.AccAddressFromBech32(playerString)
			players = append(players, &player)
		}
	}
	if counter > (p.total / 2) {
		return true, players
	}
	return false, nil
}

func (p *PendingGameState) Config() *GameConfig {
	return p.config
}