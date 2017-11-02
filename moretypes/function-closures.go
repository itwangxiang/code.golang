package main

import "fmt"

func main() {
	pos, neg := adder(), adder()

	fmt.Println(pos(1),pos(2))
	fmt.Println(pos(3))

	fmt.Println(neg(10),neg(11))
}

/**
函数的闭包
 */
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}