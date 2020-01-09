# Jikan-go

[![GoDoc](https://godoc.org/github.com/darenliang/jikan-go?status.svg)](https://godoc.org/github.com/darenliang/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/darenliang/jikan-go)](https://goreportcard.com/report/github.com/darenliang/jikan-go)

A complete Go wrapper for the Jikan API.

Documentation can be found at [godoc.org](https://godoc.org/github.com/darenliang/jikan-go)

Please refer to the official [Jikan API documentation](https://jikan.docs.apiary.io) on available request parameters.

All json data is put into maps instead of structs, so consulting the official documentation of the REST API will show you the available data fields.

To visualize the response you can printout the map using `fmt.Println()` or you can use the `jikan.PrettyPrint()` function bundled in the package to print the map with indentation.

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
	anime, _ := jikan.Anime{ID: 1}.Get()

	// You can print values of type interface{}
	fmt.Println(anime["title"])
}
```
```
Output:

Cowboy Bebop
```
##### Nested Data with Type Assertions Example
```go
package main

import (
	"fmt"
	"github.com/darenliang/jikan-go"
)

func main() {
	search, _ := jikan.Search{Type: "anime", Q: "FMAB", OrderBy: "score"}.Get()

	// Get first anime from results
	firstAnime := search["results"].([]interface{})[0].(map[string]interface{})

	// Make type assertions
	fmt.Println("Title: " + firstAnime["title"].(string))
	fmt.Printf("Members: %v thousand", firstAnime["members"].(float64)/1000)
}
```
```
Output:

Title: Fullmetal Alchemist: Brotherhood
Members: 1533.286 thousand
```
## FAQ
### Why use maps over structs for reponses?
Due to the multiple different response structures of the Jikan API, it is not very convenient to call different functions depending on the return type of the struct.

Ideally, using functions like `Anime.Get()` should get you the Jikan's anime json data using structs. However, the problem is that adding many different structs creates a few nuisances that are hard to go unnoticed.

Here are some problems when using structs for responses:

1. You have to rely on the Go Docs for all struct fields since they are different from Jikan's normal key values.
2. You have to know the exact function name to call for a specific struct.
3. Whenever a new data field is added to the Jikan API, the corresponding struct that is supposed to hold the data field will ignore it unless it is updated.

### Why do I have to perform type assertions when I want to use map data?
Golang doesn't allow dynamic return types for functions other than using interfaces.

Returning interfaces doesn't allow you to use the interface value's underlying concrete value.

Since all structs returned are in the type `map[string]interface{}`, you can only readily use first level data.

For accessing lower levels of data, you must perform type assertions in advance.

To type assert a slice/array, append `.([]interface{})` before accessing the index.

To type assert a map, append `.(map[string]interface{})` before accessing the index.

To use a golang basic type, append `.(type)` where type is the data type you want to assert. Such examples include `.(string)`, `.(float64)` and `.(bool)`.

### Why do I need to create new structs every time I use the API?
Due to the nature of Golang, it is [not possible to use optional parameters in functions](https://golang.org/doc/faq#overloading).

It starts becoming a real problem when Jikan API's "search" or "user" request is used. There is just no simple way of calling only the request parameters you need directly. Although there are such things as "variadic arguments", it does not provide the ability to only add request parameters you need.

One way around this problem is by passing structs as a function parameter. We can take advantage of Golang's zero values, allowing us to not have to specify every single request parameter.

If there is a better way of addressing this problem, please let me know.
