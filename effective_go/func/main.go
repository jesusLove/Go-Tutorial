package main

import "fmt"

// 函数
// Go 中有三种类型的函数：普通函数、匿名函数、方法
// 函数参数传递方式：
// 1. 按值传递：默认使用的是按值传递
// 2. 按引用传递：如果想传递引用，需使用 & 取地址符，此时传递的是指针值。
func main() {
	// 验证：函数签名
	fmt.Printf("add = %T\n", add) // func(int, int) int
	fmt.Printf("add2 = %T\n", add2) // func(int, int) int

	testMin()
}

// 1. 函数定义
// 形参相同类型可以简写；
// 函数签名相同: 如果形参列表和返回值列表中的变量类型一一对应，那么函数被认为签名相同。
func add(a int, b int) int { return  a + b }
func add2(a, b int) int { return  a + b }

// 2. 返回值：命名返回值和非命名返回值
func get2And3(input int) (int, int)  {
	return 2 * input, 3 * input
}
func get2And3_2(input int) (x2, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

// 3. 空白符
// 如上面 get2And3 如果我只想取第二个数值
// 第一个数值会被舍弃
func blankIdentifier() {
	_, x3 := get2And3(12)
	fmt.Printf("x3 = %d\n", x3)
}
// 4. 改变函数外部变量
// 不能直接传值，需要传指针

// 5. 变长参数
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
func testMin()  {
	x := min( 3, 1, 5, 7)
	fmt.Printf("Min = %d\n", x)
	// 简化 slice 作为参数
	s := []int{7,9,3,5,11}
	x = min(s...)
	fmt.Printf("Min = %d\n", x)
}
// 5-1 变长参数的类型不同？
// 1. 使用结构体
// 2. 使用空接口

