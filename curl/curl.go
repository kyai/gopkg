package curl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Do(method, url string, data map[string]interface{}, obj interface{}) error {

	jsonS, err := json.Marshal(data)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(jsonS)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	httpC := http.Client{}

	resp, err := httpC.Do(request)
	if err != nil {
		return err
	}

	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(rdata, obj); err != nil {
		return err
	}

	return nil
}

func Get(url string, data map[string]interface{}, obj interface{}) {
	Do("Get", url, data, obj)
}

func Post(url string, data map[string]interface{}, obj interface{}) {
	Do("Post", url, data, obj)
}
