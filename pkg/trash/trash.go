package trash

import (
	"fmt"
	"time"
)

// Pickup is from the types defined in: https://www.yonkersny.gov/home/showpublisheddocument/30455/637743931515170000
type Pickup int
const (
	UNDEFINED         Pickup = iota
	NO_PICKUP         Pickup = iota
	NO_PICKUP_HOLIDAY Pickup = iota
	TRASH             Pickup = iota
	PAPER             Pickup = iota
	COMMINGLES        Pickup = iota
)

func FromString(s string) Pickup {
	return map[string]Pickup{
		"cardboard":  PAPER,
		"commingles": COMMINGLES,
		"trash":      TRASH,
	}[s]
}

func (p Pickup) String() string {
	switch p {
	case UNDEFINED:
		return "<Undefined>"
	case NO_PICKUP:
		return "No Pickup"
	case NO_PICKUP_HOLIDAY:
		return "No Pickup because of a holiday"
	case TRASH:
		return "Trash"
	case PAPER:
		return "Paper"
	case COMMINGLES:
		return "Commingles"
	}
	return "unknown"
}

type simpleTime struct {
	year  int
	month time.Month
	day   int
}

var special = map[simpleTime]Pickup{
	{2022, time.May, 25}:       PAPER,
	{2022, time.May, 30}:       NO_PICKUP_HOLIDAY,
	{2022, time.May, 31}:       TRASH,
	{2022, time.June, 1}:       TRASH,
	{2022, time.June, 2}:       COMMINGLES,
	{2022, time.June, 3}:       TRASH,
	{2022, time.June, 8}:       PAPER,
	{2022, time.June, 15}:      COMMINGLES,
	{2022, time.June, 22}:      PAPER,
	{2022, time.June, 29}:      COMMINGLES,
	{2022, time.July, 4}:       NO_PICKUP_HOLIDAY,
	{2022, time.July, 5}:       TRASH,
	{2022, time.July, 6}:       NO_PICKUP_HOLIDAY,
	{2022, time.July, 7}:       PAPER,
	{2022, time.July, 8}:       NO_PICKUP_HOLIDAY,
	{2022, time.July, 13}:      COMMINGLES,
	{2022, time.July, 20}:      PAPER,
	{2022, time.July, 27}:      COMMINGLES,
	{2022, time.August, 3}:     PAPER,
	{2022, time.August, 10}:    COMMINGLES,
	{2022, time.August, 17}:    PAPER,
	{2022, time.August, 24}:    COMMINGLES,
	{2022, time.August, 31}:    PAPER,
	{2022, time.September, 5}:  NO_PICKUP_HOLIDAY,
	{2022, time.September, 6}:  TRASH,
	{2022, time.September, 7}:  NO_PICKUP_HOLIDAY,
	{2022, time.September, 8}:  COMMINGLES,
	{2022, time.September, 9}:  TRASH,
	{2022, time.September, 14}: PAPER,
	{2022, time.September, 21}: COMMINGLES,
	{2022, time.September, 28}: PAPER,
	{2022, time.October, 5}:    COMMINGLES,
	{2022, time.October, 10}:   NO_PICKUP_HOLIDAY,
	{2022, time.October, 11}:   TRASH,
	{2022, time.October, 12}:   NO_PICKUP_HOLIDAY,
	{2022, time.October, 13}:   PAPER,
	{2022, time.October, 14}:   TRASH,
	{2022, time.October, 19}:   COMMINGLES,
	{2022, time.October, 26}:   PAPER,
	{2022, time.November, 2}:   COMMINGLES,
	{2022, time.November, 3}:   TRASH,
	{2022, time.November, 9}:   PAPER,
	{2022, time.November, 11}:  NO_PICKUP_HOLIDAY,
	{2022, time.November, 16}:  COMMINGLES,
	{2022, time.November, 23}:  PAPER,
	{2022, time.November, 24}:  NO_PICKUP_HOLIDAY,
	{2022, time.November, 25}:  TRASH,
	{2022, time.November, 30}:  COMMINGLES,
	{2022, time.December, 7}:   PAPER,
	{2022, time.December, 14}:  COMMINGLES,
	{2022, time.December, 21}:  PAPER,
	{2022, time.December, 26}:  NO_PICKUP_HOLIDAY,
	{2022, time.December, 27}:  TRASH,
	{2022, time.December, 28}:  NO_PICKUP_HOLIDAY,
	{2022, time.December, 29}:  COMMINGLES,
	{2022, time.December, 30}:  TRASH,
	{2023, time.January, 2}: NO_PICKUP_HOLIDAY,
	{2023, time.January, 3}: TRASH,
	{2023, time.January, 5}: PAPER,
	{2023, time.January, 6}: TRASH,
	{2023, time.January, 11}: COMMINGLES,
	{2023, time.January, 16}: NO_PICKUP_HOLIDAY,
	{2023, time.January, 17}: TRASH,
	{2023, time.January, 19}: PAPER,
	{2023, time.January, 20}: TRASH,
	{2023, time.January, 25}: COMMINGLES,
}

func ForDate(when time.Time) Pickup {
	p, found := special[simpleTime{year: when.Year(), month: when.Month(), day: when.Day()}]
	if found {
		return p
	}
	if when.Weekday() == time.Monday || when.Weekday() == time.Thursday {
		return TRASH
	}
	return NO_PICKUP
}

func NextPickup(starting time.Time, pickup Pickup) (time.Time, error) {
	for i := 0; i < 14; i++ {
		starting = starting.Add(time.Hour * 24)
		if ForDate(starting) == pickup {
			return starting, nil
		}
	}
	return time.Time{}, fmt.Errorf("No upcoming %s pickup.", pickup)
}
