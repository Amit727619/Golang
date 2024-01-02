package main

import (
	"fmt"
	"time"
)
func elapsed() {
	
	for i := 0; i < 20; i++ {
		_ = i
	}
}

func main() {
	
	starttime := time.Now()

	elapsed()

	endtime := time.Now()

	elapsedtime := endtime.Sub(starttime)

	fmt.Printf("Elapsed time %s", elapsedtime)
}