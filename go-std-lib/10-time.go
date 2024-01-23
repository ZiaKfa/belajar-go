package main

import (
	"fmt"
	"time"
)

func main() {
	//time is a package that implements time-related functions
	now := time.Now()
	fmt.Println(now)

	//time.Time is a struct that contains Year, Month, Day, Hour, Minute, Second, and Nanosecond
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	//time.Time implements Weekday() function that returns the day of the week
	fmt.Println(now.Weekday())

	//time.Time implements Add() function that returns a Time value that is duration d later than the current Time
	fmt.Println(now.Add(time.Hour * 24))

	//time.Time implements Sub() function that returns the duration between two times
	fmt.Println(now.Sub(now.Add(time.Hour * 24)))

	//time.Time implements Format() function that returns a textual representation of the time value formatted according to layout
	fmt.Println("Format() function")
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(now.Format("02/01/2006 03:04:05 PM"))
	fmt.Println(now.Format("02/01/2006 15:04:05"))

	//time.Time implements Unix() function that returns the local Time corresponding to the given Unix time, sec seconds and nsec nanoseconds since January 1, 1970 UTC
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Time implements UTC() function that returns a Time value representing the same moment in time with the same nanosecond precision, but with the UTC time zone offset
	fmt.Println(now.UTC())

	//time.Time implements Local() function that returns a Time value representing the same moment in time with the same nanosecond precision, but with the current local time zone offset
	fmt.Println(now.Local())

	//time.parse() function parses a formatted string and returns the time value it represents
	fmt.Println("Parse() function")
	format := "2006-Jan-02 15:04:05"
	timeStr := "2004-Jul-10 00:00:00"
	timenow, err := time.Parse(format, timeStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(timenow)
	}

	//duration is a type that represents a signed integer count of nanoseconds
	//time.Duration implements Hours(), Minutes(), Seconds(), and Nanoseconds() functions
	fmt.Println("Duration")
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 100 * time.Minute
	var duration3 time.Duration = duration1 + duration2
	fmt.Println(duration3)
}
