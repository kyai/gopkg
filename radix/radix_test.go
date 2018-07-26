package radix

import (
	"fmt"
	"testing"
)

func TestRadix(t *testing.T) {
	result, _ := Convert("f", 16, 10)
	fmt.Println("result: ", result)
}
