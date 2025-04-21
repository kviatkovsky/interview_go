package main

import (
	"fmt"
	// "syscall"
	// "time"
)

// GOGC=off go run main.go -o big_allocation
// ps -a | grep big_alolocation
// vmmap PID

var data []byte

func main() {
	count := 0

	for {
		const GB = 1 << 30
		data = make([]byte, GB)

		// for idx := 0; idx < GB; idx += syscall.Getpagesize() {
		// 	data[idx] = 100
		// }

		fmt.Println("allocated GB:", count)
		count++

		// if count == 20 {
		// 	time.Sleep(time.Hour)
		// }
	}
}
