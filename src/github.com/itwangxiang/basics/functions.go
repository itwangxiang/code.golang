package main

import "fmt"

//函数
func main() {
	fmt.Println(add(1,2))
	fmt.Println(add2(10,20))
}

/**
函数可以没有参数或接受多个参数
注意：类型在变量名之后
 */
func add(x int, y int) int {
	return x + y
}

/**
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
 */
func add2(x, y int) int {
	return x + y
}