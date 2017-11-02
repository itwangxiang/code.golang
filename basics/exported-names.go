package main

import (
	"fmt"
	"math"
)

//导出名 - （类似公共的）

func main() {
	/**
	首字母大写的名称是被导出的（类似公共的）
	在导入包之后，你只能访问所导出的名字（即首字母大写的），
	 */
	//fmt.Println(math.pi) //错误的
	fmt.Println(math.Pi)
}