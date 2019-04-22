package main

import (
	"fmt"
	"time"
)

func toDay(day time.Weekday)  int{
	today := time.Now().Weekday()
	switch day {
	case today+0:
		return 0
	case today+1:
		return 1
	case today+2:
		return 2
	case today+3:
		return 3
	case today+4:
		return 4
	case today+5:
		return 5
	case today+6:
		return 6
	default:
		return 0 
	}
}
func main() {
	fmt.Println(toDay(time.Saturday))
}
