package goyum

import (
	"fmt"
)

type Diet struct {
	Id               string    `json:"id"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	SearchValue      string `json:"searchValue"`
	Type             string `json:"type"`
}

func (y *Yummly) Diets() error {
	var diets []Diet

	sp := NewSearchParams("")
	err := y.callYummlyApi("GET", "metadata/diet", sp, &diets)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("nice %+v\n", diets)
	return nil
}

func (y *Yummly) Ingredients() error {
	var i interface{}

	sp := NewSearchParams("")
	err := y.callYummlyApi("GET", "metadata/ingredient", sp, &i)

	if err != nil {
		return err
	}
	fmt.Printf("nice %+v\n", i)
	return nil
}
