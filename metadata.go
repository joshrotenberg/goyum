package goyum

import (
	"fmt"
)

type Diet struct {
	Id               int
	ShortDescription string
	LongDescription  string
	SearchValue      string
}

type DietMetadata struct {
	Diets []*Diet
}

func (y *Yummly) Diets() error {
	//var d DietMetadata
	fmt.Println("word up")
	var i interface{}
	sp := NewSearchParams("")
	err := y.callYummlyApi("GET", "metadata/diet", sp, &i)

	if err != nil {
		return err
	}
	fmt.Println(i)
	//t.Logf("nice %+v\n", d)
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
