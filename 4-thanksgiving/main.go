package main

import (
	"fmt"
	"time"
)

type probability struct {
	twentytwo   int
	twentythree int
	twentyfour  int
	twentyfive  int
	twentysix   int
	twentyseven int
	twentyeight int
}

func main() {
	p := probability{}
	beginYear := 0
	endYear := 2017

	for i := beginYear; i <= endYear; i++ {
		dateOfThanksgiving(i, &p)
	}

	fmt.Println()
	fmt.Println("Occurrence of Thanksgiving on each day between year ", beginYear, " and ", endYear)
	p.prettyPrint()
}

func (p probability) prettyPrint() {
	space := ""

	fmt.Println("date:  22   23   24   25   26   27   28")
	fmt.Println("occ:  ", p.twentytwo, space, p.twentythree, space, p.twentyfour, space, p.twentyfive, space, p.twentysix, space, p.twentyseven, space, p.twentyeight)
}

// Given a year, this function will return the day of the month that
// Thanksgiving fell on for that year.  Thanksgiving falls on the 4th
// Thursday of every November.  Thus, the response will be an int
// somewhere between 22 and 28.
func dateOfThanksgiving(year int, p *probability) {
	novOneDay := time.Date(year, time.November, 1, 0, 0, 0, 0, time.UTC).Weekday().String()
	switch novOneDay {
	case "Sunday":
		p.twentysix = p.twentytwo + 1
		return
	case "Monday":
		p.twentyfive = p.twentytwo + 1
		return
	case "Tuesday":
		p.twentyfour = p.twentytwo + 1
		return
	case "Wednesday":
		p.twentythree = p.twentythree + 1
		return
	case "Thursday":
		p.twentytwo = p.twentytwo + 1
		return
	case "Friday":
		p.twentyeight = p.twentytwo + 1
		return
	case "Saturday":
		p.twentyseven = p.twentytwo + 1
		return
	}
	return
}
