package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(prefix string, n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	str := fmt.Sprintf("%s%s", prefix, sb.String())
	return str
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
