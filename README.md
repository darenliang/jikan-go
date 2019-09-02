# Jikan Go

[![GoDoc](https://godoc.org/github.com/darenliang/jikan-go?status.svg)](https://godoc.org/github.com/darenliang/jikan-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/darenliang/jikan-go)](https://goreportcard.com/report/github.com/darenliang/jikan-go)

These Golang bindings cover the complete Jikan API

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
	anime, _ := jikan.GetAnime(Anime{ID: 1})
	fmt.Println(anime["title"])
}
```
```
Output:
Cowboy Bebop
```
### Why do I need to create new structs everytime I use the API?
Due to the nature of Golang, it is [not possible to use optional parameters in functions](https://golang.org/doc/faq#overloading).

It starts becoming a real problem when Jikan API's "search" or "user" request is used. There is just no simple way of calling only the request parameters you need directly. Although there are such things as "variadic arguments", it is strong typed and will not take in different data types.

One way around this problem is by using structs. We can take advantage of Golang's zero values, allowing us to not have to specify every single request parameter.

If there is a better way of addressing this problem, please let me know.