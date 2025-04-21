package main

import "unsafe"

func main() {
	var str1 = "go"
	str2 := str1 + "-go"

	str1Data := unsafe.StringData(str1)
	str2Data := unsafe.StringData(str2)

	slice1 := unsafe.Slice(str1Data, len(str1))
	slice2 := unsafe.Slice(str2Data, len(str2))

	slice2[0] = 'G'
	println("str2: ", str2)

	slice1[0] = 'G'
	println("str1: ", str1)

}
