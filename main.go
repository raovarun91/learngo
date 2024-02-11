package main

import (
	"fmt"
	"time"
)

func main() {

	random := 1
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		random += 1
	}
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("Time taken %v", elapsed)
	//var username string
	//fmt.Print("Hello! Please enter your name:")
	//fmt.Scan(&username)
	//fmt.Printf("Hi %v now this is personalized", username)
}
