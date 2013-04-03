package goyum

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"unicode"
)

type SearchResults struct {
	Attribution     map[string]interface{} `json:"attribution"`
	TotalMatchCount int                    `json:"totalMatchCount"`
	FacetCounts     map[string]interface{} `json:"facetCounts"`
	Matches         []interface{}          `json:"matches"`
	y               *Yummly
}

type SearchParams struct {
	q                 string
	RequirePictures   bool
	AllowedIngredient []string
}

func downcase(target string) string {
	a := []rune(target)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func (sp *SearchParams) values() url.Values {
	v := make(url.Values)
	s := reflect.ValueOf(sp).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		name := downcase(typeOfT.Field(i).Name)

		switch f.Kind() {
		case reflect.String:
			v.Add(name, f.String())
		case reflect.Bool:
			v.Add(name, strconv.FormatBool(f.Bool()))
		case reflect.Int:
			v.Add(name, strconv.FormatInt(f.Int(), 10))
		case reflect.Slice:
			s := f.Slice(0, f.Len())
			var buf bytes.Buffer
			buf.WriteString(name)
			buf.WriteString("[]")

			for j := 0; j < f.Len(); j++ {
				v.Add(buf.String(), s.Index(j).String())
			}
		default:
			fmt.Println("huh?")
		}
	}

	return v
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
