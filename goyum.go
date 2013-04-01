package goyum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	YummlyApiUrl     = "http://api.yummly.com/"
	YummlyApiUrlSSL  = "https://api.yummly.com/"
	YummlyApiVersion = "v1"
)

var (
	ErrorNotJSON = errors.New("Yummly API did not return JSON")
)

type Yummly struct {
	httpClient *http.Client
	AppId      string
	AppKey     string
}

type Params interface {
	values() url.Values
}

func SetCredentials(appId string, appKey string) (*Yummly, error) {
	client := &http.Client{}
	return &Yummly{httpClient: client, AppId: appId, AppKey: appKey}, nil

}

func call(y *Yummly, method, uri, params string) (response *http.Response, err error) {
	url := fmt.Sprintf("%s/%s/%s?%s", YummlyApiUrl, YummlyApiVersion, uri, params)

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	request.Header.Set("X-Yummly-App-ID", y.AppId)
	request.Header.Set("X-Yummly-App-Key", y.AppKey)
	response, err = y.httpClient.Do(request)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		e := "Yummly responded with %d"
		err = errors.New(fmt.Sprintf(e, response.StatusCode))
	}
	return
}

func (y *Yummly) callYummlyApi(method, uri string, params Params, result interface{}) error {
	response, err := call(y, method, uri, params.values().Encode())
	if err != nil {
		return err
	}

	switch response.Header.Get("Content-Type") {
	case "application/json":
		fallthrough
	case "application/json; charset=utf-8":
		var js []byte
		js, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(js, result)
		}
	default:
		err = ErrorNotJSON
	}
	return err
}
