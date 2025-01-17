package main

import (
	"fmt"
	"math"
	"strconv"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {

	// Assignment Discription
	// Program to calculate the area of the circle, First line has a value of the radius of the circle
	// constraint
	// 1. Use const PI from the package math package
	// 2. Use the Pow function from the math package
	// 3. Round area float value to 2 decimal places

	var inputString string
	var radius float64
	var area float64

	// Taking Radius as input form user
	fmt.Println("Enter the radius of circle:")
	fmt.Scan(&inputString)

	radius, err := strconv.ParseFloat(inputString, 64)

	if err != nil {
		fmt.Println("Enter valid input radius")
		return
	}
	// Calculation
	area = radius * radius * math.Pi

	// Round of to 2 decimal places
	area = roundFloat(area, 2)

	fmt.Println(area)
}
