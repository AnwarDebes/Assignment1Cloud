package model

// BookCount represents the count of books in a particular language, along with the number of authors and the fraction of books in that language compared to the total number of books in the library.
type BookCount struct {
	Language string  `json:"language"` // The language of the books, typically represented by its ISO code.
	Count    int     `json:"books"`    // The total number of books available in the specified language.
	Authors  int     `json:"authors"`  // The total number of authors who have written books in the specified language.
	Fraction float64 `json:"fraction"` // The fraction of the total books in the library that are written in the specified language.
}

// Readership represents readership statistics for a given language, including the country of readership, ISO country code, and the number of books and authors.
type Readership struct {
	Country    string  `json:"country"`    // The name of the country associated with the readership statistics.
	IsoCode    string  `json:"iso_code"`   // The ISO code of the country, typically represented as a two-letter country code.
	Books      int     `json:"books"`      // The total number of books read in the specified country/language.
	Authors    int     `json:"authors"`    // The total number of authors read in the specified country/language.
	Readership float64 `json:"readership"` // An estimated number of readers or readership metric in the specified country/language.
}

// Status represents the status of various external APIs or services that the application depends on, including their HTTP status codes, the application version, and uptime.
type Status struct {
	Qutendexapi  int    `json:"gutendexapi"`  // The HTTP status code of the Gutendex API, indicating its availability.
	Languageapi  int    `json:"languageapi"`  // The HTTP status code of the language to country mapping API, indicating its availability.
	Countriesapi int    `json:"countriesapi"` // The HTTP status code of the countries API, indicating its availability.
	Version      string `json:"version"`      // The version of the application or service.
	Uptime       string `json:"uptime"`       // The uptime of the application or service, typically represented as a duration string.
}
