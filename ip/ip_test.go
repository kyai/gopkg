package ip

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	internal, _ := Internal()
	fmt.Println(internal)
	internet, _ := Internet()
	fmt.Println(internet)

	info, err := Info(internet)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
