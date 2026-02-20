package Structs

// Country Decode the country rest API get request in InfoHandler
type Country struct {
	Flags      map[string]string `json:"flags"`
	Name       map[string]any    `json:"name"`
	Capital    []string          `json:"capital"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Area       float64           `json:"area"`
	Population int               `json:"population"`
	Continents []string          `json:"continents"`
}

// Currency Decode the rest currency API get request in ExchangeHandler
type Currency struct {
	Name       map[string]any `json:"name"`
	Borders    []string       `json:"borders"`
	Currencies map[string]any `json:"currencies"`
}

// State formats the status requests to get response
type State struct {
	RestCountriesApi string
	CurrenciesApi    string
	Version          string
	Uptime           int
}

// Exchange Decodes the currency api get request
type Exchange struct {
	Rates map[string]float64 `json:"rates"`
}

// Valuta Decodes the rest api get request
type Valuta struct {
	Currency map[string]any `json:"currencies"`
}
