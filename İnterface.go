package main

import "fmt"

func main() {
	cars1 := new(Ferrari) // mercedes ve lamborghini içinde aynı.
	cars1.Brand = "Ferrari"
	cars1.Model = "F50"
	cars1.Color = "Red"
	cars1.Speed = 400
	cars1.Price = 999
	cars1.Special = true
	fmt.Println(cars1.Information())
	fmt.Println(cars1.Run())
	fmt.Println(cars1.Stop())

	carsData(cars1) // en alttaki fonksiyondan çalıştırdık.

	cars2 := new(Mercedes)
	cars2.Brand = "Mercedes"
	cars2.Model = "CLA"
	cars2.Color = "Black"
	cars2.Speed = 300
	cars2.Price = 500
	cars2.Special = true
	fmt.Println(cars2.Information())
	fmt.Println(cars2.Run())
	fmt.Println(cars2.Stop())

	carsData(cars2)

}

//Base struct
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

// Ferrari
type Ferrari struct { // Ferrari nesnesi Car nesnesini kendine kalıtım olarak aldı.
	Car //composition
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (f Ferrari) Information() string {
	ınf := f.Brand + f.Color + f.Model
	add := "yes"
	if f.Special {
		ınf += "special:" + add
	}
	return ınf
}

// Lamborghini
type Lamborghini struct {
	Car //composition
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (f Lamborghini) Information() string {
	ınf := f.Brand + f.Color + f.Model
	add := "yes"
	if f.Special {
		ınf += "special:" + add
	}
	return ınf
}

// Mercedes
type Mercedes struct {
	Car //composition
	SpecialProduction
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (m Mercedes) Information() string {
	ınf := m.Brand + m.Color + m.Model
	add := "no"
	if m.Special {
		ınf += "special:" + add
	}
	return ınf
}

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

func carsData(c Carface) {
	fmt.Println(c.Information())
	fmt.Println(c.Run())
	fmt.Println(c.Stop())
}
