package ip

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Infomation struct {
	Ip      string
	Country string
	Region  string
	City    string
	Isp     string
}

func Info(ip string) (info Infomation, err error) {
	return getInfoFromTaobao(ip)
}

func getInfoFromTaobao(ip string) (info Infomation, err error) {
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var infoTaobao struct {
		Code int `json:"code"`
		Data struct {
			Ip      string `json:"ip"`
			Country string `json:"country"`
			Area    string `json:"area"`
			Region  string `json:"region"`
			City    string `json:"city"`
			County  string `json:"county"`
			Isp     string `json:"isp"`
		} `json:"data"`
	}

	err = json.Unmarshal(data, &infoTaobao)
	if err != nil {
		return
	}
	if infoTaobao.Code != 0 {
		err = errors.New("request has error")
		return
	}

	info.Ip = infoTaobao.Data.Ip
	info.Country = infoTaobao.Data.Country
	info.Region = infoTaobao.Data.Region
	info.City = infoTaobao.Data.City
	info.Isp = infoTaobao.Data.Isp

	return
}
