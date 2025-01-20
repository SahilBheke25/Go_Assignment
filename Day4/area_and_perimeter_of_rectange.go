package main

/*
2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

Hint:
Method Signatures for Area and Perimeter
Area() int
Perimeter() int

Formulae:
Area = length * width
Perimeter = 2 * (length + width)

Example Test Case 1:
Input: 10 20
Output:
Area of Rectangle: 200
Perimeter of Rectangle: 60
*/

import "fmt"

type Rectangle struct {
	length    int
	breadth   int
	area      int
	perimeter int
}

func (rec Rectangle) Area() int {
	return rec.length * rec.breadth
}

func (rec Rectangle) Perimeter() int {
	return 2 * (rec.length + rec.breadth)
}

func main() {

	var length, breadth int
	fmt.Println("Enter the Length and Breadth of Rectangle:")
	_, err := fmt.Scanf("%v %v", &length, &breadth)

	// input validation
	if err != nil || length < 1 || breadth < 1 {
		fmt.Println("Invalid input. Please enter integer value!!!  Hint=", err)
		return
	}

	rec := Rectangle{}
	rec.length = length
	rec.breadth = breadth
	rec.area = rec.Area()
	rec.perimeter = rec.Perimeter()

	fmt.Println("Area of the rectangle is:", rec.area)
	fmt.Println("Perimeter of the rectangle is:", rec.perimeter)
}
