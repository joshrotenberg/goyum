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
	// check result.TotalMatchCount for the number of results
	// dig around in result.Matches for the recipe you want
}
```

Full API documentation can be found at godoc.org: http://godoc.org/github.com/joshrotenberg/goyum

## About

Full disclosure: I wrote this mainly as an exercise to get familiar with Go. I'm not an employee of/contractor for Yummly, and Yummly does not support this library. If you happen to use this library and have issues or requests, please file them [here](https://github.com/joshrotenberg/goyum/issues).
