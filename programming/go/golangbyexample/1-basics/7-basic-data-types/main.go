package main

import (
	"fmt"
	"math"
)

func main() {
	// int
	fmt.Printf("int min - max: %d - %d\n", math.MinInt, math.MaxInt)
	// int8
	fmt.Printf("int8 min - max: %d - %d\n", math.MinInt8, math.MaxInt8)
	// int16
	fmt.Printf("int16 min - max: %d - %d\n", math.MinInt16, math.MaxInt16)
	// int32
	fmt.Printf("int32 min - max: %d - %d\n", math.MinInt32, math.MaxInt32)
	// int64
	fmt.Printf("int64 min - max: %d - %d\n", math.MinInt64, math.MaxInt64)

	// unsigned
	// uint
	fmt.Printf("uint min - max: %d - %d\n", uint(0), uint(math.MaxUint))
	// uint8
	fmt.Printf("uint8 min - max: %d - %d\n", 0, math.MaxUint8)
	// uint16
	fmt.Printf("uint16 min - max: %d - %d\n", 0, math.MaxUint16)
	// uint32
	fmt.Printf("uint32 min - max: %d - %d\n", 0, math.MaxUint32)
	// uint64
	fmt.Printf("uint64 min - max: %d - %d\n", 0, uint64(math.MaxUint64))
}
