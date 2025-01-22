package main

import (
	"fmt"
)

/*
3. The given program takes an integer value as input from the user i.e 1 or 2.
Option 1 represents Rectangle and Option 2 represents Square.

Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:

1. Create an interface Quadrilateral which has the following method signatures
Area() int
Perimeter() int

2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square

3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

"Area :  <value>"
"Perimeter :  <value>"


HINT: Step 2 means, to define methods in Quadrilateral interface for given shapes


Formulae:

Area of Rectangle: length * width
Perimeter of Rectangle: 2*(length + breadth)


Area of Square: side * side
Perimeter of Square: 4 * side
*/

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

// Rectangle data type and its functions
type Rectangle struct {
	length  int
	breadth int
}

func (rec Rectangle) Area() int {
	return rec.length * rec.breadth
}
func (rec Rectangle) Perimeter() int {
	return 2 * (rec.length + rec.breadth)
}

// Square data type and its functions
type Square struct {
	side int
}

func (square Square) Area() int {
	return square.side * square.side
}

func (square Square) Perimeter() int {
	return 4 * square.side
}

// Printing the area and perimeter of the quadrilateral chossen by the user.
func print(quadrilateral Quadrilateral) {
	fmt.Println("Area:", quadrilateral.Area())
	fmt.Println("Perimeter:", quadrilateral.Perimeter())
}
func main() {

	var option int
	fmt.Println("Enter 1 for RECTANGLE and 2 for SQUARE")
	_, err := fmt.Scanf("%d", &option)

	if err != nil || (option != 1 && option != 2) {
		fmt.Println("Invalid Input")
		return
	} else if option == 1 { // Rectangle Calculation.
		rectangle := Rectangle{
			length:  10,
			breadth: 20,
		}
		print(rectangle)
	} else { // Square Calculation.
		square := Square{
			side: 2,
		}
		print(square)
	}
}
