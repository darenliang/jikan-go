# Jikan Go

[![GoDoc](https://godoc.org/github.com/darenliang/jikan-go?status.svg)](https://godoc.org/github.com/darenliang/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/darenliang/jikan-go)](https://goreportcard.com/report/github.com/darenliang/jikan-go)

Go bindings for Jikan.

Since the documentation is still not complete yet, please refer to the official [Jikan API documentation](https://jikan.docs.apiary.io) for assistance on certain request parameters.

To install: `go get github.com/darenliang/jikan-go`

To import: `import "github.com/darenliang/jikan-go"` and use as `jikan`

Sample usage:
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	anime, _ := jikan.GetAnime(jikan.Anime{ID: 1})
	fmt.Println(anime["title"])
}
```
```
Output:
Cowboy Bebop
```
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	search, _ := jikan.GetSearch(jikan.Search{Type: "anime", Q: "FMAB", OrderBy: "score"})
	firstAnime := search["results"].([]interface{})[0].(map[string]interface{})
	fmt.Println(firstAnime["title"], firstAnime["mal_id"])
}
```
```
Output:
Fullmetal Alchemist: Brotherhood 5114
```
### Why do I have to perform type assertions when I want to use the nested data?
Golang doesn't allow dynamic return types for functions other than using interfaces.

Returning interfaces doesn't allow you to use the interface value's underlying concrete value.

Since all structs returned are in the type `map[string]interface{}`, you can only readily use first level data.

For accessing lower levels of data, you must perform type assertions in advance.

To type assert a slice/array, append `.([]interface{})` before accessing the index.

To type assert a map, append `.(map[string]interface{})` before accessing the index.

### Why do I need to create new structs every time I use the API?
Due to the nature of Golang, it is [not possible to use optional parameters in functions](https://golang.org/doc/faq#overloading).

It starts becoming a real problem when Jikan API's "search" or "user" request is used. There is just no simple way of calling only the request parameters you need directly. Although there are such things as "variadic arguments", it does not provide the ability to only add request parameters you need.

One way around this problem is by using structs. We can take advantage of Golang's zero values, allowing us to not have to specify every single request parameter.

If there is a better way of addressing this problem, please let me know.
