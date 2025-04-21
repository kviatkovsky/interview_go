package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T1 struct {
		x struct{}
		y int32
	}

	var t1 T1
	fmt.Println("t1 size:", unsafe.Sizeof(t1))
	// fmt.Println("t1 address:", unsafe.Pointer(&t1.x))

	type T2 struct {
		x int32
		y struct{}
	}

	var t2 T2
	fmt.Println("t2 size:", unsafe.Sizeof(t2))
}
