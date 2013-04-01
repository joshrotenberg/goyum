package goyum

import (
	"errors"
	"os"
	"testing"
)

func getTestingCredentials() (appid, appkey string, err error) {
	appid = os.Getenv("YUMMLY_APPID")
	appkey = os.Getenv("YUMMLY_APPKEY")

	if len(appid) == 0 {
		err = errors.New("Please set YUMMLY_APPID")
		return
	}

	if len(appkey) == 0 {
		err = errors.New("Please set YUMMLY_APPKEY")
		return
	}
	return
}

func TestSetCredentials(t *testing.T) {
	appid, appkey, err := getTestingCredentials()
	if err != nil {
		t.Fatal(err)
	}

	var y *Yummly
	if y, err = SetCredentials(appid, appkey); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%+v", y)
	}
	
}

