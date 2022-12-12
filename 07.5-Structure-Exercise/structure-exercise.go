package main

import (
	"fmt"
)

type coordinate struct {
	x, y int
}

type rectangle struct {
	a coordinate //top left
	b coordinate //bottom right
}



func width(rect rectangle) int {
	return (rect.b.x - rect.a.x)
}

func lenght(rect rectangle) int {
	return (rect.a.y - rect.b.y)
}

func area(rect rectangle) int {
	return lenght(rect) * width(rect)
}

func perimeter(rect rectangle) int {
	return (width(rect) * 2) + (lenght(rect) * 2)
}

func printInfo(rect rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}
func main() {

	persegi := rectangle{a: coordinate{0, 7}, b:coordinate{10, 0}}
	fmt.Println(persegi.a.x, persegi.a.y)

	fmt.Println("Width", width(persegi))
	fmt.Println("Lenght", lenght(persegi))
	printInfo(persegi)
	
	persegi.a.y *= 2
	persegi.b.x *= 2
	printInfo(persegi)
	
	// fmt.Println(coordinate)
}