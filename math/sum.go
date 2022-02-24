// For all Sum_64 functions, write tests to see what happens if you
// return a value higher than the max allowed for that data type
// to see what happens

package math

func SumInt(x, y int) int {
	return x + y
}

func SumInt8(x, y int8) int16 {
	return int16(x) + int16(y)
}

func SumInt16(x, y int16) int32 {
	return int32(x) + int32(y)
}

func SumInt32(x, y int32) int64 {
	return int64(x) + int64(y)
}

func SumInt64(x, y int64) int64 {
	return x + y
}

func SumUint(x, y uint) uint {
	return x + y
}

func SumUint8(x, y uint8) uint16 {
	return uint16(x) + uint16(y)
}

func SumUint16(x, y uint16) uint32 {
	return uint32(x) + uint32(y)
}

func SumUint32(x, y uint32) uint64 {
	return uint64(x) + uint64(y)
}

func SumUint64(x, y uint64) uint64 {
	return x + y
}

func SumFloat32(x, y float32) float32 {
	return x + y
}

func SumFloat64(x, y float64) float64 {
	return x + y
}

func SumArrayInt(x []int) int {
	var sum int

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}

func SumArrayInt8(x []int8) int {
	var sum int

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return sum
}

func SumArrayInt16(x []int16) int {
	var sum int

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return sum
}

func SumArrayInt32(x []int32) int {
	var sum int

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return sum
}

func SumArrayInt64(x []int64) int {
	var sum int

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return sum
}

func SumArrayUint(x []uint) uint {
	var sum uint

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}

func SumArrayUint8(x []uint8) uint {
	var sum uint

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return sum
}

func SumArrayUint16(x []uint16) uint {
	var sum uint

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return sum
}

func SumArrayUint32(x []uint32) uint {
	var sum uint

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return sum
}

func SumArrayUint64(x []uint64) uint {
	var sum uint

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return sum
}

func SumArrayFloat32(x []float32) float32 {
	var sum float32

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}

func SumArrayFloat64(x []float64) float64 {
	var sum float64

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}