package main

import (
	"fmt"
	"time"
)

func main() {
	go waitAndPrint("Hello")
	go waitAndPrint("Hello 2")
	go waitAndPrint("Hello 3")
	printAndWait("Bye")
}

func waitAndPrint(msg string) {
	time.Sleep(time.Second)
	fmt.Println(msg)
}

func printAndWait(msg string) {
	fmt.Println(msg)
	time.Sleep(2 * time.Second)
}
