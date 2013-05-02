package goyum

import (
	"testing"
)

func getHandle(t *testing.T) *Yummly {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	y, err = SetCredentials(appid, appkey)
	if err != nil {
		t.Fatal(err)
	}
	return y
}

func TestDiet(t *testing.T) {
	y := getHandle(t)

	if diets, err := y.Diets(); err != nil {
		t.Error(err)
	} else {
		t.Logf("diets %+v\n", diets)
	}
}

func TestIngredient(t *testing.T) {
	y := getHandle(t)

	if ingredients, err := y.Ingredients(); err != nil {
		t.Error(err)
	} else {
		t.Logf("ingredients %+v\n", ingredients)
	}
}

func TestCourse(t *testing.T) {
	y := getHandle(t)

	if courses, err := y.Courses(); err != nil {
		t.Error(err)
	} else {
		t.Logf("courses %+v\n", courses)
	}
}
