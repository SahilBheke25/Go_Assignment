package main

import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {

	var inputString string
	fmt.Println("To calculate Simple Intrest Please Enter principle, rate, time(in years) as a comma separated values: ")
	fmt.Scanln(&inputString)

	var principle, rate, time, simple_intrest, result float64

	_, err := fmt.Sscanf(inputString, "%v,%v,%v", &principle, &rate, &time)

	if err == nil {
		simple_intrest = (principle * rate * time) / 100

		result = roundFloat(simple_intrest, 2)

		fmt.Println("Simple Intrest is:", result)
	} else {
		fmt.Println("Please provide input properly.")
	}
}
