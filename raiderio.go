package raideriogo

import(
	"encoding/json"
	"net/http"
	"io/ioutil"
)

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
