package main

import "fmt"

//命名返回值
func main() {
	fmt.Println(split(17))
}

// Go 的返回值可以被命名，
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
