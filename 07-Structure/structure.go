package main

import (
	"fmt"
)

type passenger struct {
	name string
	ticketNumber int
	boarded bool
}

type insideBus struct {
	// Alias
	frontSeat passenger
}

func main() {
	type sample struct {
		field string
		a, b int
	}

	/*
	!default value
	/ data := sample {}

	!full name and value
	data := sample {
		field : "sad", 
		a: 2, 
		b: 3}

	!just add one value all will get default value
	data := sample{a: 5}


	*/

	data := sample {"sad", 2, 3}
	fmt.Println(data.field)
	
	data.field = "Happy"

	fmt.Println(data.field)

	//? Anonymous Structure

	// Must Add Value
	sample2 := struct{
		field2 string
		x, y int
	}{
		"hello", 2, 5,
	}

	fmt.Println(sample2.field2)
	
	// With Default Value
	var sample3 struct{
		field3 string
		d, e int
	}

	sample3.field3 = "Waow"
	
	fmt.Println(sample3)


	casey := passenger{"Casey", 1, false}
	fmt.Println(casey)

	// Ez Way for input many data
	var(
		bill = passenger{name: "Bill", ticketNumber: 2}
		ella = passenger{name: "Ella", ticketNumber: 3}
	)
	fmt.Println(bill, ella)

	casey.boarded = true
	bill.boarded = true

	if casey.boarded{
		fmt.Println(casey.name, "has Boarded the Bus")
	}

	bus := insideBus{bill}
	fmt.Println(bus)
	fmt.Println(bus.frontSeat.name, "is in the FrontSeat")
}