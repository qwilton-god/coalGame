package equipment

import "coalGame/company/miners"

type Product struct {
	Name string
	Cost miners.Coal
}

var AllProducts = []Product{
	{
		Name: "pickaxe",
		Cost: 3000,
	},
	{
		Name: "vent",
		Cost: 15_000,
	},
	{
		Name: "wagon",
		Cost: 50_000,
	},
}
