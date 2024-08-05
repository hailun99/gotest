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

# GO命令
```
go run main.go
```

# Go 语法


## 数组 Array

### 定义一个数组
```
var a [5]int
```

### 初始化数组
```
var a = [5]int{1,2,3,4,5}
a := [5]int{1,2,3,4,5}

var = [...]int{1,2,3,4,5,6}
a := [...]int{1,2,3,4,5,6}

a := [5]int{1:1,3:3}
```

### 访问数组
```
var a [3]int
fmt.Println(a[0])

a := [3]int{}
println(len(a), cap(a))

```

### 数组遍历
```
for i := 0; i < len(a); i ++ {
	mft.Print(a[i], "\t")
} 
fmt.Println()


for _, value := range a {
	fmt.print(value)
} 
```

### 多维数组
```
var a [][]int
```

### 初始化二维数组
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

### 定义一个切片
```
var a []string
a := []string
a := make([]string, 10)
```

### 切片初始化
```
a := arr[:5]

a := arr[1:]

a := arr[O:5]

s := make([]int, len, cap)
```

### appand与copy
```
var a = []int{1,2,3}
a := appand(a, 4,5,6)

a := []int{1,2,3,4,5}
b := []int{6,7,8}
copy(a,b)
```

## Map (集合)

### 定义一个Map
```
a := make(map[string]int)

//指定长度
a := make(map[string]int, 8)

a := map[string]int{
	"name": 1,
	"age": 2,
}
```

### 获取
```
a1 := a["apple"]
a2, ok := a["pear"]

len := len(a)
```

### 修改
```
a["apple"] = 5
```

### 遍历Map
```
for k, v := range a {
	fmt.Printf("key=%s, value=%d\n, k, v")
}
```

### 删除
```
delete(m, "apple")
```

## 函数(function)

### 定义一个函数
```
func a([parameter list]) [return_types]{
	函数体
}

func a() {
	a(1,2,3,4,5)
}
```

### 不定长变参特性
```
正确
func a(a string, b ...int)

错误
func a(b ...int, a string)
```

### 匿名函数
```
func main{a := func(){
	fmt.Println("匿名函数")
}
	a() //调用
}
```

### 闭包
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

### defer(推迟)
```
func main() {
  fmt.Println("a")
  
  defer fmt.Println("b")
  defer fmt.Println("c")
}
```

### panic
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

### 定义一个结构体
```
type Aa struct {
	name string
	age int
}
var a1 Aa
a1.name = "xiaming"
a1.age = 18
```

### 匿名结构体
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

### 匿名字段
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

### 指针
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

### new关键字
```
p2 := new(a)
```

### 结构体嵌套
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

### 嵌入
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

### 创建方法
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

### 创建映射
```
map[KeyType]ValueType
```

### 遍历
```
myColors := map[string]string{
}

for key, value := range myColors {
    fmt.Printf("Key: %s  Value: %s\n", key, value)
}
```

### 查找
```
for k,v := range capital{
    fmt.Println(k,v)
}

for k := range capital{
}

for _, v := range capital{  //无需将值改为匿名变量形式，忽略值即可。
}
```

### delete(删除)
```
delete(map,键)
```

## 并发

### goroutine
```
go 函数名( 参数列表 )
```

### 通道（channel）
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

### 存储引擎
```
MySQL支持插件式的存储引擎.
常见的存储引擎:MylSAM和lnnoDB。
```

### MyLSAM:
```
1.查询速度快
2.只支持表锁
3.不支持事务
```

### lnnoDB
```
1.整体速度快
2.支持表锁和行锁
3.支持事务
```

### 事务
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

### 切换到数据库
```
use gotest;
```

### 查看数据库下的表
```
show tables;
```

### 创建一个表
```
CREATE TABLE movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    created INT
);
```

### 插入数据
```
insert into user(id,title) values(1,哈哈哈哈)
```

### 查询表的所有数据

```
select * from movies;
```

### 单行查询
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

### 链接数据库
```
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
```

### 初始化链接
```
func queryOne(id int){}
```

## 查询

### 写单条查询记录的sql语句
```
var u1 user

sqlStr := ` select id, title,description from user where id=1;`
```

### 查看某条数据
```
queryOne(?)
```

### 执行 拿取结果
```
db.QueryRow(sqlStr).Scan(&u1, &u1.title, &u1,description)
```

### 设置数据库最大的连接数
```
db.SetMaxOpenConns(10)
```

### 多行查询
```
sqlStr := ` select id, title,description from user where id > ?;`
```

### 关闭rows
```
defer rows.Close()
```

### 取结果
```
for rows.Next() {}
```

### 插入数据
```
sqlStr := ` insert ino user(id, title) values(?,?)`

stmt, err := db.Prepare(sqlStr)

db.Exec(sqlStr)
```

### 更新数据
```
sqlStr := "movies user set id=? where title= ?"
```

### 删除数据
```
sqlStr := `delete from suer where id=?`
```

## MySQL预处理

### 预处理查询
```
defer stmt.Close()
	rows, err := stmt.Query(0)
```

### 预处理插入
```
defer stmt.Close()
	_, err = stmt.Exec("?", ?)
```

## GO实现MySQL事务
### 开始事务
```
func (db *DB) Begin() (*Tx, error)
```

### 提交事务
```
func (tx *Tx) Commit() error
```

### 回滚事务
```
func (tx *Tx) Rollback() error
```

## sqlx使用
### 安装
```
go get github.com/jmoiron/sqlx

	go get"github.com/go-sql-driver/mysql" //init()
```

### sqlx查询单条数据
```
err := db.Get(&?，?,?)
```

### sqlx多条查询
```
err := db.Select(&?,?)
```

# GO操作Redis
### 准备Redis环境
```
docker run --name redis507 -p 6379:6379 -d redis:5.0.7
```

### 链接Redis环境
```
docker run -it --network host --rm redis:5.0.7 redis-cli
```

### 安装
```
go get github.com/redis/go-redis/v8


go get github.com/redis/go-redis/v9
```

### 连接
```
import "github.com/redis/go-redis/v9"
```

### 普通连接模式
```
rdb := redis.NewClient(){}

除此之外，还可以使用 redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}
rdb := redis.NewClient(opt)
```

### TLS连接模式
```
rdb := redis.NewClient(&redis.Options{
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
		// Certificates: []tls.Certificate{cert},
    // ServerName: "your.domain.com",
	},
})
```

### Redis Sentinel模式
```
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},
})
```

### Redis Cluster模式
```
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

    // 若要根据延迟或随机路由命令，请启用以下命令之一
    // RouteByLatency: true,
    // RouteRandomly: true,
})
```

## 常用命令
### 设置/获取/删除
```
set/Get/Del

err := rdb.Set(ctx, "name", "Go-Redis", 0).Err()

val, err := rdb.Get(ctx, "name").Result()

_, err = rdb.Del(ctx, "name").Result()
```

### 哈希操作
```
HSet 和 HGetAll

err := rdb.HSet(ctx, "user:1000", "name", "John", "age", 30).Err()

name, err := rdb.HGet(ctx, "user:1000", "name").Result()
```

### 列表操作
```
LPush 和 LPop

err := rdb.LPush(ctx, "tasks", "task1", "task2").Err()

task, err := rdb.LPop(ctx, "tasks").Result()
```

### 集合操作
```
SAdd 和 SMembers

err := rdb.SAdd(ctx, "tags", "redis", "go", "database").Err()

tags, err := rdb.SMembers(ctx, "tags").Result()
```

### 有序集合操作
```
ZAdd 和 ZRange

err := rdb.ZAdd(ctx, "scores", &redis.Z{Score: 100, Member: "Alice"}, &redis.Z{Score: 200, Member: "Bob"}).Err()

scores, err := rdb.ZRangeWithScores(ctx, "scores", 0, -1).Result()
```

# Redis扩展
## 安装Redis客户端
```
go get github.com/go-redis/redis/v8
```

## 基本指令
Keys():根据正则获取keys
```
keys, err := rdb.Keys(ctx, "*").Result()
```

获取key对应值得类型
```
vType, err := rdb.Type(ctx, "key").Result()
```

Del():删除缓存项
```
n, err := rdb.Del(ctx, "key1", "key2").Result()
```

Exists():检测缓存项是否存在
```
n, err := rdb.Exists(ctx, "key1").Result()
```

Expire(),ExpireAt():设置有效期
```
res, err := rdb.Expire(ctx, "key", time.Minute * 2).Result()

res, err = rdb.ExpireAt(ctx, "key2", time.Now()).Result()
```

TTL(),PTTL():获取有效期
```
ttl, err := rdb.TTL(ctx, "key").Result()

pttl, err := rdb.PTTL(ctx, "key").Result()
```

DBSize():查看当前数据库key的数量
```
num, err := rdb.DBSize(ctx).Result()
```

FlushDB():清空当前数据
```
res, err := rdb.FlushDB(ctx).Result()
```

FlushAll():清空所有数据库
```
res, err := rdb.FlushAll(ctx).Result()
```

## 字符串(string)类型
Set():设置
```
err = rdb.Set(ctx, "key1", "value1", 0).Err()

err = rdb.Set(ctx, "key2", "value2", time.Minute * 2).Err()
```

SetEX():设置并指定过期时间
```
err := rdb.SetEX(ctx, "key", "value", time.Hour * 2).Err()
```

SetNX():设置并指定过期时间(key不存在是才能设置)
```
res, err := rdb.SetNX(ctx, "key", "value", time.Minute).Result()
```

Get():获取
```
val, err := rdb.Get(ctx, "key").Result()
```

GetRange():字符串截取
```
val, err := rdb.GetRange(ctx, "key", 1, ?).Result()
```

Incr():增加+1
```
val, err := rdb.Incr(ctx, "number").Result()
```

IncrBy():按指定步长增加
```
val, err := rdb.IncrBy(ctx, "number", 12).Result()
```

Decr():减少-1
```
val, err := rdb.Decr(ctx, "number").Result()
```

DecrBy():按指定步长减少
```
val, err := rdb.DecrBy(ctx, "number", 12).Result()
```

Append():追加
```
err := rdb.Set(ctx, "key", "hello", 0).Err()

length, err := rdb.Append(ctx, "key", " world!").Result()  //hello world
```

StrLen():获取长度
```
err := rdb.Set(ctx, "key", "hello world!", 0).Err()

length, err := rdb.StrLen(ctx, "key").Result() //12
```

## 列表(list)类型
LPush()/RPush():将元素从左/右压入链表
```
n, err := rdb.LPush(ctx, "list", 1, 2, 3).Result()
```

LInsert():在某个位置插入新元素
```
err := rdb.LInsert(ctx, "key", "before", "100", 123).Err()
```

LSet():设置某个元素的值
```
err := rdb.LSet(ctx, "list", 1, 100).Err()
```

LLen():获取链表元素个数
```
length, err := rdb.LLen(ctx, "list").Result()
```

LIndex():获取链表下标对应的元素
```
val, err := rdb.LIndex(ctx, "list", 0).Result()
```

LRange():获取某个选定范围的元素集
```
als, err := rdb.LRange(ctx, "list", 0, 2).Result()
```

LPop()/RPop() 从链表左/右侧弹出数据
```
val, err := rdb.LPop(ctx, "list").Result()
```

LRem():根据值移除元素
```
n, err := rdb.LRem(ctx, "list", 2, "100").Result()
```

## 集合(set)类型
SAdd():添加元素
```
rdb.SAdd(ctx, "team", "kobe", "jordan")
	rdb.SAdd(ctx, "team", "curry")
	rdb.SAdd(ctx, "team", "kobe") //由于kobe已经被添加到team集合中，所以重复添加时无效的
```

SPop()/SPopN:随机获取一个/多个元素
```
val, err := rdb.SPop(ctx, "team").Result()
```

SRem():删除集合里指定的值
```
n, err := rdb.SRem(ctx, "team", "kobe", "v2").Result()
```

SSMembers():获取所有成员
```
vals, err := rdb.SMembers(ctx, "team").Result()
```

SIsMember():判断元素是否在集合中
```
exists, err := rdb.SIsMember(ctx, "team", "jordan").Result()
```

SCard():获取集合元素个数
```
total, err := rdb.SCard(ctx, "team").Result()
```

SUnion():并集,SDiff():差集,SInter():交集
```
rdb.SAdd(ctx, "setA", "a", "b", "c", "d"
rdb.SAdd(ctx, "setB", "a", "d", "e", "f")

//并集
union, err := rdb.SUnion(ctx, "setA", "setB").Result()

//差集
diff, err := rdb.SDiff(ctx, "setA", "setB").Result()

//交集
inter, err := rdb.SInter(ctx, "setA", "setB").Result()
```

## 有序集合(zset)类型
ZAdd():添加元素(添加6个元素1~6,分值都是0)
```
rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 1
```

ZIncrBy():增加元素分值(分值可以为负数，表示递减)
```
rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "1")
```

ZRange()、ZRevRange():获取根据score排序后的数据段(根据分值排序后的，升序和降序的列表获取)
```
res, err := rdb.ZRevRange(ctx, "zSet", 0, 2).Result()
```

ZRangeByScore()、ZRevRangeByScore():获取score过滤后排序的数据段(根据分值过滤之后的列表,需要提供分值区间)
```
res, err := rdb.ZRangeByScore(ctx, "zSet", &redis.ZRangeBy{
		Min:    "40",
		Max:    "85",
	}).Result()
```

ZCard():获取元素个数
```
count, err := rdb.ZCard(ctx, "zSet").Result()
```

ZCount():获取区间内元素个数
```
n, err := rdb.ZCount(ctx, "zSet", "40", "85").Result()
```

ZScore():获取元素的score(获取元素分值)
```
score, err := rdb.ZScore(ctx, "zSet", "5").Result()
```

ZRank()、ZRevRank():获取某个元素在集合中的升序/降序排名
```
res, err := rdb.ZRevRank(ctx, "zSet", "2").Result()
```

ZRem():删除元素
```
res, err := rdb.ZRem(ctx, "zSet", "2").Result()
```

ZRemRangeByRank():根据排名来删除
```
//按照升序排序删除第一个和第二个元素
	res, err := rdb.ZRemRangeByRank(ctx, "zSet",  0, 1).Result()
```

ZRemRangeByScore():根据分值区间来删除
```
res, err := rdb.ZRemRangeByScore(ctx, "zSet", "40", "70").Result()
```

## 哈希(hash)类型
HSet():设置
```
rdb.HSet(ctx, "user", "key1", "value1", "key2", "value2")
rdb.HSet(ctx, "user", []string{"key3", "value3", "key4", "value4"})
rdb.HSet(ctx, "user", map[string]interface{}{"key5": "value5", "key6": "value6"})
```

HMset():批量设置
```
rdb.Del(ctx, "user")
rdb.HMSet(ctx, "user", map[string]interface{}{"name":"kevin", "age": 27, "address":"北京"})
```

HGet():获取某个元素
```
address, err := rdb.HGet(ctx, "user", "address").Result()
```

HGetAll():获取全部元素
```
user, err := rdb.HGetAll(ctx, "user").Result()
```

HDel():删除某个元素(支持一次删除多个元素)
```
res, err := rdb.HDel(ctx, "user", "name", "age").Result()
```

HExists():判断元素是否存在
```
res, err := rdb.HExists(ctx, "user", "address").Result()
```

HLen():获取长度
```
res, err := rdb.HLen(ctx, "user").Result()
```



### 执行任意命令
```
Do
```

### 判断redis.Nil
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

### zset示例
```
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
v8版本使用[]*redis.Z；此处为v9版本使用[]redis.Z
languages := []redis.Z{}
```

### 扫描或遍历所有key
```
vals, err := rdb.Keys(ctx, "user:*").Result()


iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
```

### Watch方法
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

# Go 开发环境
下载
```
 https://studygolang.com/dl

 （Windows 环境比较建议直接下载对应的包进行安装 https://studygolang.com/dl/golang/go1.14.4.windows-amd64.msi）
```

配置环境变量
```
GAPATH 工作目录设置一个目录用于下载go库

GO111MODULE 是否开启go mod模式设置为no

GODPROXY 开启go远程代理设置为 https://goproxy.io
```

## sql.db类型上用于查询的方法:
```
Context // 上下文
Query // 0行或者多行记录
QueryRow // 单行查询
QueryContext
QueryRowContext
```

Rows的方法:
```
返回的类型是: type Rpws struct{}

func(rs*Rows)Close()error //关闭

func(rs*Rows)ColumnTypes()([]*CloumnType,error) // 返回结果数据里列里的所有信息

func(rs*Rows)Columns()([]string,error) // 返回所有列名

func(rs*Rows)Err()error // 返回错误

func(rs*Rows)Next()bool // 返回遍历的结果 (每次一行)

func(rs*Rows)NextResultSet()bool // 返回多个结果集

func(rs*Rows)Scan(dest...interface{})error //拷贝出来放在变量里
```

更新(执行命令)
```
Exec

ExecContext
```

# SQL
```
创建数据库：`CREATE DATABASE database_name;`
删除数据库：`DROP DATABASE database_name;`
创建表格：`CREATE TABLE table_name (column1 datatype, column2 datatype, column3 datatype, ...);`
删除表格：`DROP TABLE table_name;`
插入数据：`INSERT INTO table_name (column1, column2, column3, ...) VALUES (value1, value2, value3, ...);`
查询数据：`SELECT * FROM table_name WHERE condition;`
```

SELECT语句：用于从一个或多个数据表中选择数据
```
SELECT column1, column2, ... FROM table_name;
```

INSERT语句：用于向数据表中插入新的行。
```
INSERT INTO table_name (column1, column2, column3, ...) VALUES (value1, value2, value3, ...);
```

UPDATE语句：用于更新数据表中的已有行
```
UPDATE table_name SET column1=value1, column2=value2, ... WHERE condition;
```

DELETE语句：用于从数据表中删除行
```
DELETE FROM table_name WHERE condition;
```

 模糊模糊查询 LIKE
 ```
SELECT column1, column2, ...
FROM table_name
WHERE column_name LIKE pattern;
 ```

出重  UNION
```
SELECT column1, column2, ...
FROM table1
UNION
SELECT column1, column2, ...
FROM table2
[ORDER BY column1, column2, ...];
```

ORDER BY(排序) 语句
```
SELECT column1, column2, ...
FROM table_name
ORDER BY column1 [ASC | DESC], column2 [ASC | DESC], ...;
//升序（ASC）或降序（DESC）
```

GROUP BY 语句,分组
```
SELECT column1, aggregate_function(column2)
FROM table_name
GROUP BY column1;
```

连接的使用
```
INNER JOIN（内连接,或等值连接）
LEFT JOIN（左连接）
RIGHT JOIN（右连接）
```
NULL 值处理
```
IS NULL: 当列的值是 NULL,此运算符返回 true。
IS NOT NULL: 当列的值不为 NULL, 运算符返回 true。
<=>: 比较操作符（不同于 = 运算符），当比较的的两个值相等或者都为 NULL 时返回 true
```

## 正则表达式
```
.：匹配任意单个字符。
^：匹配字符串的开始。
$：匹配字符串的结束。
*：匹配零个或多个前面的元素。
+：匹配一个或多个前面的元素。
?：匹配零个或一个前面的元素。
[abc]：匹配字符集中的任意一个字符。
[^abc]：匹配除了字符集中的任意一个字符以外的字符。
[a-z]：匹配范围内的任意一个小写字母。 BINARY 关键字，使得匹配区分大小写
[0-9]：匹配一个数字字符。
\w：匹配一个字母数字字符（包括下划线）。SELECT * FROM products WHERE product_name RLIKE '^[0-9]';
\s：匹配一个空白字符。
```

MySQL 事务
```
1、用 BEGIN, ROLLBACK, COMMIT 来实现
BEGIN 或 START TRANSACTION：开用于开始一个事务。
ROLLBACK 事务回滚，取消之前的更改。
COMMIT：事务确认，提交事务，使更改永久生效。

2、直接用 SET 来改变 MySQL 的自动提交模式:

SET AUTOCOMMIT=0 禁止自动提交
SET AUTOCOMMIT=1 开启自动提交
```

## ALTER 命令
添加列
```
ALTER TABLE employees
ADD COLUMN birth_date DATE;
```

修改列的数据类型
```
ALTER TABLE employees
MODIFY COLUMN salary DECIMAL(10,2);
```

修改列名
```
ALTER TABLE employees
CHANGE COLUMN old_column_name new_column_name VARCHAR(255);
```

删除列
```
ALTER TABLE employees
DROP COLUMN birth_date;
```

添加 PRIMARY KEY(添加一个主键)
```
ALTER TABLE employees
ADD PRIMARY KEY (employee_id);
```

添加 FOREIGN KEY(添加一个外键)
```
ALTER TABLE orders
ADD CONSTRAINT fk_customer
FOREIGN KEY (customer_id)
REFERENCES customers (customer_id);
```

修改表名
```
ALTER TABLE employees
RENAME TO staff;
```

## 索引
创建索引(CREATE INDEX)
```
CREATE INDEX idx_name ON students (name);
```

修改表结构(添加索引) ALTER TABLE
```
ALTER TABLE employees
ADD INDEX idx_age (age);
```

创建表的时候直接指定(CREATE TABLE)
```
CREATE TABLE students (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  age INT,
  INDEX idx_age (age)
);
```

删除索引的语法(DROP INDEX / ALTER TABLE)
```
DROP INDEX index_name ON table_name;

ALTER TABLE table_name
DROP INDEX index_name;
```

### 唯一索引(CREATE UNIQUE INDEX)
```
CREATE UNIQUE INDEX idx_email ON employees (email);
```

修改表结构添加索引
关键字:(UNIQUE)
```
ALTER TABLE employees
ADD CONSTRAINT idx_email UNIQUE (email);
```

创建表的时候直接指定
```
CREATE TABLE employees (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE
);
```

使用ALTER 命令添加和删除索引
```
mysql> ALTER TABLE testalter_tbl ADD INDEX (c);

mysql> ALTER TABLE testalter_tbl DROP INDEX c;
```

使用 ALTER 命令添加和删除主键
```
mysql> ALTER TABLE testalter_tbl MODIFY i INT NOT NULL;
mysql> ALTER TABLE testalter_tbl ADD PRIMARY KEY (i);

mysql> ALTER TABLE testalter_tbl DROP PRIMARY KEY;
```

显示索引信息
```
mysql> SHOW INDEX FROM table_name\G
........
```

## 临时表
创建临时表
```
CREATE TEMPORARY TABLE temp_table_name (
  column1 datatype,
  column2 datatype,
  ...
);
或
CREATE TEMPORARY TABLE temp_table_name AS
SELECT column1, column2, ...
FROM source_table
WHERE condition;
```

插入数据到临时表
```
INSERT INTO temp_table_name (column1, column2, ...)
VALUES (value1, value2, ...);
```

查询临时表
```
SELECT * FROM temp_table_name;
```

修改临时表
```
ALTER TABLE temp_table_name
ADD COLUMN new_column datatype;
```

删除临时表
```
DROP TEMPORARY TABLE IF EXISTS temp_table_name;
```

复制表(CREATE TABLE ... SELECT)
```
mysql> INSERT INTO clone_tbl (runoob_id,
    ->                        runoob_title,
    ->                        runoob_author,
    ->                        submission_date)
    -> SELECT runoob_id,runoob_title,
    ->        runoob_author,submission_date
    -> FROM runoob_tbl;
//Records: 3  Duplicates: 0  Warnings: 0
```

使用 mysqldump 命令
```
mysqldump -u username -p dbname old_table > old_table_dump.sql

mysql -u username -p new_dbname < old_table_dump.sql
```

## MySQL 元数据
查看所有数据库：
```
SHOW DATABASES;
```

选择数据库：
```
USE database_name;
```

查看数据库中的所有表：
```
SHOW TABLES;
```

查看表的结构：
```
DESC table_name;
```

查看表的索引：
```
SHOW INDEX FROM table_name;
```

查看表的创建语句：
```
SHOW CREATE TABLE table_name;
```

查看表的行数：
```
SELECT COUNT(*) FROM table_name;
```

查看列的信息：
```
SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_KEY
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_SCHEMA = 'your_database_name'
AND TABLE_NAME = 'your_table_name';
```

查看外键信息：
```
SELECT
    TABLE_NAME,
    COLUMN_NAME,
    CONSTRAINT_NAME,
    REFERENCED_TABLE_NAME,
    REFERENCED_COLUMN_NAME
FROM
    INFORMATION_SCHEMA.KEY_COLUMN_USAGE
WHERE
    TABLE_SCHEMA = 'your_database_name'
    AND TABLE_NAME = 'your_table_name'
    AND REFERENCED_TABLE_NAME IS NOT NULL;
```

## 系统数数据库(information_schema 数据库)
SCHEMATA 表
存储有关数据库的信息，如数据库名、字符集、排序规则等
```
SELECT * FROM information_schema.SCHEMATA;
```

TABLES 表 包含有关数据库中所有表的信息，如表名、数据库名、引擎、行数等
```
SELECT * FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'your_database_name';
```

COLUMNS 表
包含有关表中列的信息，如列名、数据类型、是否允许 NULL 等
```
SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = 'your_database_name' AND TABLE_NAME = 'your_table_name';
```

STATISTICS 表
提供有关表索引的统计信息，如索引名、列名、唯一性等。
```
SELECT * FROM information_schema.STATISTICS WHERE TABLE_SCHEMA = 'your_database_name' AND TABLE_NAME = 'your_table_name';
```

KEY_COLUMN_USAGE 表
包含有关表中外键的信息，如外键名、列名、关联表等。
```
SELECT * FROM information_schema.KEY_COLUMN_USAGE WHERE TABLE_SCHEMA = 'your_database_name' AND TABLE_NAME = 'your_table_name';
```

REFERENTIAL_CONSTRAINTS 表
存储有关外键约束的信息，如约束名、关联表等。
```
SELECT * FROM information_schema.REFERENTIAL_CONSTRAINTS WHERE CONSTRAINT_SCHEMA = 'your_database_name' AND TABLE_NAME = 'your_table_name';
```

## 序列使用 （AUTO_INCREMENT）自增主键
使用 AUTO_INCREMENT 创建表的例子：
```
CREATE TABLE example_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50)
);
```

使用 LAST_INSERT_ID() 函数来获取刚刚插入的行的自增值
```
SELECT LAST_INSERT_ID();
```

获取表的当前自增值
```
SHOW TABLE STATUS LIKE 'example_table';
```

## 处理重复数据
防止表中出现重复数据
```
双主键模式
CREATE TABLE person_tbl
(
   first_name CHAR(20) NOT NULL,
   last_name CHAR(20) NOT NULL,
   sex CHAR(10),
   PRIMARY KEY (last_name, first_name)
);

INSERT IGNORE INTO 当插入数据时，在设置了记录的唯一性后，如果插入重复数据，将不返回错误，只以警告形式返回。 而 REPLACE INTO 如果存在 primary 或 unique 相同的记录，则先删除掉。再插入新记录。
mysql> INSERT IGNORE INTO person_tbl (last_name, first_name)
    -> VALUES( 'Jay', 'Thomas');
Query OK, 1 row affected (0.00 sec)
mysql> INSERT IGNORE INTO person_tbl (last_name, first_name)
    -> VALUES( 'Jay', 'Thomas');
Query OK, 0 rows affected (0.00 sec)
```

UNIQUE（唯一）
```
CREATE TABLE person_tbl
(
   first_name CHAR(20) NOT NULL,
   last_name CHAR(20) NOT NULL,
   sex CHAR(10),
   UNIQUE (last_name, first_name)
);

```

统计重复数据
```
mysql> SELECT COUNT(*) as repetitions, last_name, first_name
    -> FROM person_tbl
    -> GROUP BY last_name, first_name
    -> HAVING repetitions > 1;
```

过滤重复数据
```
mysql> SELECT DISTINCT last_name, first_name
    -> FROM person_tbl;

mysql> SELECT last_name, first_name
    -> FROM person_tbl
    -> GROUP BY (last_name, first_name);
```

删除重复数据
```
mysql> CREATE TABLE tmp SELECT last_name, first_name, sex FROM person_tbl  GROUP BY (last_name, first_name, sex);
mysql> DROP TABLE person_tbl;
mysql> ALTER TABLE tmp RENAME TO person_tbl;

 通过INDEX（索引） 和 PRIMAY KEY（主键）来删除
mysql> ALTER IGNORE TABLE person_tbl
    -> ADD PRIMARY KEY (last_name, first_name);
```

## 及 SQL 注入
```
规定范围
addcslashes(string,characters)
```

## 导出数据
```
SELECT...INTO OUTFILE

LOAD DATA INFILE是SELECT ... INTO OUTFILE的逆操作
```

mysqldump 是 MySQL 提供的用于备份和导出数据库的命令行工具
```
$ mysqldump -u root -p --no-create-info \
            --tab=/tmp RUNOOB runoob_tbl
password ******
```

导出整个数据库
```
mysqldump -u root -p mydatabase > mydatabase_backup.sql
```

导出特定表
```
mysqldump -u username -p password -h hostname database_name table_name > output_file.sql
或
mysqldump -u root -p mydatabase mytable > mytable_backup.sql
```

导出数据库结构
```
mysqldump -u username -p password -h hostname --no-data database_name > output_file.sql
```

导出压缩文件
```
mysqldump -u username -p password -h hostname database_name | gzip > output_file.sql.gz
```

导出 SQL 格式的数据
```
$ mysqldump -u root -p RUNOOB runoob_tbl > dump.txt
password ******
```

### 将数据表及数据库拷贝至其他主机
```
如果你需要将数据拷贝至其他的 MySQL 服务器上, 你可以在 mysqldump 命令中指定数据库名及数据表。

在源主机上执行以下命令，将数据备份至 dump.txt 文件中:

$ mysqldump -u root -p database_name table_name > dump.txt
password *****
如果完整备份数据库，则无需使用特定的表名称。

如果你需要将备份的数据库导入到MySQL服务器中，可以使用以下命令，使用以下命令你需要确认数据库已经创建：

$ mysql -u root -p database_name < dump.txt
password *****
你也可以使用以下命令将导出的数据直接导入到远程的服务器上，但请确保两台服务器是相通的，是可以相互访问的：

$ mysqldump -u root -p database_name \
       | mysql -h other-host.com database_name
以上命令中使用了管道来将导出的数据导入到指定的远程主机上
```

## 导入数据
```
# mysql -uroot -p123456 < runoob.sql
```

source 命令导入
```
source 命令导入数据库需要先登录到数库终端：

mysql> create database abc;      # 创建数据库
mysql> use abc;                  # 使用已创建的数据库 
mysql> set names utf8;           # 设置编码
mysql> source /home/abc/abc.sql  # 导入备份数据库
使用 source 命令的好处是，你可以在 MySQL 命令行中直接执行，而无需退出 MySQL 并使用其他命令。
```

使用 LOAD DATA 导入数据
```
mysql> LOAD DATA LOCAL INFILE 'dump.txt' INTO TABLE mytbl;
```

使用 mysqlimport 导入数据
```
从文件 dump.txt 中将数据导入到 mytbl 数据表中, 可以使用以下命令：

$ mysqlimport -u root -p --local mytbl dump.txt
```

