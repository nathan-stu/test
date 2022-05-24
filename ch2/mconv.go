package main

import (
	"bufio"
	"os"
	"fmt"
	"gopl/test/meter"
	"strconv"
)

func main() {
	files := os.Args[1:]
	var first, second string
	var val float64
	if len(files) == 0 {
		inputs := bufio.NewScanner(os.Stdin)
		inputs.Scan()
		first = inputs.Text()
		inputs.Scan()
		second = inputs.Text()
		inputs.Scan()
		//如何停止读取呢？
		val, _ = strconv.ParseFloat(inputs.Text(), 64)		
	}else {
		first = os.Args[1]
		second = os.Args[2]
		val, _ = strconv.ParseFloat(os.Args[3], 64)
	}

	if first == "m" && second == "cm" {
		fmt.Println( meter.MToCM( meter.M(val) ) )
	}
	//so as above
}