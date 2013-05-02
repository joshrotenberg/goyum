package goyum

import (
//"fmt"
)

type Diet struct {
	Id               string `json:"id"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	SearchValue      string `json:"searchValue"`
	Type             string `json:"type"`
}

type Ingredient struct {
	Id           string `json:"id"`
	Term         string `json:"term"`
	SearchValue  string `json:"searchValue"`
	IngredientId string `json:"ingredientId"`
	UseCount     int    `json:"useCount"`
}

type Course struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	SearchValue string `json:"searchValue"`
}

type Allergy struct {
	Id               string `json:"id"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	SearchValue      string `json:"searchValue"`
}

type Cuisine struct {
}

type Holiday struct {
}

type Nutrition struct {
}

type Flavor struct {
}

func getMetaData(y *Yummly, t string, d interface{}) error {

	sp := NewSearchParams("")
	err := y.callYummlyApi("GET", "metadata/"+t, sp, d)

	if err != nil {
		return err
	}
	return nil
}

func (y *Yummly) Diets() ([]Diet, error) {
	var diets []Diet

	if len(diets) == 0 {
		err := getMetaData(y, "diet", &diets)
		if err != nil {
			return nil, err
		}
	}
	return diets, nil
}

func (y *Yummly) Ingredients() ([]Ingredient, error) {
	var ingredients []Ingredient

	err := getMetaData(y, "ingredient", &ingredients)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (y *Yummly) Courses() ([]Course, error) {
	var courses []Course

	err := getMetaData(y, "course", &courses)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (y *Yummly) Allergy() ([]Allergy, error) {
	var allergies []Allergy

	err := getMetaData(y, "allergy", &allergies)
	if err != nil {
		return nil, err
	}
	return allergies, nil
}

func (y *Yummly) Cuisine() ([]Cuisine, error) {
	var cuisines []Cuisine

	err := getMetaData(y, "cuisine", &cuisines)
	if err != nil {
		return nil, err
	}
	return cuisines, nil
}

func (y *Yummly) Holiday() ([]Holiday, error) {
	var holidays []Holiday

	err := getMetaData(y, "holiday", &holidays)
	if err != nil {
		return nil, err
	}
	return holidays, nil
}

func (y *Yummly) Nutrition() ([]Nutrition, error) {
	var nutritions []Nutrition

	err := getMetaData(y, "nutrition", &nutritions)
	if err != nil {
		return nil, err
	}
	return nutritions, nil
}

func (y *Yummly) Flavor() ([]Flavor, error) {
	var flavors []Flavor

	err := getMetaData(y, "flavor", &flavors)
	if err != nil {
		return nil, err
	}
	return flavors, nil
}
