package main

import (
	"fmt"
)

type roomStatus struct{
	name string
	cleaned bool
}

func checkCleaniness(rooms []roomStatus){
	for i:=0; i< len(rooms); i++ {
		room:= rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "Are Clean")
		}else{
			fmt.Println(room.name, "Are Dirty")
		}

	}
}

func main() {
	rooms:= []roomStatus{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
		{name: "OB"},
		{name: "Cleaner"},
	}

	checkCleaniness(rooms)

	fmt.Println("=====Do Cleaning=====")
	rooms[0].cleaned = true
	rooms[5].cleaned = true
	rooms[4].cleaned = true

	checkCleaniness(rooms)
}