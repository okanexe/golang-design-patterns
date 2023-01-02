package main

import "fmt"

type ICar interface {
	setModel(m string)
	setPower(p int)
	getModel() string
	getPower() int
}

type Car struct {
	model string
	power int
}

func (c *Car) setModel(m string) {
	c.model = m
}

func (c *Car) getModel() string {
	return c.model
}

func (c *Car) setPower(p int) {
	c.power = p
}

func (c *Car) getPower() int {
	return c.power
}

type bmw struct {
	Car
}

func newBmw() ICar {
	return &bmw{
		Car: Car{
			model: "3.20",
			power: 150,
		},
	}
}

type mercedes struct {
	Car
}

func newMercedes() ICar {
	return &mercedes{
		Car: Car{
			model: "A180",
			power: 100,
		},
	}
}

func getCar(carModel string) (ICar, error) {
	switch carModel {
	case "3.20":
		return newBmw(), nil
	case "A180":
		return newMercedes(), nil
	}
	return nil, fmt.Errorf("wrong car model passed ")
}

func main() {
	bmw320, _ := getCar("3.20")
	a180, _ := getCar("A180")

	fmt.Println("Car Model: ", bmw320.getModel())
	fmt.Println("Power: ", bmw320.getPower())

	fmt.Println("Car Model: ", a180.getModel())
	fmt.Println("Power: ", a180.getPower())
}
