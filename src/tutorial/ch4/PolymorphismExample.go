package main

import "fmt"

func main() {
	shoes := Item{"Sport Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 2},
		10.00,
	}

	displayCost(shoes)
	displayCost(eventShoes)

	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}

	displayCost(shirt)
	displayCost(video)
}

type Coster interface {
	cost() float64
}

func displayCost(c Coster) {
	fmt.Println("Cost : ", c.cost())
}

type Item struct {
	name string
	price float64
	quantity int
}

func (t Item) cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) cost() float64 {
	return t.Item.cost() * (1 - t.discountRate/100)
}

type Rental struct {
	name string
	feePerDay float64
	periodLength int
	RentalPeriod
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

func (r Rental) cost() float64 {
	return r.feePerDay * float64(r.ToDays() * r.periodLength)
}