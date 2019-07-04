package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {

		fmt.Println("1")
		time.Sleep(5 * time.Second) // simulate a page download op.
		fmt.Println("Go routine done.")
		done <- true
	}()

	fmt.Println("2")
	<-done
	fmt.Print("Program ends.")
}
