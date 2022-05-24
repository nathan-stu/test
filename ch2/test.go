package main

import (
	"fmt"
	"gopl/ch2/tempconv"
)

func main() {
	var c tempconv.Celsius = 10
	fmt.Println(tempconv.CToK(c))
	fmt.Println(tempconv.KToC(300.15))
}