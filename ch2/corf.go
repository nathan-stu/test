package main

import "fmt"

type C float64
type F float64


func main() {
	fmt.Println(CtoF(55))
}

func CtoF(c C) F {
	return F(c*9/5 + 32)
}

func FtoC(f F) C {
	return C((f-32) * 5 /9)
}