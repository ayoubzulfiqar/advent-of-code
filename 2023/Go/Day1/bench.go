package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	SumCalibration()
	endTime := time.Since(startTime)
	fmt.Printf("MilliSeconds: %v\nMicroSeconds: %v", endTime.Milliseconds(), endTime.Microseconds())
}
