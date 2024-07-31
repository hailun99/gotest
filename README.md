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

## Mysql知识点
```
DDL:操作数据库的
DML：表的增删改查
DCL:用户及权限
```

 存储引擎
```
MySQL支持插件式的存储引擎.
常见的存储引擎:MylSAM和lnnoDB。
```

MyLSAM:
```
1.查询速度快
2.只支持表锁
3.不支持事务
```

lnnoDB
```
1.整体速度快
2.支持表锁和行锁
3.支持事务
```

事务
```
把多个SQL操作当成一个整体

事务的特点:
ACID:
    1.原子性:事务要么成功要么失败,没有中间状态.
    2.一致性:数据的完整性没有被破坏
    3.隔离性：事务之间是相互隔离的.事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）
    4.持久性:事务操作的结果是不会丢失的.
```

## 创建数据库
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

插入数据
```
insert into user(id,title) values(1,哈哈哈哈)
```

查询表的所有数据

```
select * from movies;
```

单行查询
```
select title from user where id=1;

select id,title,description from user where id=1; 
```


# GO操作MySQL

```
database/sql
```

## 链接
```
func Open(drverName,dataSourcname string) (*DB,error)
```

链接数据库
```
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
```

初始化链接
```
func queryOne(id int){}
```

## 查询

写单条查询记录的sql语句
```
var u1 user

sqlStr := ` select id, title,description from user where id=1;`
```

查看某条数据
```
queryOne(?)
```

执行 拿取结果
```
db.QueryRow(sqlStr).Scan(&u1, &u1.title, &u1,description)
```

设置数据库最大的连接数
```
db.SetMaxOpenConns(10)
```

多行查询
```
sqlStr := ` select id, title,description from user where id > ?;`
```

关闭rows
```
defer rows.Close()
```

取结果
```
for rows.Next() {}
```

插入数据
```
sqlStr := ` insert ino user(id, title) values(?,?)`

stmt, err := db.Prepare(sqlStr)

db.Exec(sqlStr)
```

更新数据
```
sqlStr := "movies user set id=? where title= ?"
```

删除数据
```
sqlStr := `delete from suer where id=?`
```

## MySQL预处理

预处理查询
```
defer stmt.Close()
	rows, err := stmt.Query(0)
```

预处理插入
```
defer stmt.Close()
	_, err = stmt.Exec("?", ?)
```

## GO实现MySQL事务
开始事务
```
func (db *DB) Begin() (*Tx, error)
```

提交事务
```
func (tx *Tx) Commit() error
```

回滚事务
```
func (tx *Tx) Rollback() error
```

## sqlx使用
安装
```
go get github.com/jmoiron/sqlx

	go get"github.com/go-sql-driver/mysql" //init()
```

sqlx查询单条数据
```
err := db.Get(&?，?,?)
```

sqlx多条查询
```
err := db.Select(&?,?)
```

# GO操作REdis
准备Redis环境
```
docker run --name redis507 -p 6379:6379 -d redis:5.0.7
```

链接Redis环境
```
docker run -it --network host --rm redis:5.0.7 redis-cli
```

安装
```
go get github.com/redis/go-redis/v8


go get github.com/redis/go-redis/v9
```

连接
```
import "github.com/redis/go-redis/v9"
```

普通连接模式
```
rdb := redis.NewClient(){}

除此之外，还可以使用 redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}
rdb := redis.NewClient(opt)
```

TLS连接模式
```
rdb := redis.NewClient(&redis.Options{
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
		// Certificates: []tls.Certificate{cert},
    // ServerName: "your.domain.com",
	},
})
```

Redis Sentinel模式
```
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},
})
```

Redis Cluster模式
```
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

    // 若要根据延迟或随机路由命令，请启用以下命令之一
    // RouteByLatency: true,
    // RouteRandomly: true,
})
```

## 基本使用
执行命令
```
set/Get
```

执行任意命令
```
Do
```

判断redis.Nil
```
// getValueFromRedis redis.Nil判断
func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		// 出其他错了
		return "", err
	}
	return val, nil
}
```

zset示例
```
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
v8版本使用[]*redis.Z；此处为v9版本使用[]redis.Z
languages := []redis.Z{}
```

扫描或遍历所有key
```
vals, err := rdb.Keys(ctx, "user:*").Result()


iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
```

Watch方法
```
Watch(fn func(*Tx) error, keys ...string) error
```

# GO语言操作NSQ

启动nsqd，指定-broadcast-address=127.0.0.1来配置广播地址
```
./nsqd -broadcast-address=127.0.0.1
```

如果是在搭配nsqlookupd使用的模式下需要还指定nsqlookupd地址:
```
./nsqd -broadcast-address=127.0.0.1 -lookupd-tcp-address=127.0.0.1:4160
```


安装
```
go get -u github.com/nsqio/go-nsq
```

生产者
``` 
// nsq_producer/main.go
package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
 
    "github.com/nsqio/go-nsq"
)
 
// NSQ Producer Demo
 
var producer *nsq.Producer
 
// 初始化生产者
func initProducer(str string) (err error) {
    config := nsq.NewConfig()
    producer, err = nsq.NewProducer(str, config)
    if err != nil {
        fmt.Printf("create producer failed, err:%v\n", err)
        return err
    }
    return nil
}
 
func main() {
    nsqAddress := "127.0.0.1:4150"
    err := initProducer(nsqAddress)
    if err != nil {
        fmt.Printf("init producer failed, err:%v\n", err)
        return
    }
 
    reader := bufio.NewReader(os.Stdin) // 从标准输入读取
    for {
        data, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("read string from stdin failed, err:%v\n", err)
            continue
        }
        data = strings.TrimSpace(data)
        if strings.ToUpper(data) == "Q" { // 输入Q退出
            break
        }
        // 向 'topic_demo' publish 数据
        err = producer.Publish("topic_demo", []byte(data))
        if err != nil {
            fmt.Printf("publish msg to nsq failed, err:%v\n", err)
            continue
        }
    }
}
```

消费者
```
// nsq_consumer/main.go
package main
 
import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
 
    "github.com/nsqio/go-nsq"
)
 
// NSQ Consumer Demo
 
// MyHandler 是一个消费者类型
type MyHandler struct {
    Title string
}
 
// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
    fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
    return
}
 
// 初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
    config := nsq.NewConfig()
    config.LookupdPollInterval = 15 * time.Second
    c, err := nsq.NewConsumer(topic, channel, config)
    if err != nil {
        fmt.Printf("create consumer failed, err:%v\n", err)
        return
    }
    consumer := &MyHandler{
        Title: "沙河1号",
    }
    c.AddHandler(consumer)
 
    // if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
    if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
        return err
    }
    return nil
 
}
 
func main() {
    err := initConsumer("topic_demo", "first", "127.0.0.1:4161")
    if err != nil {
        fmt.Printf("init consumer failed, err:%v\n", err)
        return
    }
    c := make(chan os.Signal)        // 定义一个信号的通道
    signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
    <-c                              // 阻塞
}
```