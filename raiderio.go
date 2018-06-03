package raideriogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Get returns an interface corresponding to the given endpoint.
func Get(endpoint string) interface{} {
	var v interface{}

	response, err := http.Get(endpoint)
	if nil != err {
		return err.Error()
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err.Error()
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err.Error()
	}

	return v
}
