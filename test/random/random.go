package random

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomString() string {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(rando.Intn(1000)) // todo: better random string
}
