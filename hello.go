package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now().UnixNano();
	result := strconv.Itoa(fab(45))
	end := time.Now().UnixNano();
	elapseTime := strconv.FormatInt((end-start)/1e6,10)
	fmt.Printf("elapseTime:"+elapseTime+",result:"+result)
}

func fab(num int) int {
	if num < 2 {
		return 1
	}
	return fab(num-1)+fab(num-2)
}