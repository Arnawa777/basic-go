package main

import "fmt"

func double(x int) int{
	return x + x
}

func add(lhs, rhs int) int{
	return lhs + rhs
}

func greet(){
	fmt.Println("Hello from greet function!!")
}

func twoTwos() (int, int){
	return 2, 2
}

func five() int{
	return 5
}

func addThree(a, b, c int) int{
	return a + b + c
}

func greetPerson (name string){
	fmt.Println("Hello", name)
}

func hiThere() string{
	return "Hi, There"
}

func main () {
	greet()
	dozen := double(12)
	fmt.Println("A Dozen is", dozen)

	bakersDozen := add(dozen, 21)
	fmt.Println("A baker's dozen is", bakersDozen)

	greetPerson("Alice")
	fmt.Println(hiThere())

	a, b := twoTwos()
	answer := addThree(five(), a, b)

	fmt.Println(answer)
} 