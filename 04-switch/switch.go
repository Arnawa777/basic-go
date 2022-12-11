package main

import "fmt"
 
func price() int {
	return 19
}

const (
	Economy		= 0
	Business	= 1
	FirstClass	= 2
)

func yourAge() uint {
	return 9
}

func main() {
	fmt.Println("Switch")

	// Fallthrough = execution next case
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap Item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")
	default:
		fmt.Println("Other seating")
	}

	switch age := yourAge(); {
	case age == 0:
		fmt.Println("You are Newborn")
	case age < 4:
		fmt.Println("You are Toddler")
	case age < 13:
		fmt.Println("You are Child")
	case age < 18:
		fmt.Println("You are Teenager")
	default:
		fmt.Println("You are Adult")
	}
}