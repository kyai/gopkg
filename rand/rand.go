package rand

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

type rval int

func New(size int) rval {
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(int(math.Pow10(size)))
	r *= int(math.Pow10(size - len(strconv.Itoa(r))))
	return rval(r)
}

func (v rval) Int() int {
	return int(v)
}

func (v rval) String() string {
	return strconv.Itoa(v.Int())
}
