package main

import "fmt"

func average (a, b, c int) float32{
	return float32(a+b+c) /3
}

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekdays(day int) bool{
	return day <=4
}

func main() {
	quiz1, quiz2, quiz3 := 9, 9, 9 
	avgQuiz := average(quiz1, quiz2, quiz3)

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 hebat")
	}else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 Menang bro")
	}else{
		fmt.Println("Quiz 1 equal with quiz 2 ")		
	}
	fmt.Println(avgQuiz)
	fmt.Println("===========Exercise===========")

	//--Requirements:
	//* Use the accessGranted() and accessDenied() functions to display
	//  informational messages
	//* Access at any time: Admin, Manager
	//* Access weekends: Contractor
	//* Access weekdays: Member
	//* Access Mondays, Wednesdays, and Fridays: Guest


	// The day and role. Change these to check your work.
	today, role := Saturday, Guest

	// fmt.Println(today, role)
	if role == 10 || role == 20 {
		//* Access at any time: Admin, Manager
		accessGranted()
	} else if role == Contractor && weekdays(today){
		//* Access weekends: Contractor
		accessGranted()
	} else if role == Member && !weekdays(today) {
		//* Access weekdays: Member
		accessGranted()
	} else if role == 50 && (today == Monday || today == Wednesday || today == Friday) {
		//* Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
	} else{
		accessDenied()
	}
}