package types

func NewFaction(name string, config InitializationConfig) *Faction {
	return &Faction{
		Moniker:     name,
		Resources:   NewResourceSet(config.Resources),
		Population:  make(map[uint32]uint32),
		Settlements: make(map[uint32]Settlement),
	}
}

// Reap takes a Factions structures and accumulates all the resources accrued in that turn
func (f *Faction) Reap(config *ProductionRatesConfig) {
	for position, settlement := range f.Settlements {
		switch settlement {
		case Settlement_TOWN:
			if populace, ok := f.Population[position]; ok {
				f.Population[position] = populace + config.Town
			} else {
				f.Population[position] = config.Town
			}
		case Settlement_CITY, Settlement_CAPITAL:
			if populace, ok := f.Population[position]; ok {
				f.Population[position] = populace + config.City
			} else {
				f.Population[position] = config.City
			}
		case Settlement_LUMBERMILL:
			f.Resources.Wood += config.Lumbermill
		case Settlement_QUARRY:
			f.Resources.Stone += config.Quarry
		case Settlement_FARM:
			f.Resources.Food += config.Farm
		default: // Rooks for example don't produce anything
			continue
		}
	}
}

func NewResourceSet(r *ResourceSet) *ResourceSet {
	return &ResourceSet{Wood: r.Wood, Food: r.Food, Stone: r.Stone}
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
		return &ResourceSet{0,0,0}
	}
}

func (r *ResourceSet) Less(r2 *ResourceSet) bool {
	return r.Wood < r2.Wood && r.Food < r2.Food && r.Stone < r2.Stone
}

func (r *ResourceSet) Subtract(r2 *ResourceSet) {
	r.Wood -= r2.Wood
	r.Food -= r2.Food
	r.Stone -= r2.Stone
}
