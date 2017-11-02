package main

import "fmt"

const Pi = 3.14

/**
常量
 */
func main() {

	const World = "世界"
	fmt.Println("hello",World)
	fmt.Println("Happy",Pi,"Day")

	const Truth  =true
	fmt.Println("Go rules?",Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}


/**
数值常量是高精度的 值
 */
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}