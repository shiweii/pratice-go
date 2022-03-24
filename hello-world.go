package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println("Time now is", t.Format("2006-01-02 15:04:05"))
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
