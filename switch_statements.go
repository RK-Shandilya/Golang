package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch_With_no_condition() {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("time is less than 12 hours")
	case t.Hour() > 12:
		fmt.Println("time is more than 12 hours")
	default:
		fmt.Println("time is equal to 12 hours")
	}
}

func findSaturday() {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow is Saturday")
	case today + 2:
		fmt.Println("After tomorrow is Saturday")
	default:
		fmt.Println("Too far away")
	}
}

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin", "windows":
		fmt.Println(os)
	case "linux":
		fmt.Println(os)
	default:
		fmt.Println("an unknown platform")
	}

	findSaturday()
	switch_With_no_condition()
}
