package main

import "fmt"

func main() {
	var myName = "Arnawa Juan"

	fmt.Println("My Name is", myName)

	var name string = "Katty"
	fmt.Println("name = ",name)

	username := "admin"
	fmt.Println("username = ",username)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("Part 1 is", part1, "And Other is", other)
	
	// cann reuse other with :=
	part2, other := 2, 0
	fmt.Println("Part 2 is", part2, "And Other is", other)

	sum = part1 + part2
	fmt.Println("Sum =", sum)
	
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	
	fmt.Println("Lesson Name", lessonName)
	fmt.Println("Lesson Type", lessonType)

	// _ will skip it
	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

	// Exercise
	var color = "black"
	birthYear, age := 2000, 22
	var(

		firstName = "Arnawa Juan"  
		lastName = "Ibnuaji"
	)
	var ageDays int
	ageDays = 365 * age

	fmt.Println("My Name is", firstName, lastName, "Birth in", birthYear, "Age", age,"or in days is", ageDays, "days. My Favorite Color is", color)
}
