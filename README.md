# gotest

## 最基本的 git 流程

1. 克隆一个 git 仓库
```
git clone https://github.com/hailun99/gotest.git
```

2. 切换到 gotest 目录

```
cd gotest
```

3. 用 vsc 打开当前目录

```
code
```

4. 提交代码

```
git add .
git commit -m "记笔记"
git push
```

5. 查看当前状态

```
git status
```

## 第一个 go 程序

1. 创建一个 main.go 文件

```go
package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello World!")
}
```

2. 初始化 go 项目

```
go mod init github.com/hailun99/gotest
go mod tidy
```

3. 编译

```
go build
```

4. 执行编译出来的可执行文件

```
./gotest.exe
```