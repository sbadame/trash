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
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
    <style>
       body {
           font-family: helvetica, sans-serif;
           font-size: 4vw;
           background-color: #d8fbee;
           display: flex;
           flex-direction: row;
           align-content: center;
           justify-content: center;
           align-items: center;
       }
       .nopickup {
          color: #a8c7bb;
       }
       table {
           text-align: center;
           border-collapse: collapse;
       }
       td {
           padding: 5px;
           border: 2px;
       }
       table td + td {
           border-left: 2px solid black;
       }
       caption {
         caption-side: bottom;
         padding: 10px;
       }
    </style>
  </head>
  <body>
    <table>
      <caption><a href="https://www.yonkersny.gov/home/showpublisheddocument/30455/637743931515170000">Schedule PDF</a></caption>
      <thead><tr><th>Date</th><th>Pickup</th></tr></thead>
      <tbody>
        {{range .TrashDates}}
          <tr><td>{{.Date}}</td><td{{.PickupClass}}>{{.Pickup}}</td></tr>
        {{end}}
      </tbody>
    </table>
  </body>
</html>
`

type PickupDate struct {
	Date   string
	Pickup Pickup
	PickupClass template.HTMLAttr
}

func nextWeekOfDays(startDate time.Time) []PickupDate {
	var r []PickupDate
	for i := 0; i < 7; i++ {
		d := startDate.AddDate(0, 0, i)
		p := ForDate(d)

		class := ""
		if p == NO_PICKUP || p == NO_PICKUP_HOLIDAY {
			class = " class=\"nopickup\""
		}

		r = append(r, PickupDate{
			d.Format("Mon 01/02"),
			p,
			template.HTMLAttr(class),
		})
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
