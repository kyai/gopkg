package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	SetKey("6666666666666666")
	str, err := Encrypt("helloworld")
	if err != nil {
		panic(err)
	}
	fmt.Println("result: ", str)
}
