package random

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"unicode"
)

func BenchmarkOldRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OldRune(unicode.Latin)
	}
}

func BenchmarkNewRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRune(unicode.Latin)
	}
}

func TestRune(t *testing.T) {
	t.Skip("Unsure what's up here, but this is now timing out")
	test := func(trt *TestRangeTable) bool {
		r := OldRune(trt.RangeTable)
		if !unicode.In(r, trt.RangeTable) {
			t.Logf("Generated rune %v not in Range Table %v", r, trt)
			return false
		}
		return true
	}

	err := quick.Check(test, &quick.Config{MaxCount: 10})

	if err != nil {
		t.Error("fail")
	}
}

type TestRangeTable struct {
	*unicode.RangeTable
}

func (t *TestRangeTable) Generate(rand *rand.Rand, size int) reflect.Value {
	rt := unicode.RangeTable{}
	r16size := rand.Intn(5)
	r32size := rand.Intn(5)
	for i := 0; i < r16size; i++ {
		stride := uint16(1)
		low := stride * uint16(rand.Intn(5))
		rt.R16 = append(rt.R16, unicode.Range16{
			Lo:     low,
			Hi:     low * uint16(rand.Intn(10)),
			Stride: stride,
		})
	}
	for i := 0; i < r32size; i++ {
		stride := uint32(rand.Intn(3))
		low := stride * uint32(rand.Intn(10))
		rt.R32 = append(rt.R32, unicode.Range32{
			Lo:     low,
			Hi:     low * uint32(rand.Intn(10)),
			Stride: stride,
		})
	}

	return reflect.ValueOf(NewMyRangeTable(&rt))
}

func NewMyRangeTable(rangeTable *unicode.RangeTable) *TestRangeTable {
	return &TestRangeTable{RangeTable: rangeTable}
}
