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

# Go 语法


## 数组 Array

定义一个数组
```
var a [5]int
```

初始化数组
```
var a = [5]int{1,2,3,4,5}
a := [5]int{1,2,3,4,5}

var = [...]int{1,2,3,4,5,6}
a := [...]int{1,2,3,4,5,6}

a := [5]int{1:1,3:3}
```

访问数组
```
var a [3]int
fmt.Println(a[0])

a := [3]int{}
println(len(a), cap(a))

```

数组遍历
```
for i := 0; i < len(a); i ++ {
	mft.Print(a[i], "\t")
} 
fmt.Println()


for _, value := range a {
	fmt.print(value)
} 
```

多维数组
```
var a [][]int
```

初始化二维数组
```
a =[4][2]int{
	{10,11}
	{20,21}
	{30,31}
	{40，41}
}

a = [3][2]int{
	1:{20, 21}
	3：{40, 41}
}

a = [3][2]int{
	1: {0: 20}
	3: {2: 42}
}

```

## 切片

定义一个切片
```
var a []string
a := []string
a := make([]string, 10)
```

切片初始化
```
a := arr[:5]

a := arr[1:]

a := arr[O:5]

s := make([]int, len, cap)
```

appand与copy
```
var a = []int{1,2,3}
a := appand(a, 4,5,6)

a := []int{1,2,3,4,5}
b := []int{6,7,8}
copy(a,b)
```

## Map (集合)

定义一个Map
```
a := make(map[string]int)

//指定长度
a := make(map[string]int, 8)

a := map[string]int{
	"name": 1,
	"age": 2,
}
```

获取
```
a1 := a["apple"]
a2, ok := a["pear"]

len := len(a)
```

修改
```
a["apple"] = 5
```

遍历Map
```
for k, v := range a {
	fmt.Printf("key=%s, value=%d\n, k, v")
}
```

删除
```
delete(m, "apple")
```

## 函数(function)

定义一个函数
```
func a([parameter list]) [return_types]{
	函数体
}

func a() {
	a(1,2,3,4,5)
}
```

不定长变参特性
```
正确
func a(a string, b ...int)

错误
func a(b ...int, a string)
```

匿名函数
```
func main{a := func(){
	fmt.Println("匿名函数")
}
	a() //调用
}
```

闭包
```
func a(x int) func(int) int{
	fmt.Println("p\n",&x)
	return func (y int) int{
		fmt.Println("%p\n", &x)
        fmt.Println(x)
        fmt.Println(y)
    
        return x + y
	}
}

func main() {
  f := closure(10);
  fmt.Println(f(1))
  fmt.Println(f(2))
}
```

defer(推迟)
```
func main() {
  fmt.Println("a")
  
  defer fmt.Println("b")
  defer fmt.Println("c")
}
```

panic
```
func main() {
  A()
  B()
  C()
}

func A() {
    fmt.Println("FUNC A")
}

func B() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recover is B")

        }
        
    }()

    panic("B panic")

}

func C() {
    fmt.Println("FUNC C")
}
```

## 结构体

定义一个结构体
```
type Aa struct {
	name string
	age int
}
var a1 Aa
a1.name = "xiaming"
a1.age = 18
```

匿名结构体
```
var Aa sturct {
	name string
	age int
}
Aa.name = "xiaoming"
Aa.age = 18


func main() {

  a := struct {
      Name string
      Age int
  }{
      Name:"Corwien",
      Age: 20,
  }

  fmt.Println(a)

}
```

匿名字段
```
type person struct{
    string
    int
}

func main() {

	//匿名字段必须严格遵守字段类型声明的顺序
    a := person{"Corwien", 12}
    
    fmt.Println(a)

}
```

指针
```
type person struct{
    Name string
    Age int
}

func main() {

  a := person{  //&person
      Name:"Jam",
      Age:19,
  }
  // a.Age = 20
  fmt.Println(a)

  A(a)

  // 值拷贝，如果需要原来的改变，则需要添加指针
  B(&a) //B(a)

  fmt.Println(a)

}

// per 为变量名，person表示为结构体类型
func A(per person) { //*person
    per.Age = 25
    fmt.Println("A", per)
}

// per 为变量名，person表示为结构体类型
func B(per *person) {
    per.Age = 18
    fmt.Println("B", per)
}
```

new关键字
```
p2 := new(a)
```

结构体嵌套
```
type person struct{
    Name string
    Age int
    Contact struct {
        Phone, City string
        Code int           // 门牌号
    }
}

func main() {
    a := person{Name:"Corwien", Age:15}
    a.Contact.Phone = "10086"
    a.Contact.City = "Guangzhou"
    a.Contact.Code = 2007
  
    fmt.Println(a)
}
```

嵌入
```
type human struct {
    Sex int
}

type teacher struct {
    human           
    Name string
    Age int
}

type student struct {
    human
    Name string
    Age int
}

func main() {

    a := teacher{Name:"Corwien", Age:25, human: human{Sex: 1}}
    b := student{Name:"mark", Age:12, human: human{Sex: 1}}
    
    a.Name = "Jack"
    a.Age = 10
    // a.human.Sex = 0
    a.Sex = 0
    fmt.Println(a, b)
}
```

## 方法(method)

创建方法
```
type Test struct{}

// 无参数、无返回值
func (t Test) method0() {

}

// 单参数、无返回值
func (t Test) method1(i int) {

}

// 多参数、无返回值
func (t Test) method2(x, y int) {

}

// 无参数、单返回值
func (t Test) method3() (i int) {
    return
}

// 多参数、多返回值
func (t Test) method4(x, y int) (z int, err error) {
    return
}

// 无参数、无返回值
func (t *Test) method5() {

}

// 单参数、无返回值
func (t *Test) method6(i int) {

}

// 多参数、无返回值
func (t *Test) method7(x, y int) {

}

// 无参数、单返回值
func (t *Test) method8() (i int) {
    return
}

// 多参数、多返回值
func (t *Test) method9(x, y int) (z int, err error) {
    return
}
```

表达式
```
```

## 映射(Map)

创建映射
```
map[KeyType]ValueType
```

遍历
```
myColors := map[string]string{
}

for key, value := range myColors {
    fmt.Printf("Key: %s  Value: %s\n", key, value)
}
```

查找
```
for k,v := range capital{
    fmt.Println(k,v)
}

for k := range capital{
}

for _, v := range capital{  //无需将值改为匿名变量形式，忽略值即可。
}
```

delete(删除)
```
delete(map,键)
```

## 并发

goroutine
```
go 函数名( 参数列表 )
```

通道（channel）
```
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
   // 并把值赋给 v


ch := make(chan int)

通道缓冲区
ch := make(chan int, 100)
```

# MySQL

创建数据库
```
create database gotest;
```

切换到数据库

```
use gotest;
```

查看数据库下的表
```
show tables;
```

创建一个表
```
CREATE TABLE movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    created INT
);
```

查询表的所有数据

```
select * from movies;
```