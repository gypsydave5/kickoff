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

func RangeTableSize(tab *unicode.RangeTable) int {
	var count int
	for _, r16 := range tab.R16 {
		count += int((r16.Hi - r16.Lo) / r16.Stride)
	}
	for _, r32 := range tab.R32 {
		count += int((r32.Hi - r32.Lo) / r32.Stride)
	}
	return count
}

func RunesFromRange(tab *unicode.RangeTable) []rune {
	var res []rune
	for _, r16 := range tab.R16 {
		for c := r16.Lo; c <= r16.Hi; c += r16.Stride {
			res = append(res, rune(c))
		}
	}
	for _, r32 := range tab.R32 {
		for c := r32.Lo; c <= r32.Hi; c += r32.Stride {
			res = append(res, rune(c))
		}
	}
	return res
}
