package radix

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Convert(num interface{}, currentRadix, targetRadix int) (result string, err error) {

	var numStr string
	switch num.(type) {
	case string:
		numStr = num.(string)
	case int:
		numStr = strconv.Itoa(num.(int))
	case int64:
		numStr = strconv.FormatInt(num.(int64), 10)
	default:
		return "", errors.New("No support type")
	}

	if currentRadix <= 36 {
		numStr = strings.ToLower(numStr)
	}

	if !check(numStr, currentRadix) {
		return "", errors.New("current radix is error")
	}

	ten := enten(numStr, currentRadix)
	result = deten(ten, targetRadix)
	return
}

func check(num string, r int) bool {
	s := char[:r]
	for _, v := range num {
		if strings.IndexRune(s, v) < 0 {
			return false
		}
	}
	return true
}

// N to 10
func enten(num string, r int) (result int) {
	for k, v := range strings.Split(reverse(num), "") {
		result += strings.Index(char, v) * int(math.Pow(float64(r), float64(k)))
	}
	return
}

// 10 to N
func deten(num int, r int) (result string) {
	for i := num; i > 0; i /= r {
		result = string(char[i%r]) + result
	}
	return
}

// reverse
func reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
