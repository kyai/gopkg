package ip

import (
	"net/http"
)

type Info struct {
	Ip      string
	Country string
	Area    string
	Region  string
	City    string
	Isp     string
}

type infoTaobao struct {
	code int
	data struct {
		ip      string
		country string
		area    string
		region  string
		city    string
		county  string
		isp     string
	}
}

func GetInfo(ip string) (info Info, err error) {
	return
}

func getInfoFromTaobao(ip string) (info infoTaobao, err error) {
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }
	return
}
