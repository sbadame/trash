package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

const HTML = `
<!DOCTYPE html>
<html>
  <head>
    <title>Yonkers Trash Schedule</title>
  </head>
  <body>
  {{range .TrashDates}}
    <div>
      {{.Date}} - {{.Pickup}}
    </div>
  {{end}}
  </body>
</html>
`

type PickupDate struct {
	Date   string
	Pickup Pickup
}

func nextWeekOfDays(startDate time.Time) []PickupDate {
	var r []PickupDate
	for i := 0; i < 7; i++ {
		d := startDate.AddDate(0, 0, i)
		r = append(r, PickupDate{d.Format("Mon 01/02"), ForDate(d)})
	}
	return r
}

func writeHtmlForTime(out io.Writer, t time.Time) error {
	tmpl, err := template.New("page").Parse(HTML)
	if err != nil {
		return err
	}

	data := struct {
		TrashDates []PickupDate
	}{
		nextWeekOfDays(t),
	}
	return tmpl.Execute(out, data)
}

func TrashHTML(w http.ResponseWriter, req *http.Request) {
	err := writeHtmlForTime(w, time.Now())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal error: %v", err)
	}
}

func main() {
	// https://cloud.google.com/run/docs/reference/container-contract#port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	fmt.Printf("Serving on :%s\n", port)
	http.HandleFunc("/trash", TrashVoice)
	http.HandleFunc("/", TrashHTML)
	http.ListenAndServe(":"+port, nil)
}