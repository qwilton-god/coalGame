package company

import (
	"coalGame/company/equipment"
	"coalGame/company/miners"
	"context"
	"fmt"
	"sync"
	"time"
)

type Company struct {
	ctx       context.Context
	cancelCtx context.CancelFunc
	mtx       sync.Mutex
	coal      miners.Coal

	inventory    map[string]bool
	activeMiners []miners.Miner
	historyStats map[string]int
}

func RunGame() *Company {
	ctx, cancelContext := context.WithCancel(context.Background())
	company := Company{
		ctx:       ctx,
		cancelCtx: cancelContext,
		mtx:       sync.Mutex{},
		coal:      miners.Coal(0),

		inventory:    make(map[string]bool),
		historyStats: make(map[string]int),
		activeMiners: make([]miners.Miner, 0),
	}
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				company.mtx.Lock()
				company.coal++
				company.mtx.Unlock()
			}
		}
	}()
	return &company
}
func (c *Company) BuyMiner(minerType string) error {
	var miner miners.Miner
	switch minerType {
	case "basic":
		miner = miners.NewBasicMiner()
	case "normal":
		miner = miners.NewNormalMiner()
	case "advanced":
		miner = miners.NewAdvancedMiner()
	default:
		return ErrUnknownMinerType
	}
	cost := miner.Info().Salary

	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.coal < cost {
		return ErrInsufficientCoals
	}
	c.coal -= cost
	c.historyStats[minerType]++
	fmt.Println("Была соверена покупка")
	outChan := miner.Run(c.ctx)
	c.activeMiners = append(c.activeMiners, miner)

	go c.listenerMiner(outChan, miner)

	return nil
}

func (c *Company) listenerMiner(ch <-chan miners.Coal, whoIsIt miners.Miner) {
	for coal := range ch {
		c.AddCoal(coal)
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for i, m := range c.activeMiners {
		if m == whoIsIt {
			c.activeMiners[i] = nil
			c.activeMiners = append(c.activeMiners[:i], c.activeMiners[i+1:]...)
			break
		}
	}
}

func (c *Company) AddCoal(amount miners.Coal) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.coal += amount
}

func (c *Company) BuyProduct(productName string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.inventory[productName] {
		return ErrProductAlreadyBuyed
	}

	var foundProduct *equipment.Product

	for _, p := range equipment.AllProducts {
		if p.Name == productName {
			foundProduct = &p
			break
		}
	}

	if foundProduct == nil {
		return ErrUnknownProductName
	}
	cost := foundProduct.Cost
	if c.coal < cost {
		return ErrInsufficientCoals
	}
	c.coal -= cost
	c.inventory[productName] = true

	return nil
}

func (c *Company) checkWin() bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	hasPickaxe := c.inventory["pickaxe"]
	hasVent := c.inventory["vent"]
	hasWagon := c.inventory["wagon"]

	return hasPickaxe && hasVent && hasWagon
}

func (c *Company) EndGame() (GameStats, error) {
	if !c.checkWin() {
		return GameStats{}, ErrDontCompletedConditions
	}
	c.cancelCtx()
	return c.GetStats(), nil
}
