package main

import (
	"fmt"
	"fmBasics/imports"
)

func main() {
	fmt.Println("Hello World")

	ticket := imports.Ticket{ID: 1, Event: "Standup"}

	ticket.PrintEvent()
}
