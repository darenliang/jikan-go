# Jikan Go

These Golang bindings cover the complete Jikan API

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