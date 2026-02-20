package Handlers

import (
	"encoding/json"
	"errors"
	"log"
	"main/Structs"
	"net/http"
	"strconv"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		infoHandlerGet(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func infoHandlerGet(w http.ResponseWriter, r *http.Request) {
	country := Structs.Country{}
	mp := make(map[string]any)
	// tester med no
	resp, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/" + r.PathValue("p1") + "?fields=name,capital,languages,borders,area,population,continents,flags")

	if err != nil {
		// To user
		http.Error(w, "Ops something on our end went wrong: "+string(http.StatusInternalServerError), http.StatusInternalServerError)
		// To server
		log.Fatal(err.Error(), ", Could not send Get request")
		return
	}
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "No response "+strconv.Itoa(resp.StatusCode), resp.StatusCode)
		log.Println(resp.StatusCode, " Bad request from user")
		return
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&country)

	if err != nil {
		http.Error(w, "Decode Error", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	err, mp = formatingInfoResponse(country)
	if err != nil {
		http.Error(w, "Ops something on our end went wrong: "+string(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Fatal(err.Error(), ", Could not format map for Get response")
	}

	// gjør sånn at jeg skriver ut til get hver gang jeg kaller encoder.encode(noe)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	err = encoder.Encode(&mp)

	if err != nil {
		http.Error(w, "Encode Error", http.StatusBadRequest)
	}

}

func formatingInfoResponse(country Structs.Country) (error, map[string]any) {
	nyMap := map[string]any{}

	nyMap["Name"] = country.Name["common"]
	nyMap["Capital"] = country.Capital
	nyMap["Languages"] = country.Languages
	nyMap["Flag"] = country.Flags["png"]
	nyMap["Population"] = country.Population
	nyMap["Continents"] = country.Continents
	nyMap["area"] = country.Area
	nyMap["Borders"] = country.Borders

	for _, value := range nyMap {
		if value == nil {
			return errors.New("trouble formating map"), nil
		}
	}

	return nil, nyMap
}
