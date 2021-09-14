package main

import "fmt"

func main() {
	// 1. 定义 map
	// 1.1 make
	m1 := make(map[string]int)
	m1["age"] = 20
	fmt.Printf("m1 = %v\n", m1)
	// 1.2 字面量
	m2 := map[string]int{
		"age": 20,
		"name": 99,
		"sex": 1,
	}
	fmt.Printf("m2 = %v\n", m2)

	// 2. 取值
	fmt.Printf("m2[age] = %d\n", m2["age"])

	// 3. 遍历: 迭代顺序是随机的
	for key, value := range m2 {
		fmt.Printf("key = %s, value = %d\n", key, value)
	}

	// 4. 索引访问 map，空置判断
	age, ok := m2["age"]
	if !ok {
		fmt.Println("error, is not a key")
	}
	fmt.Printf("m2 age = %d\n", age)
}
