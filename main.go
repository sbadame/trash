package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sbadame/trash/pkg/trash"
	"html/template"
	"net/http"
	"os"
	"time"
)

var accessControlAllowOrigin string

func init() {
	flag.StringVar(&accessControlAllowOrigin, "access-control-allow-origin", "", "When set will add the Access-Control-Allow-Origin header to requests with the value provided.")
}

//go:embed static/schedule_2025.pdf
var schedule_2025_pdf []byte

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
      <caption><a href="/schedule_2025.pdf">Schedule PDF 2025</a></caption>
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
	Date        string
	Pickup      trash.Pickup
	PickupClass template.HTMLAttr
}

func nextWeekOfDays(startDate time.Time) []PickupDate {
	var r []PickupDate
	for i := 0; i < 7; i++ {
		d := startDate.AddDate(0, 0, i)
		p := trash.ForDate(d)

		class := ""
		if p == trash.NO_PICKUP || p == trash.NO_PICKUP_HOLIDAY {
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

func TrashHTML(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.New("page").Parse(HTML)
	if err != nil {
		http.Error(w, "Internal Error parsing the HTML template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		TrashDates []PickupDate
	}{
		nextWeekOfDays(time.Now()),
	}

	if accessControlAllowOrigin != "" {
		w.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Error executing the HTML template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func TrashJSON(w http.ResponseWriter, req *http.Request) {
	startDate := time.Now()

	pickups := make(map[int64]string) // milliseconds from unix epoch --> pickup string
	for i := 0; i < 7; i++ {
		d := startDate.AddDate(0, 0, i)
		pickups[d.Unix()*1000] = trash.ForDate(d).String()
	}

	payload := struct {
		Dates map[int64]string `json:"dates"`
	}{
		pickups,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Error converting data to json: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if accessControlAllowOrigin != "" {
		w.Header().Set("Access-Control-Allow-Origin", accessControlAllowOrigin)
	}
	w.Write(j)
}

func DownloadSchedule2025(w http.ResponseWriter, req *http.Request) {
	if _, err := w.Write(schedule_2025_pdf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Error: "+err.Error())
	}
}

func main() {
	flag.Parse()

	// https://cloud.google.com/run/docs/reference/container-contract#port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	fmt.Printf("Serving on :%s\n", port)
	http.HandleFunc("/", TrashHTML)
	http.HandleFunc("/index.json", TrashJSON)
	http.HandleFunc("/schedule_2025.pdf", DownloadSchedule2025)
	http.ListenAndServe(":"+port, nil)
}
