package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRegression(t *testing.T) {
	r := ` { "queryResult": { "parameters": { "date-time": "" } } }`
	req := httptest.NewRequest("POST", "/trash", strings.NewReader(r))
	w := httptest.NewRecorder()
	Trash(w, req)
	if w.Code != 200 {
		t.Errorf("Got: %d, want 200\n%s", w.Code, w.Body)
	}
}

func day(m time.Month, d int) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal("Unable to load 'America/New_York' as a location.")
	}
	return time.Date(2022, m, d, 0, 0, 0, 0, loc)
}

func NextPickupS(p Pickup, t time.Time) string {
	buf := new(bytes.Buffer)
	NextPickupResponse(buf, p, t)
	return buf.String()
}

func PickForDayS(t time.Time) string {
	buf := new(bytes.Buffer)
	PickForDayResponse(buf, t)
	return buf.String()
}

func TestFooFoo(t *testing.T) {
	tests := []struct {
		dateTime string
		pickup   string
		want     string
	}{
		{
			dateTime: "2022-05-27T03:55:20.550Z",
			pickup:   "",
			want:     PickForDayS(day(time.May, 27)),
		},
		{
			dateTime: "2022-05-27T03:55:20.550Z",
			pickup:   "cardboard",
			want:     NextPickupS(PAPER, day(time.June, 8)),
		},
	}

	for _, tc := range tests {
		r := fmt.Sprintf(`{ "queryResult": { "parameters": { "date-time": "%s", "Pickup": "%s" } } }`, tc.dateTime, tc.pickup)
		req := httptest.NewRequest("POST", "/trash", strings.NewReader(r))
		w := httptest.NewRecorder()

		Trash(w, req)

		result := w.Result()
		responseText, _ := io.ReadAll(result.Body)
		s := fmt.Sprintf("%s", responseText)

		if w.Code != 200 || s != tc.want {
			t.Errorf("Got Code: %d, want 200\nGot response: %s\bWant response: %s", w.Code, w.Body, tc.want)
		}
	}
}
