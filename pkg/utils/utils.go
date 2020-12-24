package utils

import (
	"bytes"
	"math/rand"
	"time"
)

func RandPassword(length int, upper, lower, digit bool) string {
	upperStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerStr := "abcdefghijklmnopqrstuvwxyz"
	digitStr := "0123456789"

	randPool := ""
	if upper {
		randPool += upperStr
	}

	if lower {
		randPool += lowerStr
	}

	if digit {
		randPool += digitStr
	}

	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result bytes.Buffer
	for i := 0; i < length; i++ {
		result.WriteByte(randPool[rr.Int63()%int64(len(randPool))])
	}

	return result.String()
}
