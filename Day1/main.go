package main

import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}


func calculateSimpleIntrest(){

	var inputString string;
	fmt.Println("To calculate Simple Intrest Please Enter principle, rate, time(in years) as a comma separated values: ")
	fmt.Scanln(&inputString)

	var principle, rate, time, simple_intrest, result float64
 
	_, err := fmt.Sscanf(inputString, "%v,%v,%v", &principle, &rate, &time)

	if(err == nil){
		simple_intrest = (principle * rate * time) / 100

		result = roundFloat(simple_intrest, 2)

		fmt.Println("Simple Intrest is:",result)
	}else{
		fmt.Println("Please provide input properly.")
	}

}

const PI float64 = 3.14

func areaOfCircle(){
// 	2. Program to calculate the area of the circle, First line has a value of the radius of the circle
// constraint
// 1. Use const PI from the package math package
// 2. Use the Pow function from the math package
// 3. Round area float value to 2 decimal places
	var radius, area float64;
	fmt.Println("Enter the radius of circle:")
	fmt.Scan(&radius)

	area = radius * radius * math.Pi

	fmt.Println(area)

}

func main(){
	
	var option int 
	fmt.Println("To Calculate Simple Intrest Enter 1 OR For Area Of Circle Enter 0")
	fmt.Scan(&option)
	if(option == 1){
		calculateSimpleIntrest()
	}else{
		areaOfCircle()
	}
}
