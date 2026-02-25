package Handlers

import (
	"encoding/json"
	"log"
	"main/Structs"
	"maps"
	"net/http"
	"slices"
	"strings"
)

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		exchangeHandlerGet(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func exchangeHandlerGet(w http.ResponseWriter, r *http.Request) {
	var exchange Structs.Currency
	mp := make(map[string]any)

	// Rest countries
	resp, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/" + r.PathValue("p1") + "?fields=name,borders,currencies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&exchange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err, mp = formatingExchangeResponse(exchange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(mp)
	if err != nil {
		http.Error(w, "Ops, something went wrong", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

}

func formatingExchangeResponse(exchange Structs.Currency) (error, map[string]any) {
	getResponseMap := make(map[string]any)   // The get response that is going to be sent to User
	var valutaRates Structs.Exchange         // The exchangerates of the asked for country
	valutaNames := make(map[string][]string) // Map that collects all valuta names of all countries
	var names []string                       // Temporary variable to construct the valutaNames map
	var valuta Structs.Valuta                // Struct to decode get request into ´
	tempOutput := make(map[string]float64)   // Formating the output for the get request

	// fetching neighboring counteries valuta names
	for i := 0; i < len(exchange.Borders); i++ {
		//fmt.Println(exchange.Borders[i])
		resp, err := http.Get("http://129.241.150.113:8080/v3.1/alpha/" + exchange.Borders[i] + "?fields=currencies")
		if err != nil {
			return err, nil
		}
		// Decoder get request
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&valuta)
		if err != nil {
			return err, nil
		}

		// Appends Currency names to slice
		for k, _ := range valuta.Currency {
			names = append(names, k)
		}
		// constructs map
		valutaNames[exchange.Borders[i]] = names
	}

	// If there is more than one currency the loop covers all the currencies
	for key, _ := range exchange.Currencies {
		// sending get request to currencies to fetch exchange rate for each valuta that original country
		resp, err := http.Get("http://129.241.150.113:9090/currency/" + key)
		if err != nil {
			return err, nil
		}

		decoder := json.NewDecoder(resp.Body)

		err = decoder.Decode(&valutaRates)
		if err != nil {
			return err, nil
		}

		// Loop through the exchange rate map
		for valuta1, value := range valutaRates.Rates {
			// searching for neighboring countries exchange rates
			for _, valuta2 := range valutaNames {
				// Loops thorugh a countries currencies (if more than one)
				for _, name := range valuta2 {
					// check if the currency is a neighbor currency
					if strings.Compare(valuta1, name) == 0 {
						// TODO: denne skriver over allerede eksisterende variabler hvis det er mer enn 1 base currency. Så må fikse det
						tempOutput[valuta1] = value
					}
				}
			}
		}

	}

	// extracts all currencies of a country
	tempBaseCurrency := slices.Collect(maps.Keys(exchange.Currencies))

	// Constructs the get response for the user
	getResponseMap["Country"] = exchange.Name["common"]
	getResponseMap["Base-Currency"] = tempBaseCurrency
	getResponseMap["Exchange-Rates"] = tempOutput

	return nil, getResponseMap

}
