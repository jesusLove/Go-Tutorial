package main

import (
	"fmt"
	"time"
)

/**
	* 结构体：聚合的数据类型，由零个或多个任意类型的值聚合的实体。
	*
 */

// Employee
// * 首字母大写标识：成员导出的
// * 成员顺序有意义，不同的结构体类型
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

func main() {
	// 2. 结构体字面值: 赋值给一个结构体指针
	e1 := &Employee{10, "lee", "Qingdao", time.Now(), "ShanDong", 5000, 20}
	fmt.Printf("el = %v\n", e1)

	// 3. 读取结构体中属性值
	fmt.Printf("e1.ID = %d\n", e1.ID)

	// 4. 修改值
	e1.Address = "Qingdao LaoShan"
	fmt.Printf("e1.Address = %s\n", e1.Address)
}
