package miners

import (
	"context"
	"time"
)

type BasicMiner struct {
	Stats MinerStats
}

func NewBasicMiner() *BasicMiner {
	return &BasicMiner{
		Stats: MinerStats{
			Class:           "basic",
			Salary:          5,
			Energy:          30,
			CoalPerTick:     1,
			Cooldown:        3,
			MoreCoalPerMine: 0,
		},
	}
}

func (b *BasicMiner) Run(ctx context.Context) <-chan Coal {
	out := make(chan Coal)
	go func() {
		defer close(out)
		ticker := time.NewTicker(time.Duration(b.Stats.Cooldown) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if b.Stats.Energy <= 0 {
					return
				}
				out <- Coal(b.Stats.CoalPerTick)

				b.Stats.CoalPerTick += b.Stats.MoreCoalPerMine
				b.Stats.Energy--
			}
		}
	}()
	return out
}

func (b *BasicMiner) Info() MinerStats {
	return b.Stats
}
