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
	{2025, time.January, 1}:    NO_PICKUP_HOLIDAY,
	{2025, time.January, 2}:    PAPER,
	{2025, time.January, 3}:    TRASH,
	{2025, time.January, 8}:    COMMINGLES,
	{2025, time.January, 15}:   PAPER,
	{2025, time.January, 20}:   NO_PICKUP_HOLIDAY,
	{2025, time.January, 21}:   TRASH,
	{2025, time.January, 23}:   COMMINGLES,
	{2025, time.January, 24}:   TRASH,
	{2025, time.January, 29}:   PAPER,
	{2025, time.February, 5}:   COMMINGLES,
	{2025, time.February, 12}:  NO_PICKUP_HOLIDAY,
	{2025, time.February, 13}:  PAPER,
	{2025, time.February, 14}:  TRASH,
	{2025, time.February, 17}:  NO_PICKUP_HOLIDAY,
	{2025, time.February, 18}:  TRASH,
	{2025, time.February, 20}:  COMMINGLES,
	{2025, time.February, 21}:  TRASH,
	{2025, time.February, 26}:  PAPER,
	{2025, time.March, 5}:      COMMINGLES,
	{2025, time.March, 12}:     PAPER,
	{2025, time.March, 19}:     COMMINGLES,
	{2025, time.March, 26}:     PAPER,
	{2025, time.April, 2}:      COMMINGLES,
	{2025, time.April, 9}:      PAPER,
	{2025, time.April, 16}:     COMMINGLES,
	{2025, time.April, 23}:     PAPER,
	{2025, time.April, 30}:     COMMINGLES,
	{2025, time.May, 7}:        PAPER,
	{2025, time.May, 14}:       COMMINGLES,
	{2025, time.May, 21}:       PAPER,
	{2025, time.May, 26}:       NO_PICKUP_HOLIDAY,
	{2025, time.May, 27}:       TRASH,
	{2025, time.May, 29}:       COMMINGLES,
	{2025, time.May, 30}:       TRASH,
	{2025, time.June, 4}:       PAPER,
	{2025, time.June, 11}:      COMMINGLES,
	{2025, time.June, 18}:      PAPER,
	{2025, time.June, 19}:      NO_PICKUP_HOLIDAY,
	{2025, time.June, 20}:      TRASH,
	{2025, time.June, 25}:      COMMINGLES,
	{2025, time.July, 2}:       PAPER,
	{2025, time.July, 4}:       NO_PICKUP_HOLIDAY,
	{2025, time.July, 9}:       COMMINGLES,
	{2025, time.July, 15}:      PAPER,
	{2025, time.July, 23}:      COMMINGLES,
	{2025, time.July, 30}:      PAPER,
	{2025, time.August, 6}:     COMMINGLES,
	{2025, time.August, 13}:    PAPER,
	{2025, time.August, 20}:    COMMINGLES,
	{2025, time.August, 27}:    PAPER,
	{2025, time.September, 1}:  NO_PICKUP_HOLIDAY,
	{2025, time.September, 2}:  TRASH,
	{2025, time.September, 4}:  COMMINGLES,
	{2025, time.September, 5}:  TRASH,
	{2025, time.September, 10}: PAPER,
	{2025, time.September, 17}: COMMINGLES,
	{2025, time.September, 24}: PAPER,
	{2025, time.October, 1}:    COMMINGLES,
	{2025, time.October, 8}:    PAPER,
	{2025, time.October, 13}:   NO_PICKUP_HOLIDAY,
	{2025, time.October, 14}:   TRASH,
	{2025, time.October, 16}:   COMMINGLES,
	{2025, time.October, 17}:   TRASH,
	{2025, time.October, 22}:   PAPER,
	{2025, time.October, 29}:   COMMINGLES,
	{2025, time.November, 4}:   NO_PICKUP_HOLIDAY,
	{2025, time.November, 6}:   PAPER,
	{2025, time.November, 7}:   TRASH,
	{2025, time.November, 11}:  NO_PICKUP_HOLIDAY,
	{2025, time.November, 13}:  COMMINGLES,
	{2025, time.November, 14}:  TRASH,
	{2025, time.November, 19}:  PAPER,
	{2025, time.November, 26}:  COMMINGLES,
	{2025, time.November, 27}:  NO_PICKUP_HOLIDAY,
	{2025, time.November, 28}:  TRASH,
	{2025, time.December, 3}:   PAPER,
	{2025, time.December, 10}:  COMMINGLES,
	{2025, time.December, 17}:  PAPER,
	{2025, time.December, 24}:  COMMINGLES,
	{2025, time.December, 25}:  NO_PICKUP_HOLIDAY,
	{2025, time.December, 26}:  TRASH,
	{2025, time.December, 31}:  PAPER,
	{2024, time.January, 1}:    NO_PICKUP_HOLIDAY,
	{2024, time.January, 2}:    TRASH,
	{2024, time.January, 4}:    PAPER,
	{2024, time.January, 5}:    TRASH,
	{2024, time.January, 10}:   COMMINGLES,
	{2024, time.January, 15}:   NO_PICKUP_HOLIDAY,
	{2024, time.January, 16}:   TRASH,
	{2024, time.January, 18}:   PAPER,
	{2024, time.January, 19}:   TRASH,
	{2024, time.January, 24}:   COMMINGLES,
	{2024, time.January, 31}:   PAPER,
	{2024, time.February, 7}:   COMMINGLES,
	{2024, time.February, 12}:  NO_PICKUP_HOLIDAY,
	{2024, time.February, 13}:  TRASH,
	{2024, time.February, 15}:  PAPER,
	{2024, time.February, 16}:  TRASH,
	{2024, time.February, 19}:  NO_PICKUP_HOLIDAY,
	{2024, time.February, 20}:  TRASH,
	{2024, time.February, 22}:  COMMINGLES,
	{2024, time.February, 23}:  TRASH,
	{2024, time.February, 28}:  PAPER,
	{2024, time.March, 6}:      COMMINGLES,
	{2024, time.March, 13}:     PAPER,
	{2024, time.March, 20}:     COMMINGLES,
	{2024, time.March, 27}:     PAPER,
	{2024, time.April, 3}:      COMMINGLES,
	{2024, time.April, 10}:     PAPER,
	{2024, time.April, 17}:     COMMINGLES,
	{2024, time.April, 24}:     PAPER,
	{2024, time.May, 1}:        COMMINGLES,
	{2024, time.May, 8}:        PAPER,
	{2024, time.May, 15}:       COMMINGLES,
	{2024, time.May, 22}:       PAPER,
	{2024, time.May, 27}:       NO_PICKUP_HOLIDAY,
	{2024, time.May, 28}:       TRASH,
	{2024, time.May, 30}:       COMMINGLES,
	{2024, time.May, 31}:       TRASH,
	{2024, time.June, 5}:       PAPER,
	{2024, time.June, 12}:      COMMINGLES,
	{2024, time.June, 19}:      NO_PICKUP_HOLIDAY,
	{2024, time.June, 20}:      PAPER,
	{2024, time.June, 21}:      TRASH,
	{2024, time.June, 26}:      COMMINGLES,
	{2024, time.July, 3}:       PAPER,
	{2024, time.July, 4}:       NO_PICKUP_HOLIDAY,
	{2024, time.July, 5}:       TRASH,
	{2024, time.July, 10}:      COMMINGLES,
	{2024, time.July, 17}:      PAPER,
	{2024, time.July, 24}:      COMMINGLES,
	{2024, time.July, 31}:      PAPER,
	{2024, time.August, 7}:     COMMINGLES,
	{2024, time.August, 14}:    PAPER,
	{2024, time.August, 21}:    COMMINGLES,
	{2024, time.August, 28}:    PAPER,
	{2024, time.September, 2}:  NO_PICKUP_HOLIDAY,
	{2024, time.September, 3}:  TRASH,
	{2024, time.September, 5}:  COMMINGLES,
	{2024, time.September, 6}:  TRASH,
	{2024, time.September, 11}: PAPER,
	{2024, time.September, 18}: COMMINGLES,
	{2024, time.September, 25}: PAPER,
	{2024, time.October, 2}:    COMMINGLES,
	{2024, time.October, 9}:    PAPER,
	{2024, time.October, 14}:   NO_PICKUP_HOLIDAY,
	{2024, time.October, 15}:   TRASH,
	{2024, time.October, 17}:   COMMINGLES,
	{2024, time.October, 18}:   TRASH,
	{2024, time.October, 23}:   PAPER,
	{2024, time.October, 30}:   COMMINGLES,
	{2024, time.November, 5}:   NO_PICKUP_HOLIDAY,
	{2024, time.November, 7}:   PAPER,
	{2024, time.November, 8}:   TRASH,
	{2024, time.November, 11}:  NO_PICKUP_HOLIDAY,
	{2024, time.November, 12}:  TRASH,
	{2024, time.November, 14}:  COMMINGLES,
	{2024, time.November, 15}:  TRASH,
	{2024, time.November, 20}:  PAPER,
	{2024, time.November, 27}:  COMMINGLES,
	{2024, time.November, 28}:  NO_PICKUP_HOLIDAY,
	{2024, time.November, 29}:  TRASH,
	{2024, time.December, 4}:   PAPER,
	{2024, time.December, 11}:  COMMINGLES,
	{2024, time.December, 18}:  PAPER,
	{2024, time.December, 25}:  NO_PICKUP_HOLIDAY,
	{2024, time.December, 26}:  COMMINGLES,
	{2024, time.December, 27}:  TRASH,
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
