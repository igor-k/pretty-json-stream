package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; ; i++ {
		fmt.Printf("{\"n\":%v,\"s\":\"abc_%v\"}\n", i, i)
		time.Sleep(2 * time.Second)
	}
}
