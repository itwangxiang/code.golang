package main

import (
	"math"
	"fmt"
)

//数据类型转化
/**
表达式 T(v) 将值 v 转换为类型 T
 */
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
