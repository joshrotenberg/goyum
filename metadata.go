package goyum

import (
	"fmt"
	"net/url"
)

type MetaDataParams struct {
}

func (sp *MetaDataParams) values() url.Values {
	return make(url.Values)
}

func (sp *MetaDataParams) Encode() string {
	return "rad"
}

const (
	YummlyMetadataUrl = "http://meta.yummly.com"
)

func (y *Yummly) Ingredients() error {
	var i interface{}
	var p MetaDataParams
	err := y.callYummlyApi("GET", "metadata/ingredient", &p, &i)

	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", i)

	return nil
}
