package random

import (
	"fmt"
	"unicode"
)

func OldRune(table *unicode.RangeTable) rune {
	runes := RunesFromRange(table)
	return runes[Uint(len(runes))]
}

func NewRune(table *unicode.RangeTable) rune {
	return RuneFromRangeTable(table, Uint(RangeTableSize(table)))
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

func Range32Size(r unicode.Range32) int {
	return int((r.Hi - r.Lo) / r.Stride)
}

func Range16Size(r unicode.Range16) int {
	return int((r.Hi - r.Lo) / r.Stride)
}

func RuneFromRange16(r unicode.Range16, index int) rune {
	return rune(r.Lo + (uint16(index) * r.Stride))
}

func RuneFromRange32(r unicode.Range32, index int) rune {
	return rune(r.Lo + (uint32(index) * r.Stride))
}

func RuneFromRangeTable(tab *unicode.RangeTable, index int) rune {
	i := index
	for _, r16 := range tab.R16 {
		if i <= int(Range16Size(r16)) {
			return RuneFromRange16(r16, i)
		}
		i -= Range16Size(r16)
	}
	for _, r32 := range tab.R32 {
		if i <= int(Range32Size(r32)) {
			return RuneFromRange32(r32, i)
		}
		i -= Range32Size(r32)
	}

	panic(fmt.Sprintf("index %v, final index: %v, range table size: %v", index, i, RangeTableSize(tab)))
}
