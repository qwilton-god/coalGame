package miners

import (
	"context"
)

type Coal int

type Miner interface {
	Run(ctx context.Context) <-chan Coal
	Info() MinerStats
}

type MinerStats struct {
	Class           string
	Salary          Coal
	Energy          int
	CoalPerTick     int
	Cooldown        int
	MoreCoalPerMine int
}
