package main

import (
	"testing"
	"time"
)

func TestForDate(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal("Unable to load 'America/New_York' as a location.")
	}
	tests := []struct {
		time time.Time
		want Pickup
	}{
		{time: time.Date(2022, time.May, 25, 0, 0, 0, 0, loc), want: PAPER},
		{time: time.Date(2022, time.May, 30, 0, 0, 0, 0, loc), want: NO_PICKUP_HOLIDAY},
		{time: time.Date(2022, time.June, 3, 0, 0, 0, 0, loc), want: TRASH},
		{time: time.Date(2022, time.June, 13, 0, 0, 0, 0, loc), want: TRASH},
		{time: time.Date(2022, time.June, 15, 0, 0, 0, 0, loc), want: COMMINGLES},
		{time: time.Date(2022, time.June, 15, 1, 1, 1, 1, loc), want: COMMINGLES},
	}

	for _, tc := range tests {
		got := ForDate(tc.time)
		if got != tc.want {
			t.Errorf("date: %v got: %v want: %v", tc.time, got, tc.want)
		}
	}
}

func TestNextPickup(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal("Unable to load 'America/New_York' as a location.")
	}
	tests := []struct {
		from    time.Time
		pickup  Pickup
		want    time.Time
		wantErr error
	}{
		{
			from:    time.Date(2022, time.May, 27, 0, 0, 0, 0, loc),
			pickup:  TRASH,
			want:    time.Date(2022, time.May, 31, 0, 0, 0, 0, loc),
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		got, err := NextPickup(tc.from, tc.pickup)
		if err != tc.wantErr || got != tc.want {
			t.Errorf("from: %v pickup: %v, got: %v err: %v, want: %v, wantErr: %v", tc.from, tc.pickup, got, err, tc.want, tc.wantErr)
		}
	}
}
