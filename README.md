# Jikan-go

[![GoDoc](https://godoc.org/github.com/darenliang/jikan-go?status.svg)](https://godoc.org/github.com/darenliang/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/darenliang/jikan-go)](https://goreportcard.com/report/github.com/darenliang/jikan-go)

**The library has changed significantly. The current version is not backwards compatible with older versions.**

A complete Go wrapper for the Jikan API.

Documentation can be found at [pkg.go.dev](https://pkg.go.dev/github.com/darenliang/jikan-go)

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
	anime, _ := jikan.GetAnime(1)
	fmt.Println(anime.Title)
}
```
```
Output:

Cowboy Bebop
```
##### Search Query Example
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	search, _ := jikan.GetSearch(jikan.SearchQuery{
		Type: "anime",
		Q:    "Cowboy Bebop",
	})
	fmt.Println(search.Results[0].Score)
}
```
```
Output:

8.79
```
##### Troubleshooting

If it is necessary to modify the http client (eg. modify timeout), you can access the client via `jikan.Client`.