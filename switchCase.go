package main

import (
	"fmt"
	"time"
)

// SwitchGo :
func SwitchGo() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
