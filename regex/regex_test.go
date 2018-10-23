package regex

import (
	"fmt"
	"testing"
)

func TestV(t *testing.T) {
	fmt.Println(Verify("123@qq.com").IsEmail())
}
