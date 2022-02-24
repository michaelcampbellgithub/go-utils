package math

func MeanInt(x []int) float64 {
	var sum int
	xLen := len(x)

	for i := 0; i < xLen; i++ {
		sum += int(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanInt8(x []int8) float64 {
	var sum int
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanInt16(x []int16) float64 {
	var sum int
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanInt32(x []int32) float64 {
	var sum int
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanInt64(x []int64) float64 {
	var sum int
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += int(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanUint(x []uint) float64 {
	var sum uint
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanUint8(x []uint8) float64 {
	var sum uint
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanUint16(x []uint16) float64 {
	var sum uint
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanUint32(x []uint32) float64 {
	var sum uint
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanUint64(x []uint64) float64 {
	var sum uint
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += uint(x[i])
	}

	return float64(sum) / float64(xLen)
}

func MeanFloat32(x []float32) float32 {
	var sum float32
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum / float32(xLen)
}

func MeanFloat64(x []float64) float64 {
	var sum float64
	xLen := len(x)

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum / float64(xLen)
}