package goyum

import (
	_ "fmt"
	_ "net/url"
	_ "unicode"
)

type SearchResults struct [

	y *Yummly
}

type SearchParams struct {
	Foo[] []string
}

func (y *Yummly) SearchRecipes(q string, params *SearchParams) (*SearchResult, error) {
	var result SearchResult
	err := y.callYummlyApi("GET", fmt.Sprintf("/

}
