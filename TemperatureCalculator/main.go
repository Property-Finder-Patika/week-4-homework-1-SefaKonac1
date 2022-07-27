package main

import "fmt"

/*degrees struct store different temperature measurement units*/
type Degrees struct {
	Celcius    float32
	Fahrenheit float32
	Kelvin     float32
}

/*an interface was defined to do convertion operations between different temperature units*/
type ConvertionMethods interface {
	Kelvin2Fahrenheit()
	Kelvin2Celsius()
	Fahrenheit2Celsius()
	Fahrenheit2Kelvin()
	Celsius2Fahrenheit()
	Celsius2Kelvin()
}

func (d Degrees) Kelvin2Fahrenheit() float32 {
	return (d.Kelvin*9/5 - 459.67)
}

func (d Degrees) Kelvin2Celsius() float32 {
	return (d.Kelvin - 273.15)
}

func (d Degrees) Fahrenheit2Celsius() float32 {
	return ((d.Fahrenheit - 32) * 5 / 9)
}

func (d Degrees) Fahrenheit2Kelvin() float32 {
	return ((d.Fahrenheit + 459.67) * 5 / 9)
}

func (d Degrees) Celsius2Fahrenheit() float32 {
	return (d.Celcius*9/5 + 32)
}

func (d Degrees) Celsius2Kelvin() float32 {
	return (d.Celcius + 273.15)
}

func main() {

	/*stores degrees struct object with arbitrary values*/
	testDegrees := Degrees{
		Celcius:    65,
		Kelvin:     72,
		Fahrenheit: -120,
	}

	fmt.Printf("%f celcius is equal to %f fahrenheit ", testDegrees.Celcius, testDegrees.Celsius2Fahrenheit())
	fmt.Printf("%f celcius is equal to %f kelvin \n", testDegrees.Celcius, testDegrees.Celsius2Kelvin())

	fmt.Printf("%f kelvin is equal to %f fahrenheit ", testDegrees.Kelvin, testDegrees.Kelvin2Fahrenheit())
	fmt.Printf("%f kelvin is equal to %f celcius \n", testDegrees.Kelvin, testDegrees.Kelvin2Celsius())

	fmt.Printf("%f fahrenheit is equal to %f celcius ", testDegrees.Fahrenheit, testDegrees.Fahrenheit2Celsius())
	fmt.Printf("%f fahrenheit is equal to %f kelvin \n", testDegrees.Fahrenheit, testDegrees.Fahrenheit2Kelvin())

}
