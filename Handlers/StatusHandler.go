package Handlers

import (
	"encoding/json"
	"log"
	"main/Structs"
	"net/http"
	"time"
)

const VERSION string = "v1"

var start = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Checking if the request is a Get. Error njhandling if not
	if r.Method == http.MethodGet {
		statusHandlerGet(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func statusHandlerGet(w http.ResponseWriter, r *http.Request) {
	info := Structs.State{}

	// Rest Countries endpoint check
	resp, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/no")
	if err != nil {
		http.Error(w, "Ops something went wrong on our side"+string(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err.Error())
	}

	// Currency endpoint check
	resp1, err1 := http.Get("http://129.241.150.113:9090/currency/NOK")
	if err1 != nil {
		http.Error(w, "Ops something went wrong on our side"+string(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err1.Error())
	}

	// Making the Get response for my application
	info.CurrenciesApi = resp1.Status
	info.RestCountriesApi = resp.Status
	info.Version = VERSION
	info.Uptime = int(time.Since(start).Seconds())

	encoder := json.NewEncoder(w)
	err = encoder.Encode(info)
	if err != nil {
		http.Error(w, "Ops, something went wrong with encoding", http.StatusInternalServerError)
		log.Println(err)
	}

}
