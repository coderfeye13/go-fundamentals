package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8       // 255
	var smallI int32 = math.MaxInt32 // 2147483647
	var bigI uint64 = math.MaxUint64 // 18446744073709551615

	fmt.Printf("Before: b=%d, smallI=%d, bigI=%d\n", b, smallI, bigI)

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Printf("After:  b=%d, smallI=%d, bigI=%d\n", b, smallI, bigI)
}
