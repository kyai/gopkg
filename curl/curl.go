package curl

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type curl struct {
	url    string
	data   interface{}
	body   BodyType
	param  map[string]string
	method string
	header map[string]string
	cookie map[string]string
}

type Response *http.Response

type BodyType int

const (
	_ BodyType = iota
	TEXT
	FORM
	JSON
	XML
)

// New request
func New(method, url string) *curl {
	return &curl{url: url, method: method}
}

// Set data and the type, default is JSON
func (this *curl) Data(data interface{}, body ...BodyType) *curl {
	this.data = data
	if len(body) > 0 {
		this.body = body[0]
	} else {
		this.body = JSON
	}
	return this
}

// Set params
func (this *curl) Param(param map[string]string) *curl {
	this.param = param
	return this
}

// Set headers
func (this *curl) Header(header map[string]string) *curl {
	this.header = header
	return this
}

// Set cookies
func (this *curl) Cookie(cookie map[string]string) *curl {
	this.cookie = cookie
	return this
}

// Combined urls and parameters
func (this *curl) urlWithParam() (err error) {
	if this.param == nil {
		return
	}

	var u *url.URL
	if u, err = url.Parse(this.url); err != nil {
		return
	}

	q := u.Query()
	for k, v := range this.param {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	this.url = u.String()

	return
}

func (this *curl) request() (response Response, err error) {
	if this.url == "" {
		return nil, errors.New("No url")
	} else {
		if err = this.urlWithParam(); err != nil {
			return
		}
	}
	if this.method == "" {
		return nil, errors.New("No method")
	} else {
		this.method = strings.ToUpper(this.method)
	}

	var payload io.Reader

	if this.data != nil && this.method != "GET" {
		switch this.body {
		case TEXT:
			payload = bytes.NewReader([]byte(this.data.(string)))
		case JSON:
			if jdata, err := json.Marshal(this.data); err != nil {
				return nil, err
			} else {
				payload = bytes.NewReader(jdata)
			}
		case XML:
			if xdata, err := xml.Marshal(this.data); err != nil {
				return nil, err
			} else {
				payload = bytes.NewReader(xdata)
			}
		}
	}

	request, err := http.NewRequest(this.method, this.url, payload)
	if err != nil {
		return
	}

	httpClient := http.Client{}

	response, err = httpClient.Do(request)
	if err != nil {
		return
	}

	return
}

func (this *curl) Do() (body []byte, header http.Header, err error) {
	var resp Response
	if resp, err = this.request(); err != nil {
		return
	}
	header = resp.Header
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
