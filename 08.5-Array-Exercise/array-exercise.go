package main

import (
	"fmt"
)

type product struct{
	price int
	name string
}

func printStats(list [4]product){
	var cost, totalItems int

	for i:=0; i<len(list); i++{
		item:= list[i]
		cost += item.price

		if item.name != "" {
			totalItems++
		}
	}

	fmt.Println("Last item of the list: ", list[totalItems - 1])
	fmt.Println("Total Items: ", totalItems)
	fmt.Println("Total Cost: ", cost)
}

func main(){
	shopList := [4]product{
		{1, "Banana"},
		{2, "Tomato"},
		{6, "Chicken"},
	}

	printStats(shopList)

	shopList[3] = product{10, "Meat"}
	printStats(shopList)
}