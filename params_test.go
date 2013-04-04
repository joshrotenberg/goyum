package goyum

import (
	"testing"
	_ "fmt"
)

func TestSearchParams(t *testing.T) {
	sp := NewSearchParams()
	sp.RequirePictures(true).AddAllowedIngredients("garlic", "oranges")
	sp.AddExcludedIngredients("chalk", "bricks")
	t.Logf("%s", sp.Encode())
}
