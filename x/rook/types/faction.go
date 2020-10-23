package types

const (
	townPopulationRate       = 1
	cityPopulationRate       = 2
	lumbermillProductionRate = 1
	quarryProductionRate     = 1
	farmProductionRate       = 1
)

func NewFaction(name string) *Faction {
	return &Faction{
		Moniker:     name,
		Resources:   NewResourceSet(),
		Population:  make(map[uint32]uint32),
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
			f.Resources.Wood += lumbermillProductionRate
		case Settlement_QUARRY:
			f.Resources.Stone += quarryProductionRate
		case Settlement_FARM:
			f.Resources.Food += farmProductionRate
		default: // Rooks for example don't produce anything
			continue
		}
	}
}

func NewResourceSet() *ResourceSet {
	return &ResourceSet{Wood: 0, Food: 0, Stone: 0}
}

func ConstructionResources(config *SettlementCostsConfig, settlement Settlement) *ResourceSet {
	switch settlement {
	case Settlement_TOWN:
		return config.Town
	case Settlement_CITY:
		return config.City
	case Settlement_CAPITAL:
		return config.Capital
	case Settlement_FARM:
		return config.Farm
	case Settlement_QUARRY:
		return config.Quarry
	case Settlement_LUMBERMILL:
		return config.Lumbermill
	case Settlement_ROOK:
		return config.Rook
	default:
		return NewResourceSet()
	}
}

func (r *ResourceSet) Less(r2 *ResourceSet) bool {
	return r.Wood < r2.Wood && r.Food < r2.Food && r.Stone < r2.Stone
}
