package curl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Do(method, url string, data map[string]interface{}, header map[string]string) (respData interface{}, err error) {

	if method == "" {
		err = errors.New("No method")
		return
	}
	method = strings.ToUpper(method)

	if header == nil {
		header = make(map[string]string)
		header["Content-Type"] = "application/json"
	}

	jsonS, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(jsonS)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return
	}

	for k, v := range header {
		request.Header.Set(k, v)
	}

	httpC := http.Client{}

	resp, err := httpC.Do(request)
	if err != nil {
		return
	}

	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	ctype := strings.ToLower(header["Content-Type"])

	if ctype == "application/json" {
		if err = json.Unmarshal(rdata, &respData); err != nil {
			return
		}
	} else if ctype == "application/xml" || ctype == "text/xml" {
		if err = xml.Unmarshal(rdata, &respData); err != nil {
			return
		}
	} else {
		respData = string(rdata)
	}

	return
}

func Get(url string, data map[string]interface{}, header map[string]string) (respData interface{}, err error) {
	return Do("Get", url, data, header)
}

func Post(url string, data map[string]interface{}, header map[string]string) (respData interface{}, err error) {
	return Do("Post", url, data, header)
}
