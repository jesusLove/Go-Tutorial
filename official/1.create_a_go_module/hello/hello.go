package main

import (
	"fmt"
	"itlee.top/greetings"
	"log"
)
func main() {
	// 设置 log 前缀
	log.SetPrefix("greetings:")
	// 设置 log 表示：隐藏日期、时间等
	log.SetFlags(0)

	names := []string{"lee", "wang", "zhang"}
	messages, err := greetings.Hellos(names)
	// 如果返回错误，打印到控制台并退出程序。
	if err != nil {
		// 等价于：执行 Print 然后执行 Exit
		log.Fatal(err)
	}

	fmt.Println(messages)
}