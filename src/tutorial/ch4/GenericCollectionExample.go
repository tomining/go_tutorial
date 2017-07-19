package main

import (
	"fmt"
	"strings"
)

func main() {
	shoes := Item{"Sport Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 2},
		10.00,
	}

	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}

	// Generic Collection
	items := Items{shoes, eventShoes, shirt, video}
	displayCost(items)

	fmt.Println(shoes)
	fmt.Println(eventShoes)
	fmt.Println(shirt)
	fmt.Println(video)
	fmt.Println(items)
}

// Generic Collection
type Items []Coster

func (ts Items) cost() (c float64) {
	for _, t := range ts {
		c += t.cost()
	}
	return
}
// Generic Collection

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
	return r.feePerDay * float64(r.ToDays()*r.periodLength)
}

// Java의 ToString 개념인 듯
type String interface {
	String() string
}

func (t Item) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.cost())
}

func (t DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", t.Item.String(), t.cost(), t.discountRate, "%")
}

func (t Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.cost())
}

func (ts Items) String() string {
	var s []string
	for _, t := range ts {
		s = append(s, fmt.Sprint(t))
	}
	return fmt.Sprintf("%d items. total: %.0f\n\t-%s", len(ts), ts.cost(), strings.Join(s, "\n\t- "))
}