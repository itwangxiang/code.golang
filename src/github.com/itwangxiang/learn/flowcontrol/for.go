package main

import "fmt"

func main() {
	sum := 0
	//Go 只有一种循环结构—— for 循环
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	sum1 := 1
	//循环初始化语句和后置语句都是可选的
	for ; sum1 < 1000; {
		sum1 += sum1
	}
	fmt.Println(sum1)


	sum2 := 1
	//for 是 Go 的 “while”
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
