package main

import "fmt"

func main() {
	for i := 0; i<10; i++{
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "So Sad")
	} 

	sum := 0
	for i := 1; i<=1; i++{
		sum += i
		fmt.Println("Sum is", sum)
	}
	
	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement Sum is", sum)
	}

	for i := 1; i<=50; i++{
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 {
			fmt.Println(i, "Fizz")
		}else if divisibleBy5 {
			fmt.Println(i, "Buzz")
		}else if divisibleBy3 && divisibleBy5{
			fmt.Println(i, "FizzBuzz")
		}else {
			fmt.Println(i)
				
		}
		
	} 
}