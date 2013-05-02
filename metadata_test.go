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

func TestAllergies(t *testing.T) {
	y := getHandle(t)

	if allergies, err := y.Allergy(); err != nil {
		t.Error(err)
	} else {
		t.Logf("allergies %+v\n", allergies)
	}
}

func TestCuisines(t *testing.T) {
	y := getHandle(t)

	if cuisines, err := y.Cuisine(); err != nil {
		t.Error(err)
	} else {
		t.Logf("cuisines %+v\n", cuisines)
	}
}

func TestHoliday(t *testing.T) {
	y := getHandle(t)

	if holidays, err := y.Holiday(); err != nil {
		t.Error(err)
	} else {
		t.Logf("holidays %+v\n", holidays)
	}
}

func TestNutrition(t *testing.T) {
	y := getHandle(t)

	if nutrition, err := y.Nutrition(); err != nil {
		t.Error(err)
	} else {
		t.Logf("nutrition %+v\n", nutrition)
	}
}

func TestFlavor(t *testing.T) {
	y := getHandle(t)

	if flavors, err := y.Flavor(); err != nil {
		t.Error(err)
	} else {
		t.Logf("flavors %+v\n", flavors)
	}
}

