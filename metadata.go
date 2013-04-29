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

	err := getMetaData(y, "diet", &diets)
	if err != nil {
		return nil, err
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
