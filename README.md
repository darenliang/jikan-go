# Jikan Go

These Golang API bindings are for Jikan's REST API v3.

Almost all of the Jikan's API can be accessed through these bindings except the API's metadata reference.

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
	anime, _ := jikan.GetAnimeInfo(1)
	fmt.Println(anime.Title)
}
```
```
Output:
Cowboy Bebop
```