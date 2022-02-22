package math

import (
	"testing"
	"math/rand"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	cases := []struct {
		a, b, result int
	}{
		{1, 2, 1},
		{3, 2, 2},
		{4, 4, 4},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinInt(testCase.a, testCase.b),
		)
	}
}

func TestMinInt8(t *testing.T) {
	cases := []struct {
		a, b, result int8
	}{
		{-128, 0, -128},
		{0, 127, 0},
		{0, -128, -128},
		{0, 0, 0},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinInt8(testCase.a, testCase.b),
		)
	}
}

func TestMinInt16(t *testing.T) {
	cases := []struct {
		a, b, result int16
	}{
		{-32768, 0, -32768},
		{0, 32767, 0},
		{0, -32768, -32768},
		{32767, 32767, 32767},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinInt16(testCase.a, testCase.b),
		)
	}
}

func TestMinInt32(t *testing.T) {
	cases := []struct {
		a, b, result int32
	}{
		{-2147483648, 0, -2147483648},
		{0, 2147483647, 0},
		{0, -2147483648, -2147483648},
		{2147483647, 2147483647, 2147483647},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinInt32(testCase.a, testCase.b),
		)
	}
}

func TestMinInt64(t *testing.T) {
	cases := []struct {
		a, b, result int64
	}{
		{-9223372036854775808, 0, -9223372036854775808},
		{0, 9223372036854775807, 0},
		{0, -9223372036854775808, -9223372036854775808},
		{9223372036854775807, 9223372036854775807, 9223372036854775807},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinInt64(testCase.a, testCase.b),
		)
	}
}

func TestMinUint(t *testing.T) {
	cases := []struct {
		a, b, result uint
	}{
		{1, 2, 1},
		{3, 2, 2},
		{4, 4, 4},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinUint(testCase.a, testCase.b),
		)
	}
}

func TestMinUint8(t *testing.T) {
	cases := []struct {
		a, b, result uint8
	}{
		{0, 127, 0},
		{128, 255, 128},
		{127, 0, 0},
		{255, 255, 255},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinUint8(testCase.a, testCase.b),
		)
	}
}

func TestMinUint16(t *testing.T) {
	cases := []struct {
		a, b, result uint16
	}{
		{0, 32767, 0},
		{32768, 65535, 32768},
		{32767, 0, 0},
		{65535, 65535, 65535},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinUint16(testCase.a, testCase.b),
		)
	}
}

func TestMinUint32(t *testing.T) {
	cases := []struct {
		a, b, result uint32
	}{
		{0, 2147483647, 0},
		{2147483648, 4294967295, 2147483648},
		{2147483647, 0, 0},
		{4294967295, 4294967295, 4294967295},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinUint32(testCase.a, testCase.b),
		)
	}
}

func TestMinUint64(t *testing.T) {
	cases := []struct {
		a, b, result uint64
	}{
		{0, 9223372036854775807, 0},
		{9223372036854775808, 18446744073709551615, 9223372036854775808},
		{9223372036854775807, 0, 0},
		{18446744073709551615, 18446744073709551615, 18446744073709551615},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinUint64(testCase.a, testCase.b),
		)
	}
}

func TestMinFloat32(t *testing.T) {
	cases := []struct {
		a, b, result float32
	}{
		{0, 1.2345678, 0},
		{1.2345678, 0, 0},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinFloat32(testCase.a, testCase.b),
		)
	}
}

func TestMinFloat64(t *testing.T) {
	cases := []struct {
		a, b, result float64
	}{
		{0, 1.234567890123, 0},
		{1.234567890123, 0, 0},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MinFloat64(testCase.a, testCase.b),
		)
	}
}

func BenchmarkMinInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinInt(rand.Int(), rand.Int())
	}
}

// need to think up test cases for this
// func BenchmartMinInt8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMinInt16(b *testing.B) {}

func BenchmarkMinInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinInt32(rand.Int31(), rand.Int31())
	}
}

func BenchmarkMinInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinInt64(rand.Int63(), rand.Int63())
	}
}

// need to think up test cases for this
// func BenchmarkMinUint(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMinUint8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMinUint16(b *testing.B) {}

func BenchmarkMinUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinUint32(rand.Uint32(), rand.Uint32())
	}
}

func BenchmarkMinUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinUint64(rand.Uint64(), rand.Uint64())
	}
}

func BenchmarkMinFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinFloat32(rand.Float32(), rand.Float32())
	}
}

func BenchmarkMinFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinFloat64(rand.Float64(), rand.Float64())
	}
}
