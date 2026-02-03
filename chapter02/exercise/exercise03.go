package main

import (
	"fmt"
	"math"
)

func exercise03() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI int64 = math.MaxInt64

	b++
	smallI++
	bigI++

	fmt.Println(b, smallI, bigI) // 0 -2147483648 -9223372036854775808
}
