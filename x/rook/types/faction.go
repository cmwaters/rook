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
	for position, settlement := range f.Settlements {
		switch settlement {
		case Settlement_TOWN:
			if populace, ok := f.Population[position]; ok {
				f.Population[position] = populace + townPopulationRate
			} else {
				f.Population[position] = townPopulationRate
			}
		case Settlement_CITY, Settlement_CAPITAL:
			if populace, ok := f.Population[position]; ok {
				f.Population[position] = populace + cityPopulationRate
			} else {
				f.Population[position] = cityPopulationRate
			}
		case Settlement_LUMBERMILL:
			f.Wood += lumbermillProductionRate
		case Settlement_QUARRY:
			f.Wood += quarryProductionRate
		case Settlement_FARM:
			f.Food += farmProductionRate
		default: // Rooks for example don't produce anything
			continue
		}
	}
}
