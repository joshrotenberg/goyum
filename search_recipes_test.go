package goyum

import (
	"testing"
)

func TestSearchRecipes(t *testing.T) {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	y, err = SetCredentials(appid, appkey)
	sp := new(SearchParams)
	sp.RequirePictures = true
	sp.AllowedIngredient = append(sp.AllowedIngredient, "garlic")
	sp.AllowedIngredient = append(sp.AllowedIngredient, "onion")
	if res, err := y.SearchRecipes("onion+soup", sp); err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", res.TotalMatchCount)
	}

}
