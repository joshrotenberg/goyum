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
	sp := NewSearchParams("onion soup")
	sp.RequirePictures(true).MaxTotalTimeInSeconds(60).AddAllowedIngredients("garlic", "onions")
	sp.AddExcludedIngredients("chalk", "bricks").AddFlavorMin("sweet", 0.123232).AddFlavorMax("meaty", 0.7)
	sp.AddNutritionMax("K", 3.5).AddNutritionMin("CHOCDF", 0).AddAllowedCourses("course^course-Soups")
	sp.AddFacetIngredient()
	if res, err := y.SearchRecipes(sp); err != nil {
		fmt.Println(err)
		t.Error(err)
	} else {
		t.Logf(" res is %+v\n", res)
	}
}
