package regex

import (
	"regexp"
	"strconv"
)

type verify struct {
	s string
}

func Verify(v interface{}) *verify {
	var s string
	switch v.(type) {
	case string:
		s = v.(string)
	case int:
		s = strconv.Itoa(v.(int))
	}
	return &verify{s: s}
}

func (v *verify) IsUname() bool {
	pat := `^\D+[^@#\.]*$`
	return v.match(pat)
}

func (v *verify) IsPhone() bool {
	pat := `^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`
	return v.match(pat)
}

func (v *verify) IsEmail() bool {
	pat := `^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`
	return v.match(pat)
}

func (v *verify) match(pattern string) bool {
	if ok, err := regexp.MatchString(pattern, v.s); ok && err == nil {
		return true
	}
	return false
}
