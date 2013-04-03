package goyum

import (
	"fmt"
	"testing"
)

func TestSearchRecipes(t *testing.T) {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	y, err = SetCredentials(appid, appkey)
	sp := NewSearchParams()
	sp.RequirePictures = true
	sp.AllowedIngredient = append(sp.AllowedIngredient, "garlic")
	sp.AllowedIngredient = append(sp.AllowedIngredient, "onion")
	sp.Flavors.SweetMin = 1
	fmt.Println(sp)
	if res, err := y.SearchRecipes("onion+soup", sp); err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", res.TotalMatchCount)
	}
}
