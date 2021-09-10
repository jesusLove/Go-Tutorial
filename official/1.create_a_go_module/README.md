# 创建 module

定义一个 greetings 文件夹。使用 `go mod init` 命令设置模块名：`itlee.top/greetings`。模块正式发布时，该路径就是 Go 工具下载模块的路径。

创建 greetings.go 文件，编写模块内容。

* 声明包名 `package greetings`，Go 语言通过代码包来管理代码。
* 实现 `Hello` 函数，其接受一个名为 name 字符串参数，返回一个字符串。

> 注意：Hello 首字母大写才能被外部访问。

```go
package greetings

import "fmt"

func Hello(name string) string  {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return  message
}
```

# 调用 module 

在 greetings 文件夹同级创建 hello 文件夹，同时初始化 mod 文件和创建 hello.go 文件。

**如何在 hello.go 文件中使用 greetings.go 定义的 Hello 函数？**

首先导入包 `import itlee.top/greetings`，此时会发现找不到 `greetings.Hello` 函数？ 
由于 greetings 只存在于本地并没有发布，go 工具无法下载和使用。所以需要配置下，让它在本地查找 `itlee.top/greetings` 包。

1. 使用 `go mod edit` 命令编辑 `itlee.top/hello` 模块，将获取 `greetings` 模块路径重定向到本地。
`go mod edit -replace itlee.top/greetings=../greetings`
2. 上面命令只是进行依赖重定向还未跟踪，需要使用 `go mod tidy` 命令跟踪依赖。

配置完成后，就可以在使用 `itlee.top/greetings` 模块。

```vgo
// hello/go.mod
module itlee.top/hello

go 1.16
// 执行 go mod edit ... 后生成
replace itlee.top/greetings => ../greetings
// 执行 go mod tidy 生成
require itlee.top/greetings v0.0.0-00010101000000-000000000000
```
> `v0.0.0-xxx ` 模块路径后面的数字是一个伪版本号——一个生成的数字，用来代替语义版本号。

```go
package main

import (
	"fmt"
	"itlee.top/greetings" // 由于为发布本地无法找到 greeting 包。
)
func main() {
	message := greetings.Hello("lee")
	fmt.Println(message)
}
```
# 返回和处理 Error

```go
// 省略...
func Hello(name string) (string, error)  {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return  message, nil
}
```
> Hello 函数返回两个值 (string, error) 。

在调用 Hello 函数时，需要接受两个返回值。

```go
func main() {
	// 设置 log 前缀
	log.SetPrefix("greetings:")
	// 设置 log 表示：隐藏日期、时间等
	log.SetFlags(0)

	message, err := greetings.Hello("")
	// 如果返回错误，打印到控制台并退出程序。
	if err != nil {
		// 等价于：执行 Print 然后执行 Exit
		log.Fatal(err)
	}

	fmt.Println(message)
}
```
代码很简单没有啥好说的。 `log` 是一个日志包, [log package 文档](https://pkg.go.dev/log)


# 返回一个随机问候

在 greetings.go 文件添加：

```go
import (
"math/rand"
"time"
)
// 初始化 rand
func init() {
	rand.Seed(time.Now().UnixNano())
}
// 返回随机字符串
func randomFormat() string {
	// 定义 slice
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
```

# 添加 Test 

```go
package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Lee"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Lee")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Lee") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
```