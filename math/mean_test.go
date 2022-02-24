package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanInt(t *testing.T) {
	cases := []struct{
		x      []int
		result float64
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3.5},
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{-1, 1}, 0},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanInt(testCase.x),
		)
	}
}

func TestMeanInt8(t *testing.T) {
	cases := []struct{
		x      []int8
		result float64
	}{
		{[]int8{-128, 127}, -0.5},
		{[]int8{-128, -20, 30, 127}, 2.25},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanInt8(testCase.x),
		)
	}
}

func TestMeanInt16(t *testing.T) {
	cases := []struct{
		x      []int16
		result float64
	}{
		{[]int16{-32768, 32767}, -0.5},
		{[]int16{-32768, -1024, 5120, 32767}, 1023.75},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanInt16(testCase.x),
		)
	}
}

func TestMeanInt32(t *testing.T) {
	cases := []struct{
		x      []int32
		result float64
	}{
		{[]int32{-2147483648, 2147483647}, -0.5},
		{[]int32{-2147483648, -3072000, 7168000, 2147483647}, 1023999.75},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanInt32(testCase.x),
		)
	}
}

func TestMeanInt64(t *testing.T) {
	cases := []struct{
		x      []int64
		result float64
	}{
		{[]int64{-9223372036854775808, 9223372036854775807}, -0.5},
		{[]int64{-9223372036854775808, -5020039968768, 6044039968768, 9223372036854775807}, 255999999999.75},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanInt64(testCase.x),
		)
	}
}

func TestMeanUint(t *testing.T) {
	cases := []struct{
		x      []uint
		result float64
	}{
		{[]uint{1, 2, 3, 4, 5, 6}, 3.5},
		{[]uint{1024, 2048, 6144}, 3072},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanUint(testCase.x),
		)
	}
}

func TestMeanUint8(t *testing.T) {
	cases := []struct{
		x      []uint8
		result float64
	}{
		{[]uint8{128, 255}, 191.5},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanUint8(testCase.x),
		)
	}
}

func TestMeanUint16(t *testing.T) {
	cases := []struct{
		x      []uint16
		result float64
	}{
		{[]uint16{32767, 65535}, 49151},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanUint16(testCase.x),
		)
	}
}

func TestMeanUint32(t *testing.T) {
	cases := []struct{
		x      []uint32
		result float64
	}{
		{[]uint32{2147483647, 4294967295}, 3221225471},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanUint32(testCase.x),
		)
	}
}

func TestMeanUint64(t *testing.T) {
	cases := []struct{
		x      []uint64
		result float64
	}{
		{[]uint64{9223372036854775807, 18446744073709551615}, 4.611686018427388e+18},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanUint64(testCase.x),
		)
	}
}

func TestMeanFloat32(t *testing.T) {
	cases := []struct{
		x      []float32
		result float32
	}{
		{[]float32{1.2345678, 8.7654321}, 5},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanFloat32(testCase.x),
		)
	}
}

func TestMeanFloat64(t *testing.T) {
	cases := []struct{
		x      []float64
		result float64
	}{
		{[]float64{1.234567890123, 6.234567890123}, 3.7345678901230004},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			MeanFloat64(testCase.x),
		)
	}
}

// need to think up test cases for this
// func BenchmarkMeanInt(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanInt8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanInt16(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanInt32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanInt64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanUint(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanUint8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanUint16(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanUint32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanUint64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanFloat32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkMeanFloat64(b *testing.B) {}
