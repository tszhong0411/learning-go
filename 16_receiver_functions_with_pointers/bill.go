package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...$%v \n", k+":", v)
	}

	// add tip
	fs += fmt.Sprintf("%-15v ...%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-15v ...$%0.2f", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
