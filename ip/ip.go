package ip

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Url []string

var urls Url

var TimeOut int = 5

func init() {
	urls = append(urls, "http://myexternalip.com/raw")
	urls = append(urls, "http://ipinfo.io/ip")
}

func (this *Url) Add(v string) {
	urls = append(urls, v)
}

func (this *Url) Set(list []string) {
	urls = list
}

func (this *Url) Get() Url {
	return urls
}

func Internet() (ip string, err error) {
	ipChan := make(chan string)
	for _, v := range urls {
		go httpget(v, ipChan)
	}

	select {
	case ip = <-ipChan:
		return
	case <-time.After(time.Second * time.Duration(TimeOut)):
		return "", errors.New("Timeout")
	}
}

func Internal() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return
}

func httpget(url string, ipChan chan string) (ip string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ip = string(data)
	ip = strings.Replace(ip, "\n", "", -1)
	ipChan <- ip
	return
}
