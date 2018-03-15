package main

import (
	"fmt"
	"strings"
)

/**
数组
 */
func main() {

	/**
	类型 [n]T 是一个有 n 个类型为 T 的值的数组
	 */
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)


	/**
	一个 slice 会指向一个序列的值，并且包含了长度信息
	 */
	fmt.Println("slice")
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	//len(s) 返回 slice s 的长度
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}


	/**
	slice 的 slice
	slice 可以包含任意的类型，包括另一个 slice
	 */
	// Create a tic-tac-toe board.
	fmt.Println("slice 的 slice")
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)


	/**
	slice 切片
	slice 可以重新切片，创建一个新的 slice 值指向相同的数组
	表达式 s[lo:hi]
	表示从 lo 到 hi-1 的 slice 元素，含前端，不包含后端。因此 s[lo:lo] 是空的，而 s[lo:lo+1] 有一个元素。
	 */
	fmt.Println("slice 切片")
	s1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s1 ==", s1)
	fmt.Println("s1[1:4] ==", s1[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s1[:3] ==", s1[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s1[4:] ==", s1[4:])



	/**
	构造 slice
	slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组
	 */
	fmt.Println("构造 slice")
	aa := make([]int, 5)
	printSlice("aa", aa)
	bb := make([]int, 0, 5)
	printSlice("bb", bb)
	cc := bb[:2]
	printSlice("cc", cc)
	dd := cc[2:5]
	printSlice("dd", dd)


	/**
	nil slice
	slice 的零值是 nil 。
	一个 nil 的 slice 的长度和容量是 0
	 */
	fmt.Println("nil slice")
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}


	/**
	向 slice 添加元素
	func append(s []T, vs ...T) []T
	 */
	fmt.Println("向 slice 添加元素")
	var aaa []int
	printSlice("aaa", aaa)

	// append works on nil slices.
	aaa = append(aaa, 0)
	printSlice("aaa", aaa)

	// the slice grows as needed.
	aaa = append(aaa, 1)
	printSlice("aaa", aaa)

	// we can add more than one element at a time.
	aaa = append(aaa, 2, 3, 4)
	printSlice("aaa", aaa)
}


func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}


func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}