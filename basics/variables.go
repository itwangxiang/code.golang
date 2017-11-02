package main

import "fmt"

//变量
func main() {
	var i int
	fmt.Println(i,c,python,java)


	var j,k,h = true,false,"no!"
	fmt.Println(a,b,j,k,h)

	//在函数中， := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
	//函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， := 结构不能使用在函数外
	e,d,g := true,false,"no"
	fmt.Println(e,d,g)
}

//var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面
//var 语句可以定义在包或函数级别
var c,python,java bool

//初始化变量
var a,b int = 1,2