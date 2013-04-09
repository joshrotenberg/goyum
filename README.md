# goyum

A Go library for the Yummly API.

This is currently a work in progress.

## Overview

goyum is a client library for the [Yummly](http://yummly.com) REST API written in [Go](http://golang.org).

## Docs

```go
var y *Yummly
y, err = SetCredentials(myAppId, myAppKey)
sp := NewSearchParams("onion soup")
sp.RequirePictures(true).AddAllowedIngredients("garlic", "onions")
sp.AddExcludedIngredients("potato").AddFlavorMax("meaty", 0.5)
result, err := y.SearchRecipes(sp)
if err != nil {
	// deal with err
} else {
	// dig around in the result for the recipe you want
}
```

Full API documentation can be found at godoc.org: http://godoc.org/github.com/joshrotenberg/goyum


