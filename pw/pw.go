package pw

import (
	"errors"
	"math/rand"
	"time"
)

var (
	Len int
	Str string
)

func init() {
	level(0)
}

func level(i int) {
	switch i {
	case 0:
		Len = 6
		Str = char(0)
	case 1:
		Len = 8
		Str = char(0) + char(1)
	case 2:
		Len = 8
		Str = char(0) + char(1) + char(2)
	default: // level 3 and more
		Len = 10
		Str = char(0) + char(1) + char(2) + char(3)
	}
}

func char(i int) string {
	switch i {
	case 0:
		return "0123456789"
	case 1:
		return "abcdefghijklmnopqrstuvwxyz"
	case 2:
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 3:
		return "@#$%^&*+-=_"
	}
	return ""
}

func Create(args ...interface{}) (string, error) {
	var pwd string
	var err error
	if len(args) == 1 {
		l, _ := args[0].(int)
		level(l)
	} else {
		for _, arg := range args {
			switch arg.(type) {
			case int:
				Len, _ = arg.(int)
			case string:
				Str, _ = arg.(string)
			default:
				return pwd, errors.New("Type error")
			}
		}
	}
	pwd = CreatePwd(Len, Str)
	return pwd, err
}

func CreatePwd(l int, s string) (pwd string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		pwd += string([]byte(s)[r.Intn(len(s))])
	}
	return
}
