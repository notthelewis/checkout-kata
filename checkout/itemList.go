package checkout

type SpecialPrice struct {
	Quantity uint
	Price    int
}

type Item struct {
	UnitPrice    int
	SpecialPrice SpecialPrice
}

var allItems map[rune]Item = make(map[rune]Item)

func init() {
	allItems['A'] = Item{
		UnitPrice: 50,
		SpecialPrice: SpecialPrice{
			Quantity: 3,
			Price:    130,
		},
	}

	allItems['B'] = Item{
		UnitPrice: 30,
		SpecialPrice: SpecialPrice{
			Quantity: 2,
			Price:    45,
		},
	}

	allItems['C'] = Item{
		UnitPrice: 20,
	}

	allItems['D'] = Item{
		UnitPrice: 15,
	}
}
