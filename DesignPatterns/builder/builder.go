package main

import "fmt"

type Wrapping struct{}
type Bottle struct{}

type Packing interface {
	pack() string
}

type Item interface {
	name() string
	Packing() Packing
	price() float64
}

type Burger struct {
	Item Item
}

func (p *Wrapping) pack() string {
	return "Wrapper"
}
func (p *Burger) Packing() Packing {
	return &Wrapping{}
}

type VegBurger struct {
	Burger Burger
}

// Packing implements Item.
func (*VegBurger) Packing() Packing {
	return &Wrapping{}
}

func (p *VegBurger) name() string {
	return "Veg Burger"
}

func (p *VegBurger) price() float64 {
	return 25.0
}

type ChickenBurger struct {
	Burger Burger
}

// Packing implements Item.
func (*ChickenBurger) Packing() Packing {
	return &Wrapping{}
}

func (p *ChickenBurger) name() string {
	return "Chicken Burger"
}

func (p *ChickenBurger) price() float64 {
	return 50.0
}

type ColdDrink struct {
	Item Item
}

func (p *Bottle) pack() string {
	return "Bottle"
}

func (p *ColdDrink) Packing() Packing {
	return &Bottle{}
}

type Pepsi struct {
	ColdDrink ColdDrink
}

// Packing implements Item.
func (*Pepsi) Packing() Packing {
	return &Bottle{}
}

func (p *Pepsi) name() string {
	return "Pepsi"
}

func (p *Pepsi) price() float64 {
	return 30.0
}

type Coke struct {
	ColdDrink ColdDrink
}

// Packing implements Item.
func (*Coke) Packing() Packing {
	return &Bottle{}
}

func (p *Coke) name() string {
	return "Coke"
}

func (p *Coke) price() float64 {
	return 35.0
}

type Meal struct {
	Items []Item
}

func (p *Meal) addItem(item Item) {
	p.Items = append(p.Items, item)
}

func (p *Meal) getCost() float64 {
	var cost float64 = 0.0
	for _, item := range p.Items {
		cost += item.price()
	}
	return cost
}

func (p *Meal) showItems() {
	for _, item := range p.Items {
		println("Item: ", item.name())
		println("Packing: ", item.Packing().pack())
		fmt.Printf("Price: %.2f\n", item.price())
	}
}

type MealBuilder struct{}

func (p *MealBuilder) prepareVegMeal() *Meal {
	var meal Meal
	meal.addItem(&VegBurger{})
	meal.addItem(&Coke{})
	return &meal
}

func (p *MealBuilder) prepareNonVegMeal() *Meal {
	var meal Meal
	meal.addItem(&ChickenBurger{})
	meal.addItem(&Pepsi{})
	return &meal
}

func main() {
	var mealBuilder MealBuilder
	var vegMeal = mealBuilder.prepareVegMeal()
	println("Veg Meal")
	vegMeal.showItems()
	fmt.Printf("Total Cost: %.2f\n", vegMeal.getCost())

	var nonVegMeal = mealBuilder.prepareNonVegMeal()
	println("Non-Veg Meal")
	nonVegMeal.showItems()
	fmt.Printf("Total Cost: %.2f\n", nonVegMeal.getCost())
}
