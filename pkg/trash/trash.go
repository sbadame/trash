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
	{2023, time.January, 2}:    NO_PICKUP_HOLIDAY,
	{2023, time.January, 3}:    TRASH,
	{2023, time.January, 5}:    PAPER,
	{2023, time.January, 6}:    TRASH,
	{2023, time.January, 11}:   COMMINGLES,
	{2023, time.January, 16}:   NO_PICKUP_HOLIDAY,
	{2023, time.January, 17}:   TRASH,
	{2023, time.January, 19}:   PAPER,
	{2023, time.January, 20}:   TRASH,
	{2023, time.January, 25}:   COMMINGLES,
	{2023, time.February, 1}:   PAPER,
	{2023, time.February, 8}:   COMMINGLES,
	{2023, time.February, 13}:  NO_PICKUP_HOLIDAY,
	{2023, time.February, 14}:  TRASH,
	{2023, time.February, 15}:  NO_PICKUP_HOLIDAY,
	{2023, time.February, 16}:  PAPER,
	{2023, time.February, 17}:  TRASH,
	{2023, time.February, 20}:  NO_PICKUP_HOLIDAY,
	{2023, time.February, 21}:  TRASH,
	{2023, time.February, 22}:  NO_PICKUP_HOLIDAY,
	{2023, time.February, 23}:  COMMINGLES,
	{2023, time.February, 24}:  TRASH,
	{2023, time.March, 1}:      PAPER,
	{2023, time.March, 8}:      COMMINGLES,
	{2023, time.March, 15}:     PAPER,
	{2023, time.March, 22}:     COMMINGLES,
	{2023, time.March, 29}:     PAPER,
	{2023, time.April, 5}:      COMMINGLES,
	{2023, time.April, 12}:     PAPER,
	{2023, time.April, 19}:     COMMINGLES,
	{2023, time.April, 26}:     PAPER,
	{2023, time.May, 3}:        COMMINGLES,
	{2023, time.May, 10}:       PAPER,
	{2023, time.May, 17}:       COMMINGLES,
	{2023, time.May, 24}:       PAPER,
	{2023, time.May, 29}:       NO_PICKUP_HOLIDAY,
	{2023, time.May, 30}:       TRASH,
	{2023, time.May, 31}:       NO_PICKUP_HOLIDAY,
	{2023, time.June, 1}:       COMMINGLES,
	{2023, time.June, 2}:       TRASH,
	{2023, time.June, 7}:       PAPER,
	{2023, time.June, 14}:      COMMINGLES,
	{2023, time.June, 19}:      NO_PICKUP_HOLIDAY,
	{2023, time.June, 20}:      TRASH,
	{2023, time.June, 21}:      NO_PICKUP_HOLIDAY,
	{2023, time.June, 22}:      PAPER,
	{2023, time.June, 23}:      TRASH,
	{2023, time.June, 28}:      COMMINGLES,
	{2023, time.July, 4}:       NO_PICKUP_HOLIDAY,
	{2023, time.July, 5}:       NO_PICKUP_HOLIDAY,
	{2023, time.July, 6}:       PAPER,
	{2023, time.July, 7}:       TRASH,
	{2023, time.July, 12}:      COMMINGLES,
	{2023, time.July, 19}:      PAPER,
	{2023, time.July, 26}:      COMMINGLES,
	{2023, time.August, 2}:     PAPER,
	{2023, time.August, 9}:     COMMINGLES,
	{2023, time.August, 16}:    PAPER,
	{2023, time.August, 23}:    COMMINGLES,
	{2023, time.August, 30}:    PAPER,
	{2023, time.September, 3}:  NO_PICKUP_HOLIDAY,
	{2023, time.September, 4}:  TRASH,
	{2023, time.September, 5}:  NO_PICKUP_HOLIDAY,
	{2023, time.September, 6}:  COMMINGLES,
	{2023, time.September, 7}:  TRASH,
	{2023, time.September, 12}: PAPER,
	{2023, time.September, 19}: COMMINGLES,
	{2023, time.September, 26}: PAPER,
	{2023, time.October, 4}:    COMMINGLES,
	{2023, time.October, 9}:    NO_PICKUP_HOLIDAY,
	{2023, time.October, 10}:   TRASH,
	{2023, time.October, 11}:   NO_PICKUP_HOLIDAY,
	{2023, time.October, 12}:   PAPER,
	{2023, time.October, 13}:   TRASH,
	{2023, time.October, 18}:   COMMINGLES,
	{2023, time.October, 25}:   PAPER,
	{2023, time.November, 1}:   COMMINGLES,
	{2023, time.November, 8}:   PAPER,
	{2023, time.November, 10}:  NO_PICKUP_HOLIDAY,
	{2023, time.November, 15}:  COMMINGLES,
	{2023, time.November, 22}:  PAPER,
	{2023, time.November, 23}:  NO_PICKUP_HOLIDAY,
	{2023, time.November, 24}:  TRASH,
	{2023, time.November, 29}:  COMMINGLES,
	{2023, time.December, 6}:   PAPER,
	{2023, time.December, 13}:  COMMINGLES,
	{2023, time.December, 20}:  PAPER,
	{2023, time.December, 25}:  NO_PICKUP_HOLIDAY,
	{2023, time.December, 26}:  TRASH,
	{2023, time.December, 27}:  NO_PICKUP_HOLIDAY,
	{2023, time.December, 28}:  COMMINGLES,
	{2023, time.December, 29}:  TRASH,
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
