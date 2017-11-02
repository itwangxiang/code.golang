package main

import "fmt"

//多值返回
func main() {
	a,b := swap("hello","world")
	fmt.Println(a,b)
}

//函数可以返回任意数的返回值
func swap(x, y string) (string, string) {
	return y,x
}