package main

import (
	"fmt"
	"github.com/Blxssy/DevOps-HW/cmd/calc"
)

func main() {
	res := calc.Add(2, 6)
	fmt.Println("Result:", res)

	res = calc.Sub(2, 6)
	fmt.Println("Result:", res)

	res = calc.Mul(2, 6)
	fmt.Println("Result:", res)
}
