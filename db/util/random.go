package util

import (
	"math/rand"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func RandomFloat(minVal, maxVal float64) float64 {
	if minVal >= maxVal {
		panic("minVal must be less than maxVal") // Use panic for critical errors
	}
	return minVal + (rand.Float64() * (maxVal - minVal))
}

func RandomString(numOfChars int) string {
	var sb strings.Builder
	for i := 0; i < numOfChars; i++ {
		char := letters[rand.Intn(len(letters))]
		sb.WriteByte(char)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomCurrency() string {
	currencies := []string{"GBP", "G", "S", "K"}
	return currencies[rand.Intn(len(currencies))]
}
