package random

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func String16() string {
	return StringOfLength(16)
}

func StringOfLength(length int) string {
	b := strings.Builder{}
	runes := RunesFromRange(unicode.Latin)
	for i := 0; i < length; i++ {
		b.WriteRune(runes[Uint(len(runes))])
	}

	return b.String()
}

func Uint(max int) int {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rando.Intn(max)
}

func OneOf[T any](items ...T) T {
	return items[Uint(len(items))]
}
