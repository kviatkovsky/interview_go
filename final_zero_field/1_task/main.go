package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T1 struct {
		x int8
		y int32
		z int8
	}

	var t1 T1
	fmt.Println("t1 size:", unsafe.Sizeof(t1))

	type T2 struct {
		x int32
		y int8
		z int8
	}

	var t2 T2
	fmt.Println("t2 size:", unsafe.Sizeof(t2))

	type T3 struct {
		x struct{}
		y struct{}
		z struct{}
	}

	var t3 T3
	fmt.Println("t3 size:", unsafe.Sizeof(t3))
}
