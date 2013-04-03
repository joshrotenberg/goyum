package goyum

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type Flavors map[string]int

type SearchParams struct {
	q                     string
	RequirePictures       bool
	AllowedIngredient     []string
	AllowedDiet           []string
	AllowedAllergy        []string
	AllowedCuisine        []string
	AllowedCourse         []string
	AllowedHoliday        []string
	ExcludedIngredient    []string
	ExcludedCuisine       []string
	ExcludedCourse        []string
	ExcludedHoliday       []string
	MaxTotalTimeInSeconds int
	MaxResult             int
	Start                 int
	FacetField            []string
	Flavors               map[string]int
}

func NewSearchParams() *SearchParams {
	p := new(SearchParams)
	p.Flavors = make(map[string]int)
	return p
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
			if f.Int() > 0 {
				v.Add(name, strconv.FormatInt(f.Int(), 10))
			}
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
