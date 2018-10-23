package rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(New(5).String())
	}
}
