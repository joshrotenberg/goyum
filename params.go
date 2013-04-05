package goyum

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
)

type SearchParams struct {
	q string
	v url.Values
}

func NewSearchParams(query string) *SearchParams {
	p := new(SearchParams)
	p.v = make(url.Values)
	p.v.Add("q", query)
	return p
}

// General parameters
func (sp *SearchParams) RequirePictures(r bool) *SearchParams {
	sp.v.Set("requirePictures", fmt.Sprintf("%t", r))
	return sp
}

func (sp *SearchParams) MaxTotalTimeInSeconds(seconds uint64) *SearchParams {
	sp.v.Set("maxTotalTimeInSeconds", strconv.FormatUint(seconds, 10))
	return sp
}

func (sp *SearchParams) MaxResult(results uint64) *SearchParams {
	sp.v.Set("maxResult", strconv.FormatUint(results, 10))
	return sp
}

func (sp *SearchParams) Start(start uint64) *SearchParams {
	sp.v.Set("start", strconv.FormatUint(start, 10))
	return sp
}

// Facets
func addFacet(v *url.Values, name string) {
	v.Add("facetField[]", name)
}

func (sp *SearchParams) AddFacetIngredient() *SearchParams {
	addFacet(&sp.v, "ingredient")
	return sp
}

func (sp *SearchParams) AddFacetDiet() *SearchParams {
	addFacet(&sp.v, "diet")
	return sp
}

// Ingredients
func addSliceToValues(v *url.Values, k string, s []string) {
	for _, s := range s {
		v.Add(k, s)
	}
}

func (sp *SearchParams) AddAllowedIngredients(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "allowedIngredient[]", is)
	return sp
}

func (sp *SearchParams) AddExcludedIngredients(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "excludedIngredient[]", is)
	return sp
}

// Diets
func (sp *SearchParams) AddAllowedDiets(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "allowedDiet[]", is)
	return sp
}

func (sp *SearchParams) AddExcludedDiets(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "excludedDiet[]", is)
	return sp
}

// Allergies
func (sp *SearchParams) AddAllowedAllergys(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "allowedAllergy[]", is)
	return sp
}

func (sp *SearchParams) AddExcludedAllergys(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "excludedAllergy[]", is)
	return sp
}

// Cuisine
func (sp *SearchParams) AddAllowedCuisines(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "allowedCuisine[]", is)
	return sp
}

func (sp *SearchParams) AddExcludedCuisines(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "excludedCuisine[]", is)
	return sp
}

// Courses
func (sp *SearchParams) AddAllowedCourses(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "allowedCourse[]", is)
	return sp
}

func (sp *SearchParams) AddExcludedCourses(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "excludedCourse[]", is)
	return sp
}

// Holidays
func (sp *SearchParams) AddAllowedHolidays(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "allowedHoliday[]", is)
	return sp
}

func (sp *SearchParams) AddExcludedHolidays(is ...string) *SearchParams {
	addSliceToValues(&sp.v, "excludedHoliday[]", is)
	return sp
}

func addAttribute(v *url.Values, class, att, minmax string, num float64) {
	var buf bytes.Buffer
	buf.WriteString(class)
	buf.WriteString(".")
	buf.WriteString(att)
	buf.WriteString(".")
	buf.WriteString(minmax)
	v.Add(buf.String(), strconv.FormatFloat(num, 'f', -1, 64))
}

// Flavors
func (sp *SearchParams) AddFlavorMin(flavor string, min float64) *SearchParams {
	addAttribute(&sp.v, "flavor", flavor, "min", min)
	return sp
}

func (sp *SearchParams) AddFlavorMax(flavor string, max float64) *SearchParams {
	addAttribute(&sp.v, "flavor", flavor, "max", max)
	return sp
}

// Nutrition
func (sp *SearchParams) AddNutritionMin(nutrient string, min float64) *SearchParams {
	addAttribute(&sp.v, "nutrition", nutrient, "min", min)
	return sp
}

func (sp *SearchParams) AddNutritionMax(nutrient string, max float64) *SearchParams {
	addAttribute(&sp.v, "nutrition", nutrient, "max", max)
	return sp
}

// Encode the string
func (sp *SearchParams) Encode() string {
	return sp.v.Encode()
}
