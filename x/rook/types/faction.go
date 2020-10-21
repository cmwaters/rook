package types

const (
	townPopulationRate       = 1
	cityPopulationRate       = 2
	lumbermillProductionRate = 1
	quarryProductionRate     = 1
	farmProductionRate       = 1
)



type position struct {
	X uint16
	Y uint16
}

func NewFaction(name string) *Faction {
	return &Faction{
		Moniker:    name,
		Food:       0,
		Wood:       0,
		Stone:      0,
		Population: make(map[uint32]uint32),
		Settlements: make(map[uint32]Settlement),
	}
}

// Reap takes a Factions structures and accumulates all the resources accrued in that turn
func (f *Faction) Reap() {
	for position, settlement := range f.settlements {
		switch settlement {
		case SETTLEMENT_TOWN:
			if populace, ok := f.population[position]; ok {
				f.population[position] = populace + townPopulationRate
			} else {
				f.population[position] = townPopulationRate
			}
		case SETTLEMENT_CITY, Settlement_SETTLEMENT_CAPITAL:
			if populace, ok := f.population[position]; ok {
				f.population[position] = populace + cityPopulationRate
			} else {
				f.population[position] = cityPopulationRate
			}
		case Lumbermill:
			f.wood += lumbermillProductionRate
		case Quarry:
			f.wood += quarryProductionRate
		case Farm:
			f.food += farmProductionRate
		default: // Outposts for example don't produce anything
			continue
		}
	}
}
