package model

// HOME_PAGE represents the root endpoint of the web service.
const HOME_PAGE = "/"

// Book_Count is the endpoint for accessing book count statistics.
const Book_Count = "/librarystats/v1/bookcount/"

// Book_Language is a template for the book count endpoint that includes a query parameter for filtering by language.
const Book_Language = "/librarystats/v1/bookcount/?language="

// Book_Language_Dutch is a specific instance of the Book_Language endpoint, pre-configured for Dutch ('de') language.
const Book_Language_Dutch = "/librarystats/v1/bookcount/?language=de"

// Book_Language_Dutch_French is a specific instance of the Book_Language endpoint, pre-configured for both Dutch ('de') and French ('fr') languages.
const Book_Language_Dutch_French = "/librarystats/v1/bookcount/?language=de,fr"

// TOTAL_BOOKS_IN_Library is a constant representing the total number of books available in the library.
// This value is used for calculating statistics such as the fraction of books available in a specific language.
const TOTAL_BOOKS_IN_Library int = 72810

// GUTENDEX is the base URL for the GutenDEX API, which is used to retrieve information about books.
const GUTENDEX = "https://gutendex.com/books?languages="

// ReadershipLink is the endpoint for accessing readership statistics.
const ReadershipLink = "/librarystats/v1/readership/"

// ReadershipExampleRequest is an example request to the readership endpoint for Norwegian ('no') language.
const ReadershipExampleRequest = "/librarystats/v1/readership/no"

// ReadershipExampleRequestWithLimit is an example request to the readership endpoint for Norwegian ('no') language with a limit query parameter.
// This example limits the response to the top 5 results.
const ReadershipExampleRequestWithLimit = "/librarystats/v1/readership/no/?limit=5"

// LANGUAGE_TO_COUNTRY is the base URL for an API that maps languages to countries.
const LANGUAGE_TO_COUNTRY = "http://129.241.150.113:3000/"

// POPULATION_API is the base URL for an API that provides population data for countries.
const POPULATION_API = "http://129.241.150.113:8080/"

// Status_Endpoint is the endpoint for accessing the status of various services and APIs the application depends on.
const Status_Endpoint = "/librarystats/v1/status/"
