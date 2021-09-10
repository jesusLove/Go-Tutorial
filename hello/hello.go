package main

import (
	"fmt"
	"itlee.top/greetings" // 由于为发布本地无法找到 greeting 包。
)
// * 解决方法
// 将 itlee.top/greeting 重定向到本地使用 go mod edit 命令
// go mod edit -replace example.com/greetings=../greetings
// 在 hello 目录的命令提示符下，运行 go mod tidy 命令以同步 example.com/hello 模块的依赖项，
// 添加代码所需但尚未在模块中跟踪的依赖项。
func main() {
	message := greetings.Hello("lee")
	fmt.Println(message)
}