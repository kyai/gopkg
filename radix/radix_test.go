package radix

import (
	"fmt"
	"testing"
)

func TestRadix(t *testing.T) {
	result, err := Convert("1", 2, 10)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("result: ", result)
}
