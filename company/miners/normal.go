package miners

import (
	"context"
	"time"
)

type NormalMiner struct {
	Stats MinerStats
}

func NewNormalMiner() *NormalMiner {
	return &NormalMiner{
		Stats: MinerStats{
			Class:           "normal",
			Salary:          50,
			Energy:          45,
			CoalPerTick:     3,
			Cooldown:        2,
			MoreCoalPerMine: 0,
		},
	}
}

func (b *NormalMiner) Run(ctx context.Context) <-chan Coal {
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

func (b *NormalMiner) Info() MinerStats {
	return b.Stats
}
