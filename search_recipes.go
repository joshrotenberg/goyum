package goyum

import (
)

type SearchResults struct {
	Attribution     map[string]interface{} `json:"attribution"`
	TotalMatchCount int                    `json:"totalMatchCount"`
	FacetCounts     map[string]interface{} `json:"facetCounts"`
	Matches         []interface{}          `json:"matches"`
	y               *Yummly
}

func (y *Yummly) SearchRecipes(q string, params *SearchParams) (*SearchResults, error) {
	var result SearchResults
	params.q = q
	//fmt.Println(params)
	err := y.callYummlyApi("GET", "recipes", params, &result)
	if err != nil {
		return nil, err
	}
	result.y = y
	return &result, nil

}
