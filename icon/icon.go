package icon

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func Get(u string) (string, error) {
	var ico string
	// get real url
	u = strings.ToLower(u)
	// get domain
	uparse, err := url.Parse(u)
	if err != nil {
		return ico, err
	}
	domain := uparse.Scheme + "://" + uparse.Host

	var resp *http.Response

	// try to get it
	// if resp, err = http.Get(domain + "/favicon.ico"); err == nil && resp.StatusCode == 200 {
	// 	return domain + "/favicon.ico", err
	// }

	resp, err = http.Get(domain)
	if err != nil {
		fmt.Println(err)
		return ico, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("err2")
		return ico, err
	}

	str := string(b)
	reg := regexp.MustCompile(`(?i:<link ).+(?i:href\s*=\s*"|'\s*)(.+)("|')`)
	// regHref := regexp.MustCompile(``)
	// fmt.Printf("%q\n", reg.FindAllString(str, -1))
	for _, v := range reg.FindAllString(str, -1) {
		fmt.Println(v)
		// ico = regHref.FindStringSubmatch(str)[1]
	}
	return ico, err

	// return ico, err
}

func verifyUrl(url string) error {
	if resp, err := http.Get(url); err == nil && resp.StatusCode == 200 {
		return nil
	}
	return errors.New("Error")
}
