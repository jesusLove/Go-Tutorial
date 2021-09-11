package main

import "fmt"

func main()  {
	// 1. 数组定义，数组长度 len(a)
	var a[3]int
	fmt.Println(a[1])
	fmt.Println(a[len(a) - 1])
	// 2. 数组字面量
	b := [3]int{1,2,3}
	// 3. 遍历数组
	for i, v := range b {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}
	// 省略索引
	for _, v := range b {
		fmt.Printf("value = %d\n", v)
	}
	// 4. 数组比较
	c1 := [2]int{1,2}
	c2 := [...]int{1,2}
	c3 := [2]int{1,3}
	fmt.Println(c1 == c2, c1 == c3, c2 == c3)
	// 数组长度是类型的一部分，长度不同数组永远不相等。
	//c4 := [3]int{1,2}
	//fmt.Println(c1 == c4)

	// 5. 字面量指定索引和初始值
	d := [...]int{9:-1}
	for i, v := range d {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}

}
