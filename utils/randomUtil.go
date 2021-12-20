package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateRandomNum() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
}
