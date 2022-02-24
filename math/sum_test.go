package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInt(t *testing.T) {
	cases := []struct{
		x, y, result int
	}{
		{0, 0, 0},
		{1, 1, 2},
		{-1, -1, -2},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumInt(testCase.x, testCase.y),
		)
	}
}

func TestSumInt8(t *testing.T) {
	cases := []struct{
		x ,y   int8
		result int16
	}{
		{-128, -128, -256},
		{-127, 127, 0},
		{127, 127, 254},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumInt8(testCase.x, testCase.y),
		)
	}
}

func TestSumInt16(t *testing.T) {
	cases := []struct{
		x, y   int16
		result int32
	}{
		{-32768, -32768, -65536},
		{-32767, 32767, 0},
		{32767, 32767, 65534},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumInt16(testCase.x, testCase.y),
		)
	}
}

func TestSumInt32(t *testing.T) {
	cases := []struct{
		x, y   int32
		result int64
	}{
		{-2147483648, -2147483648, -4294967296},
		{-2147483647, 2147483647, 0},
		{2147483647, 2147483647, 4294967294},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumInt32(testCase.x, testCase.y),
		)
	}
}

func TestSumInt64(t *testing.T) {
	cases := []struct{
		x, y, result int64
	}{
		// these are all special cases where the int64 overflows
		{9223372036854775807, 9223372036854775807, -2},
		{9223372036854775807, 1, -9223372036854775808},
		{9223372036854775807, 922337203685477580, -8301034833169298229},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumInt64(testCase.x, testCase.y),
		)
	}
}

func TestSumUint(t *testing.T) {
	cases := []struct{
		x, y, result uint
	}{
		{0, 0, 0},
		{1, 1, 2},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumUint(testCase.x, testCase.y),
		)
	}
}

func TestSumUint8(t *testing.T) {
	cases := []struct{
		x, y   uint8
		result uint16
	}{
		{127, 127, 254},
		{255, 255, 510},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumUint8(testCase.x, testCase.y),
		)
	}
}

func TestSumUint16(t *testing.T) {
	cases := []struct{
		x, y   uint16
		result uint32
	}{
		{32767, 32767, 65534},
		{65535, 65535, 131070},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumUint16(testCase.x, testCase.y),
		)
	}
}

func TestSumUint32(t *testing.T) {
	cases := []struct{
		x, y   uint32
		result uint64
	}{
		{2147483647, 2147483647, 4294967294},
		{4294967295, 4294967295, 8589934590},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumUint32(testCase.x, testCase.y),
		)
	}
}

func TestSumUint64(t *testing.T) {
	cases := []struct{
		x, y, result uint64
	}{
		// these are all special cases where the uint64 overflows
		{18446744073709551615, 18446744073709551615, 18446744073709551614},
		{18446744073709551615, 1, 0},
		{18446744073709551615, 1844674407370955161, 1844674407370955160},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumUint64(testCase.x, testCase.y),
		)
	}
}

func TestSumFloat32(t *testing.T) {
	cases := []struct{
		x, y, result float32
	}{
		{1.2345678, 1.2345678, 2.4691356},
		{6.1234567, 6.1234567, 12.246913},
		{61.234567, 61.234567, 122.46913},
		{612.34567, 612.34567, 1224.6913},
		{6123.4567, 6123.4567, 12246.913},
		{61234.567, 61234.567, 122469.13},
		{612345.67, 612345.67, 1.2246914e+06},
		{6123456.7, 6123456.7, 1.2246913e+07},
		{61234567, 61234567, 1.2246914e+08},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumFloat32(testCase.x, testCase.y),
		)
	}
}

func TestSumFloat64(t *testing.T) {
	cases := []struct{
		x, y, result float64
	}{
		{1.234567890123, 1.234567890123, 2.469135780246},
		{6.234567890123, 6.234567890123, 12.469135780246},
		{62.34567890123, 62.34567890123, 124.69135780246},
		{623.4567890123, 623.4567890123, 1246.9135780246},
		{6234.567890123, 6234.567890123, 12469.135780246},
		{62345.67890123, 62345.67890123, 124691.35780246},
		{623456.7890123, 623456.7890123, 1.2469135780246e+06},
		{6234567.890123, 6234567.890123, 1.2469135780246e+07},
		{62345678.90123, 62345678.90123, 1.2469135780246e+08},
		{623456789.0123, 623456789.0123, 1.2469135780246e+09},
		{6234567890.123, 6234567890.123, 1.2469135780246e+10},
		{62345678901.23, 62345678901.23, 1.2469135780246e+11},
		{623456789012.3, 623456789012.3, 1.2469135780246e+12},
		{6234567890123, 6234567890123, 1.2469135780246e+13},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumFloat64(testCase.x, testCase.y),
		)
	}
}

func TestSumArrayInt(t *testing.T) {
	cases := []struct{
		x      []int
		result int
	}{
		{[]int{0, 0}, 0},
		{[]int{1, 1}, 2},
		{[]int{-1, -1}, -2},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayInt(testCase.x),
		)
	}
}

func TestSumArrayInt8(t *testing.T) {
	cases := []struct{
		x      []int8
		result int
	}{
		{[]int8{-128, -128}, -256},
		{[]int8{-127, 127}, 0},
		{[]int8{127, 127}, 254},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayInt8(testCase.x),
		)
	}
}

func TestSumArrayInt16(t *testing.T) {
	cases := []struct{
		x      []int16
		result int
	}{
		{[]int16{-32768, -32768}, -65536},
		{[]int16{-32767, 32767}, 0},
		{[]int16{32767, 32767}, 65534},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayInt16(testCase.x),
		)
	}
}

func TestSumArrayInt32(t *testing.T) {
	cases := []struct{
		x      []int32
		result int
	}{
		{[]int32{-2147483648, -2147483648}, -4294967296},
		{[]int32{-2147483647, 2147483647}, 0},
		{[]int32{2147483647, 2147483647}, 4294967294},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayInt32(testCase.x),
		)
	}
}

func TestSumArrayInt64(t *testing.T) {
	cases := []struct{
		x      []int64
		result int
	}{
		// these are all special cases where the int64 overflows
		{[]int64{9223372036854775807, 9223372036854775807}, -2},
		{[]int64{9223372036854775807, 1}, -9223372036854775808},
		{[]int64{9223372036854775807, 922337203685477580}, -8301034833169298229},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayInt64(testCase.x),
		)
	}
}

func TestSumArrayUint(t *testing.T) {
	cases := []struct{
		x      []uint
		result uint
	}{
		{[]uint{0, 0}, 0},
		{[]uint{1, 1}, 2},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayUint(testCase.x),
		)
	}
}

func TestSumArrayUint8(t *testing.T) {
	cases := []struct{
		x      []uint8
		result uint
	}{
		{[]uint8{127, 127}, 254},
		{[]uint8{255, 255}, 510},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayUint8(testCase.x),
		)
	}
}

func TestSumArrayUint16(t *testing.T) {
	cases := []struct{
		x      []uint16
		result uint
	}{
		{[]uint16{32767, 32767}, 65534},
		{[]uint16{65535, 65535}, 131070},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayUint16(testCase.x),
		)
	}
}

func TestSumArrayUint32(t *testing.T) {
	cases := []struct{
		x      []uint32
		result uint
	}{
		{[]uint32{2147483647, 2147483647}, 4294967294},
		{[]uint32{4294967295, 4294967295}, 8589934590},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayUint32(testCase.x),
		)
	}
}

func TestSumArrayUint64(t *testing.T) {
	cases := []struct{
		x      []uint64
		result uint
	}{
		// these are all special cases where the uint64 overflows
		{[]uint64{18446744073709551615, 18446744073709551615}, 18446744073709551614},
		{[]uint64{18446744073709551615, 1}, 0},
		{[]uint64{18446744073709551615, 1844674407370955161}, 1844674407370955160},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayUint64(testCase.x),
		)
	}
}

func TestSumArrayFloat32(t *testing.T) {
	cases := []struct{
		x      []float32
		result float32
	}{
		{[]float32{1.2345678, 1.2345678}, 2.4691356},
		{[]float32{6.1234567, 6.1234567}, 12.246913},
		{[]float32{61.234567, 61.234567}, 122.46913},
		{[]float32{612.34567, 612.34567}, 1224.6913},
		{[]float32{6123.4567, 6123.4567}, 12246.913},
		{[]float32{61234.567, 61234.567}, 122469.13},
		{[]float32{612345.67, 612345.67}, 1.2246914e+06},
		{[]float32{6123456.7, 6123456.7}, 1.2246913e+07},
		{[]float32{61234567, 61234567}, 1.2246914e+08},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayFloat32(testCase.x),
		)
	}
}

func TestSumArrayFloat64(t *testing.T) {
	cases := []struct{
		x      []float64
		result float64
	}{
		{[]float64{1.234567890123, 1.234567890123}, 2.469135780246},
		{[]float64{6.234567890123, 6.234567890123}, 12.469135780246},
		{[]float64{62.34567890123, 62.34567890123}, 124.69135780246},
		{[]float64{623.4567890123, 623.4567890123}, 1246.9135780246},
		{[]float64{6234.567890123, 6234.567890123}, 12469.135780246},
		{[]float64{62345.67890123, 62345.67890123}, 124691.35780246},
		{[]float64{623456.7890123, 623456.7890123}, 1.2469135780246e+06},
		{[]float64{6234567.890123, 6234567.890123}, 1.2469135780246e+07},
		{[]float64{62345678.90123, 62345678.90123}, 1.2469135780246e+08},
		{[]float64{623456789.0123, 623456789.0123}, 1.2469135780246e+09},
		{[]float64{6234567890.123, 6234567890.123}, 1.2469135780246e+10},
		{[]float64{62345678901.23, 62345678901.23}, 1.2469135780246e+11},
		{[]float64{623456789012.3, 623456789012.3}, 1.2469135780246e+12},
		{[]float64{6234567890123, 6234567890123}, 1.2469135780246e+13},
	}

	for _, testCase := range cases {
		assert.Equal(
			t,
			testCase.result,
			SumArrayFloat64(testCase.x),
		)
	}
}

// need to think up test cases for this
// func BenchmarkSumInt(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumInt8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumInt16(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumInt32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumInt64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumUint(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumUint8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumUint16(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumUint32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumUint64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumFloat32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumFloat64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayInt(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayInt8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayInt16(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayInt32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayInt64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayUint(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayUint8(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayUint16(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayUint32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayUint64(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayFloat32(b *testing.B) {}

// need to think up test cases for this
// func BenchmarkSumArrayFloat64(b *testing.B) {}
