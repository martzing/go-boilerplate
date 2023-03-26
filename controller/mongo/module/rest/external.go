package demoRest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RandomNames() []string {
	resp, err := http.Get("https://names.drycodes.com/10?separator=space")

	if err != nil {
		panic(err)
	}

	result := []string{}

	respData, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(respData, &result)

	return result
}
