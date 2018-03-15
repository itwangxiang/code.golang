package main

import "fmt"

/**
结构体
一个结构体（ struct ）就是一个字段的集合
 */
func main() {
	fmt.Println(Vertex{1, 2})


	//结构体字段使用点号来访问
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//结构体字段可以通过结构体指针来访问
	v0 := Vertex{1, 2}
	p := &v0
	p.X = 1e9
	fmt.Println(v0)

	fmt.Println(v1, p1, v2, v3)
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p1  = &Vertex{1, 2} // 类型为 *Vertex
)

type Vertex struct {
	X int
	Y int
}