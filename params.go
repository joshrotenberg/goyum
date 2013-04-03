package goyum

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type Flavors struct {
	SweetMin   int
	SweetMax   int
	MeatyMax   int
	MeatyMin   int
	SourMin    int
	SourMax    int
	BitterMin  int
	BitterMax  int
	PiquantMin int
	PiquantMax int
}

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
	Flavors               *Flavors
}

func NewSearchParams() *SearchParams {
	p := new(SearchParams)
	p.Flavors = new(Flavors)
	return p
}

func (f *Flavors) values() url.Values {
	v := make(url.Values)
	return v
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
		case reflect.Map:
			k := f.MapKeys()
			fmt.Println(k)
		case reflect.Struct:
			switch f.Type().String() {
			case "*goyum.Flavors":
				fmt.Println("its a Doof")
			default:
				fmt.Println("no idea")
			}
		case reflect.Ptr:
			fmt.Println("hits a " + f.Kind().String())
		default:
			fmt.Println("its a " + f.Type().String())
		}
	}

	return v
}
