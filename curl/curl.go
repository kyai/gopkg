package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Do(method, url string, data map[string]interface{}, obj interface{}) {

	jsonS, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	reader := bytes.NewReader(jsonS)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		fmt.Println(err.Error())
	}

	request.Header.Set("Content-Type", "application/json")

	httpC := http.Client{}

	resp, err := httpC.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	if err = json.Unmarshal(rdata, obj); err != nil {
		fmt.Println(err.Error())
	}
}

func Get(url string, data map[string]interface{}, obj interface{}) {
	Do("Get", url, data, obj)
}

func Post(url string, data map[string]interface{}, obj interface{}) {
	Do("Post", url, data, obj)
}
