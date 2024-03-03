
# Assignment1Cloud

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Contributing](#contributing)
- [Credits](#credits)
- [License](#license)

## Introduction

This project, Assignment1Cloud, is a library statistics tool designed to provide users with comprehensive data on book counts, readership statistics, and service statuses across different languages. It aims to support libraries, readers, and researchers in accessing extensive bibliographic data, enhancing resource management, and facilitating literary studies.
## Getting Started
To get a local copy up and running, follow these simple steps.
### Prerequisites

Before you begin, ensure you have the following installed:

Go (version 1.21)

Access to an internet connection for external API calls
### Installation
1. Clone the repository:

   git clone https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2024-workspace/anwar/assignment1cloud.git

2. Navigate to the project directory:

   cd Assignment1Cloud

3. Install Go dependencies:

   go get .

4. Run the project:

    go run main.go
## Usage

This section provides a detailed guide on how to use the application, including the available endpoints and their purposes.

### Endpoints

The application provides the following endpoints:

1. `GET /book-count`: Retrieves the number of books in the library by language.
2. `GET /readership`: Retrieves readership statistics including country, ISO code, and total numbers of books and authors.
3. `GET /status`: Checks the status of the application's dependencies.


#### `GET /book-count`

This endpoint retrieves the number of books in the library by language.

##### Request

```http
Method: GET
Path: bookcount/?language={:two_letter_language_code+}/
```
Example requests:

http://localhost:8080/librarystats/v1/bookcount/?language=no

http://localhost:8080/librarystats/v1/bookcount/?language=no,fi

##### Response

```
[
  {
     "language": "no",
     "books": 21,
     "authors": 14,
     "fraction": 0.0005
  },
  {
     "language": "fi",
     "books": 2798,
     "authors": 228,
     "fraction": 0.0671
  }
]

```

#### `GET /readership`

This endpoint retrieves readership statistics including country, ISO code, and total numbers of books and authors.

##### Request

```http
Method: GET
Path: readership/{:two_letter_language_code}{?limit={:number}}

```

Example requests:

http://localhost:8080/librarystats/v1/readership/no

http://localhost:8080/librarystats/v1/readership/no?limit=3

##### Response

```
[ 
  {
     "country": "Norway",
     "isocode": "NO",
     "books": 21,
     "authors": 14,
     "readership": 5379475
  },
  {
     "country": "Svalbard and Jan Mayen",
     "isocode": "SJ",
     "books": 21,
     "authors": 14,
     "readership": 2562
  },
  {
     "country": "Iceland",
     "isocode": "IS",
     "books": 21,
     "authors": 14,
     "readership": 366425
  }
]

```


#### `GET /status`

This endpoint checks the status of the application's dependencies.

##### Request

```http
Method: GET
Path: status/

```

Example request:

http://localhost:8080/librarystats/v1/status/


##### Response

```
{
   "gutendexapi": "<http status code for gutendex API>",
   "languageapi: "<http status code for language2countries API>", 
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}

```


## Features

Assignment1Cloud includes the following features:

1. Book count retrieval by language.
2. Readership statistics including country, ISO code, and total numbers of books and authors.
3. API status checks for dependencies.
4. Multi-language support for queries.


## Credits
This project was developed by Anwar Debes as part of an academic assignment for the Cloud Computing course at NTNU Gjøvik.

