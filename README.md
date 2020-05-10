# Jikan-go

[![GoDoc](https://godoc.org/github.com/darenliang/jikan-go?status.svg)](https://godoc.org/github.com/darenliang/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/darenliang/jikan-go)](https://goreportcard.com/report/github.com/darenliang/jikan-go)

**The library has changed significantly. The current version is not backwards compatible with older versions.**

A complete Go wrapper for the Jikan API.

Documentation can be found at [godoc.org](https://godoc.org/github.com/darenliang/jikan-go)

Refer to the official [Jikan API documentation](https://jikan.docs.apiary.io) for clarification on request parameters.

### Installation

To install/update: `go get -u github.com/darenliang/jikan-go`

To import: `import "github.com/darenliang/jikan-go"` and use as `jikan`

### Usage

##### Basic Example
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	// Get anime
	anime, _ := jikan.GetAnime(1)
	fmt.Println(anime.Title)
}
```
```
Output:

Cowboy Bebop
```
##### Search Query Example

Consult the [Jikan search reference](https://jikan.docs.apiary.io/#reference/0/search) for a list of search parameters.

```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
	"net/url"
)

func main() {
	// Setup query
	query := url.Values{}
	query.Set("q", "Cowboy Bebop")
	query.Set("type", "tv")

	// Search anime
	search, _ := jikan.GetSearchAnime(query)
	fmt.Println(search.Results[0].Score)
}
```
```
Output:

8.79
```
##### Troubleshooting

If it is necessary to modify the http client (eg. modify timeout), you can access the client via `jikan.Client`.

Example: setting the client timeout from 60 seconds (default) to 10 seconds.
```go
jikan.Client.Timeout = time.Second * 10
```
