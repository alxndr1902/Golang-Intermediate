package main

import (
	"fmt"
	"time"
)

func main() {
	// current local time
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2025, time.December, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	// parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") // Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	// formatting time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("06-01-02 15:04"))

	oneDayLatter := t.Add(time.Hour * 24)
	fmt.Println(oneDayLatter)
	fmt.Println(oneDayLatter.Weekday())

	fmt.Println("Rounded time", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Jakarta")
	t = time.Date(2025, time.December, 25, 12, 30, 15, 0, time.UTC)

	// convert this to the specific time zone
	tLocal := t.In(loc)

	// perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original time (UTC):", t)
	fmt.Println("Original time Local:", tLocal)
	fmt.Println("Rounded time (UTC):", roundedTime)
	fmt.Println("Rounded time Local:", roundedTimeLocal)

	// truncating time
	fmt.Println("Truncated time: ", t.Truncate(time.Hour))

	// time in new york
	new_york, _ := time.LoadLocation("America/New_York")
	timeInNewYork := time.Now().In(new_york)
	fmt.Println("Time in new york:", timeInNewYork)

	t1 := time.Date(2025, time.December, 25, 12, 30, 15, 0, time.UTC)
	t2 := time.Date(2025, time.December, 25, 18, 30, 15, 0, time.UTC)

	duration := t2.Sub(t1)
	fmt.Println(duration)

	// compare times
	fmt.Println("t2 is after t1?", t2.After(t1))
}
