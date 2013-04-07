package goyum

import (
	"testing"
)

func TestIngredients(t *testing.T) {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	y, err = SetCredentials(appid, appkey)
	y.Diets()
}
