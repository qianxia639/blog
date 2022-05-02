package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateRandom() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
