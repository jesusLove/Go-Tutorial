package main

import "fmt"

func main() {
	// 1. 切片的定义
	s1 := []int{1,2,3}
	s2 := make([]int, 3)
	fmt.Printf("s1 = %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	// 2. 初始化指定切片容量
	s3 := make([]int, 5, 8)
	// 切片容量实际上代表它的底层数组的长度
	fmt.Printf("s3 = %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	// 3. 切片截取
	s4 := []int{1,2,3,4,5,6,7,8}
	// 下标区间 [3,6)
	s5 := s4[3:6]
	// 输出 [4,5,6]，容量为 5
	// 容量计算公式：底层数组长度 8 减去切片表达式起始索引 3，即 5。
	fmt.Printf("s5 = %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))

	// 4. 切片追加
	s6 := make([]int, 3, 5)
	s6 = append(s6, 10)
	fmt.Printf("s6 = %v, len=%d, cap=%d\n", s6, len(s6), cap(s6))
	// 如果继续添加，超出切片容量怎么办？
	s6 = append(s6, 11, 12, 13)
	// 此时 cap 容量为 10
	fmt.Printf("s6 = %v, len=%d, cap=%d\n", s6, len(s6), cap(s6))
}

/**
切片容量的增长？

Go 语言会在切片容量不足时进行扩容。原理：扩容并不是改变原来的切片而是生成一个容量更大的切片，
然后将原来的切片元素和新元素一起拷贝到新切片中。一般情况下，新切片的容量是原来切片的 2 倍。
当原切片的长度大于或等于 1024 时，切片会以 1.25 倍作为基数扩容。
 */
