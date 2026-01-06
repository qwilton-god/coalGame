package company

import (
	"coalGame/company/equipment"
	"coalGame/company/miners"
)

type GameStats struct {
	TotalCoal    miners.Coal     `json:"total_coal"`
	ActiveMiners int             `json:"active_miners_count"`
	History      map[string]int  `json:"history_stats"`
	Inventory    map[string]bool `json:"inventory"`
}

func (c *Company) GetStats() GameStats {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	historyCopy := make(map[string]int)
	for k, v := range c.historyStats {
		historyCopy[k] = v
	}

	inventoryCopy := make(map[string]bool)
	for k, v := range c.inventory {
		inventoryCopy[k] = v
	}

	return GameStats{
		TotalCoal:    c.coal,
		ActiveMiners: len(c.activeMiners),
		History:      historyCopy,
		Inventory:    inventoryCopy,
	}
}

func (c *Company) GetMinerPriceInfo() map[string]int {
	priceInfo := map[string]int{
		"basic":    5,
		"normal":   50,
		"advanced": 450,
	}

	return priceInfo
}

func (c *Company) GetActiveMinersOnClass(class string) ([]miners.Miner, error) {
	activeList := make([]miners.Miner, 0)
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, miner := range c.activeMiners {
		if miner.Info().Class == class {
			activeList = append(activeList, miner)
		}
	}

	if len(activeList) == 0 {
		return activeList, ErrNotFound
	}

	return activeList, nil
}

func (c *Company) EquipmentInfo() []equipment.Product {
	return equipment.AllProducts
}

func (c *Company) EquipmentBuyedInfo() map[string]bool {
	return c.GetStats().Inventory
}
