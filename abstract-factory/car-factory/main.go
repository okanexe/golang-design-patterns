package main

import "fmt"

// Abstract Product
type ISedan interface {
	getModel() string
	setModel(string)
	getPower() int
	setPower(int)
}

type sedan struct {
	model string
	power int
}

func (s *sedan) getModel() string {
	return s.model
}

func (s *sedan) setModel(m string) {
	s.model = m
}

func (s *sedan) getPower() int {
	return s.power
}

func (s *sedan) setPower(p int) {
	s.power = p
}

// Abstract Product
type ISUV interface {
	getModel() string
	setModel(string)
	getPower() int
	setPower(int)
	isFourWheelDrive() bool
}

type suv struct {
	model          string
	power          int
	fourWheelDrive bool
}

func (s *suv) getModel() string {
	return s.model
}

func (s *suv) setModel(m string) {
	s.model = m
}

func (s *suv) getPower() int {
	return s.power
}

func (s *suv) setPower(p int) {
	s.power = p
}

func (s suv) isFourWheelDrive() bool {
	return s.fourWheelDrive
}

type ICarFactory interface {
	makeSedan() ISedan
	makeSUV() ISUV
}

// Concrete Products
type BMWSedan struct {
	sedan
}

type BMWSUV struct {
	suv
}

// Concrete Factory
type bmwFactory struct{}

func (bmw *bmwFactory) makeSedan() ISedan {
	return &BMWSedan{
		sedan: sedan{
			model: "3.20",
			power: 200,
		},
	}
}

func (bmw *bmwFactory) makeSUV() ISUV {
	return &BMWSUV{
		suv: suv{
			model: "X5",
			power: 250,
		},
	}
}

// Concrete Products
type mercedesSedan struct {
	sedan
}

type mercedesSUV struct {
	suv
}

// Concrete Factory
type mercedesFactory struct{}

func (m mercedesFactory) makeSedan() ISedan {
	return &mercedesSedan{
		sedan: sedan{
			model: "A180",
			power: 100,
		},
	}
}

func (m mercedesFactory) makeSUV() ISUV {
	return &mercedesSUV{
		suv: suv{
			model: "GLA200",
			power: 400,
		},
	}
}

// Factory Pattern
func GetSportFactory(brand string) (ICarFactory, error) {
	switch brand {
	case "bmw":
		return &bmwFactory{}, nil
	case "mercedes":
		return &mercedesFactory{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed ")
	}
}

func main() {
	bmwFactory, _ := GetSportFactory("bmw")
	mercedesFactory, _ := GetSportFactory("mercedes")

	sedan1 := bmwFactory.makeSedan()
	suv1 := bmwFactory.makeSUV()

	printSedanDetails(sedan1)
	printSUVDetails(suv1)

	sedan2 := mercedesFactory.makeSedan()
	suv2 := mercedesFactory.makeSUV()

	printSedanDetails(sedan2)
	printSUVDetails(suv2)
}

func printSedanDetails(s ISedan) {
	fmt.Printf("Model: %s", s.getModel())
	fmt.Println()
	fmt.Printf("Power: %d", s.getPower())
	fmt.Println()
}

func printSUVDetails(s ISUV) {
	fmt.Printf("Model: %s", s.getModel())
	fmt.Println()
	fmt.Printf("Power: %d", s.getPower())
	fmt.Println()
}
