package main

import (
	"fmt"
	"time"
)

func main() {
	then := time.Date(2021, 10, 17, 22, 31, 33, 651387237, time.UTC)
	fmt.Println(time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		time.Now().Nanosecond(),
		time.Now().Location(),
		time.Now().Weekday())

	fmt.Println(time.Now().Before(time.Now()),
		time.Now().After(time.Now()),
		time.Now().Equal(time.Now()))

	diff := time.Now().Sub(then)

	fmt.Println(diff.Hours(),
		diff.Minutes(),
		diff.Seconds(),
		diff.Nanoseconds())

	//  `Add` 将时间后移一个时间间隔
	//  `-` 来将时间前移一个时间间隔
	fmt.Println(then.Add(diff), then.Add(-diff))
}
