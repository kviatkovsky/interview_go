package main

import "fmt"

func main() {
	var x = []string{"A", "M", "C"}

	for i, s := range x {
		fmt.Println(i, s)
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}

	fmt.Println(x)
}
