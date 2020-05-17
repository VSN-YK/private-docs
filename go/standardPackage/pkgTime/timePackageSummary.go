package pkgTime

import (
	"fmt"
	"time"
)

func TimePackageSummary() {
	now := time.Now()
	fmt.Printf("%v\nYear:%v\nMonth:%v\n%T\n", now, now.Year(), now.Month(), now.Weekday())

	fmt.Printf("%s\n", time.Tuesday.String())

	//time duration const index of summary
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	//duration convert
	h, _ := time.ParseDuration("12h15m30s")
	fmt.Printf("Key of Hours\t%f\nKey of Minutes\t%f\n", h.Hours(), h.Minutes())

	//Example in Add
	fmt.Println(now)
	fmt.Println(now.Add(2*time.Hour + 30*time.Minute))

	//Example in sub (next FIFA)
	fifa := time.Date(2022, 11, 21, 0, 0, 0, 0, time.Local)
	fmt.Printf("%v\n", fifa)
	fmt.Printf("%v\n", fifa.Sub(now))
	fmt.Printf("%v", now.Format("2006/01/02"))

	//TODO : use context
	for range time.Tick(1 * time.Minute) {
		fmt.Println("Tick!!")
	}
}
