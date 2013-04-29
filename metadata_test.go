package goyum

import (
	"sort"
	"testing"
)

func TestDiet(t *testing.T) {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	y, err = SetCredentials(appid, appkey)
	if diets, err := y.Diets(); err != nil {
		t.Error(err)
	} else {
		//t.Logf("diets is %+v\n", diets)
		found := sort.Search(len(diets), func(i int) bool {
			t.Log(diets[i])
			return diets[i].Id == "386"
		})
		t.Log(found)
		for i := 0; i < len(diets); i++ {
			shortDescription := diets[i].ShortDescription
			if shortDescription == "Vegan" {
				t.Logf("'%s'\n", shortDescription)
			}
		}
	}
}

func TestIngredient(t *testing.T) {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	y, err = SetCredentials(appid, appkey)
	if ingredients, err := y.Ingredients(); err != nil {
		t.Error(err)
	} else {
		t.Logf("first ingredient is %+v\n", ingredients[0])
		for i := 0; i < len(ingredients); i++ {
			useCount := ingredients[i].UseCount
			if useCount > 0 {
				t.Log(useCount)
			}
		}
	}
}
