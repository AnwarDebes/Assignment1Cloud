package model

type BookCount struct {
	Language string  `json:"language"`
	Count    int     `json:"books"`
	Authors  int     `json:"authors"`
	Fraction float64 `json:"fraction"`
}
type Readership struct {
	Country    string  `json:"country"`
	IsoCode    string  `json:"iso_code"` // Changed to string to match typical ISO code formats
	Books      int     `json:"books"`
	Authors    int     `json:"authors"`
	Readership float64 `json:"readership"`
}
type Status struct {
	Qutendexapi  int    `json:"gutendexapi"`
	Languageapi  int    `json:"languageapi"`
	Countriesapi int    `json:"countriesapi"`
	Version      string `json:"version"`
	Uptime       string `json:"uptime"`
}
