package miners

import (
	"context"
	"time"
)

type AdvancedMiner struct {
	Stats MinerStats
}

func NewAdvancedMiner() *AdvancedMiner {
	return &AdvancedMiner{
		Stats: MinerStats{
			Class:           "advanced",
			Salary:          450,
			Energy:          60,
			CoalPerTick:     10,
			Cooldown:        1,
			MoreCoalPerMine: 3,
		},
	}
}

func (b *AdvancedMiner) Run(ctx context.Context) <-chan Coal {
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

func (b *AdvancedMiner) Info() MinerStats {
	return b.Stats
}
