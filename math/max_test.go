package math

import (
	"testing"
	"math/rand"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	cases := []struct {
		a, b, result int
	}{
		{1, 2, 2},
		{3, 2, 3},
		{4, 4, 4},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxInt(testCase.a, testCase.b),
		)
	}
}

func TestMaxInt8(t *testing.T) {
	cases := []struct {
		a, b, result int8
	}{
		{-128, 0, 0},
		{0, 127, 127},
		{0, -128, 0},
		{0, 0, 0},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxInt8(testCase.a, testCase.b),
		)
	}
}

func TestMaxInt16(t *testing.T) {
	cases := []struct {
		a, b, result int16
	}{
		{-32768, 0, 0},
		{0, 32767, 32767},
		{0, -32768, 0},
		{32767, 32767, 32767},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxInt16(testCase.a, testCase.b),
		)
	}
}

func TestMaxInt32(t *testing.T) {
	cases := []struct {
		a, b, result int32
	}{
		{-2147483648, 0, 0},
		{0, 2147483647, 2147483647},
		{0, -2147483648, 0},
		{2147483647, 2147483647, 2147483647},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxInt32(testCase.a, testCase.b),
		)
	}
}

func TestMaxInt64(t *testing.T) {
	cases := []struct {
		a, b, result int64
	}{
		{-9223372036854775808, 0, 0},
		{0, 9223372036854775807, 9223372036854775807},
		{0, -9223372036854775808, 0},
		{9223372036854775807, 9223372036854775807, 9223372036854775807},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxInt64(testCase.a, testCase.b),
		)
	}
}

func TestMaxUint(t *testing.T) {
	cases := []struct {
		a, b, result uint
	}{
		{1, 2, 2},
		{3, 2, 3},
		{4, 4, 4},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxUint(testCase.a, testCase.b),
		)
	}
}

func TestMaxUint8(t *testing.T) {
	cases := []struct {
		a, b, result uint8
	}{
		{0, 127, 127},
		{128, 255, 255},
		{127, 0, 127},
		{255, 255, 255},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxUint8(testCase.a, testCase.b),
		)
	}
}

func TestMaxUint16(t *testing.T) {
	cases := []struct {
		a, b, result uint16
	}{
		{0, 32767, 32767},
		{32768, 65535, 65535},
		{32767, 0, 32767},
		{65535, 65535, 65535},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxUint16(testCase.a, testCase.b),
		)
	}
}

func TestMaxUint32(t *testing.T) {
	cases := []struct {
		a, b, result uint32
	}{
		{0, 2147483647, 2147483647},
		{2147483648, 4294967295, 4294967295},
		{2147483647, 0, 2147483647},
		{4294967295, 4294967295, 4294967295},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxUint32(testCase.a, testCase.b),
		)
	}
}

func TestMaxUint64(t *testing.T) {
	cases := []struct {
		a, b, result uint64
	}{
		{0, 9223372036854775807, 9223372036854775807},
		{9223372036854775808, 18446744073709551615, 18446744073709551615},
		{9223372036854775807, 0, 9223372036854775807},
		{18446744073709551615, 18446744073709551615, 18446744073709551615},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxUint64(testCase.a, testCase.b),
		)
	}
}

func TestMaxFloat32(t *testing.T) {
	cases := []struct {
		a, b, result float32
	}{
		{0, 1.2345678, 1.2345678},
		{1.2345678, 0, 1.2345678},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxFloat32(testCase.a, testCase.b),
		)
	}
}

func TestMaxFloat64(t *testing.T) {
	cases := []struct {
		a, b, result float64
	}{
		{0, 1.234567890123, 1.234567890123},
		{1.234567890123, 0, 1.234567890123},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MaxFloat64(testCase.a, testCase.b),
		)
	}
}

func BenchmarkMaxInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxInt(rand.Int(), rand.Int())
	}
}

// need to think up test cases for this
// func BenchmartMaxInt8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMaxInt16(b *testing.B) {}

func BenchmarkMaxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxInt32(rand.Int31(), rand.Int31())
	}
}

func BenchmarkMaxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxInt64(rand.Int63(), rand.Int63())
	}
}

// need to think up test cases for this
// func BenchmarkMaxUint(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMaxUint8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMaxUint16(b *testing.B) {}

func BenchmarkMaxUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxUint32(rand.Uint32(), rand.Uint32())
	}
}

func BenchmarkMaxUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxUint64(rand.Uint64(), rand.Uint64())
	}
}

func BenchmarkMaxFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxFloat32(rand.Float32(), rand.Float32())
	}
}

func BenchmarkMaxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxFloat64(rand.Float64(), rand.Float64())
	}
}
