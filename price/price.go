package price

import (
	"encoding/csv"
	"os"
)

type SpecialPrice struct {
	Quantity int
	Price    int
}

type Item struct {
	UnitPrice    int
	SpecialPrice SpecialPrice
}

func ReadPrice() {
    f, err := os.OpenFile("prices.csv", os.O_RDONLY, 0600)
    if err != nil {
        return nil, err
    }
}
