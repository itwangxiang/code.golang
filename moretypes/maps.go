package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

/**
map
映射键到值

在 map m 中插入或修改一个元素：

m[key] = elem
获得元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键存在：

elem, ok = m[key]
如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值

 */
func main() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])


	var mm = make(map[string]string)
	mm["key1"] = "value1"
	mm["key2"] = "value2"
	fmt.Println(mm)
	fmt.Println(mm["key2"])
	delete(mm,"key1")
	fmt.Println(mm)
	v,ok := mm["key1"]
	fmt.Println("The value:", v, "Present?", ok)

}
