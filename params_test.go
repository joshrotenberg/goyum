package goyum

import (
	"net/url"
	"testing"
)

func TestSearchParams(t *testing.T) {
	sp := NewSearchParams("onion soup")
	sp.RequirePictures(true).MaxTotalTimeInSeconds(60).AddAllowedIngredients("garlic", "oranges")
	sp.AddExcludedIngredients("chalk", "bricks").AddFlavorMin("sweet", 0.123232).AddFlavorMax("meaty", 0.7)
	sp.AddNutritionMax("K", 3.5).AddNutritionMin("CHOCDF", 0)
	v, err := url.ParseQuery(sp.Encode())
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%+v", sp.Encode())
		t.Logf("%+v", v)
	}
}
