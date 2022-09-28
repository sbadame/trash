package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func writeResponse(w io.Writer, message string) error {
	// https://developers.google.com/assistant/df-asdk/reference/dialogflow-webhook-json#simple-response-example-df
	_, err := fmt.Fprintf(w, `{
  "payload": {
    "google": {
      "expectUserResponse": false,
      "richResponse": {
        "items": [
          {
            "simpleResponse": {
              "textToSpeech": "%s",
              "displayText": "%s"
            }
          }
        ]
      }
    }
  }
}`, message, message)
	return err
}

func NextPickupResponse(w io.Writer, p Pickup, when time.Time) {
	writeResponse(w, fmt.Sprintf("The next %s pickup is %s.", p, when.Format("Monday January 2")))
}

func PickForDayResponse(w io.Writer, when time.Time) {
	writeResponse(w, fmt.Sprintf("The pickup for %s is %s", when.Weekday(), ForDate(when)))
}

func TrashVoice(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	log.Printf("%q\n", dump)

	// https://cloud.google.com/dialogflow/es/docs/fulfillment-webhook#webhook_request
	requestJson, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var r map[string]interface{}
	if err := json.Unmarshal(requestJson, &r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queryResult, ok := r["queryResult"].(map[string]interface{})
	if !ok {
		http.Error(w, "queryResult is not a map[string]interface{}", http.StatusInternalServerError)
		return
	}

	p, ok := queryResult["parameters"].(map[string]interface{})
	if !ok {
		http.Error(w, "parameters is not a map[string]interface{}", http.StatusInternalServerError)
		return
	}

	parameters := make(map[string]string)
	for k, v := range p {
		parameters[k] = v.(string)
	}

	t := time.Now()
	dt := parameters["date-time"]
	if dt != "" {
		t, err = time.Parse(time.RFC3339, dt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	pickupStr := parameters["Pickup"]
	if pickupStr != "" {
		// Given that pickup type was selected, the query is about getting the next date.
		pickup := FromString(pickupStr)
		np, err := NextPickup(t, pickup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		NextPickupResponse(w, pickup, np)
		return
	}

	// No pickup was specified but a date was, giving the pickup type for that date.
	PickForDayResponse(w, t)
}
