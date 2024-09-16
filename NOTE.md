# GO语言
## Array
定义一个数组
```
var a []int/string
```

数组初始化
```
var arr1 = [5]int {0,1,2,3,4} 
arr2 := [5]int {0,1,2,3,4}
```

多维数组
```
var arr1 [][]int
```

遍历
```
for _, value = rang arr1{
  fmt.print(value)
}
```

## Slice
```
var a []string
a := []string
a := make([]string, 10)
```

切片初始化
```
var slice0 []int = arr[start:end] 
var slice1 []int = arr[:end]        
var slice2 []int = arr[start:]        
var slice3 []int = arr[:] 
```

切片追加
```
var arr = []{1,2,3}
c := append(a,4,5,6)
```

copy
```
copy(arr)
```

遍历
```
 for index, value := range slice 
```

## Map
```
map[KeyType]ValueType
```

初始化Map
```
make(map[KeyType]ValueType, [cap])
```

遍历Map
```
for k, v := range scoreMap
```

查找
```
value, ok := map[key]
```

delete
```
delete(map, key)
```

## 函数
定义一个func
```
func arr {
  arr(1,2)
}
```

## 结构体
定义给个结构体
```
type a struct {
  name string
  id int
} 
```

## 并发
创建一个channel
```
ch := make(chan int)
```

# GO操作MySQL

 GO语言操作sql
```
database/sql
```

链接
```
func Open(drverName,dataSourcname string) (*DB,error)
```

链接数据库
```
db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
```

## 增删改查
### 查
```
var u suer // 声明这个库
sqlStr := "selest id,name,age from suer where id=?"
row := db.QueryRow(sqlStr, id)
err := row.Scan(&u.id, &u.name, &u.age)
```

多行查询
```
sqlStr := "selest id,name,age from suer where id > ?"
rows, err := db.Query(sqlStr, id)

defer rows.Close()

for rows.Next(){
    var u user
    err := rows.Scan(&u.id, &u.name, &u.age)
    }
```

### 增
```
sqlStr := "insert into user(name, age) values (?,?)"
ret, err := db.Exec(sqlstr, "?"， ?)

theID, err := ret.LastInsertId() // 新插入数据的id
```

增加多行
```
sqlStr := "insert into user(name, age) values (?,?)"
insertInfo := map[string]int{
    
}
for k, v := range insertInfo {
		ret, err := db.Exec(sqlStr, k, v
}
    theID, err := ret.LastInsertId() // 新插入数据的id
```


### 更新数据
```
sqlStr := "update user set age = ?, id = ?"
ret, err := db.Exes(sqlStr, age , id)

rowNum, err := ret.RowsAffected()	// 操作影响的行数
```

### 删除数据
```
sqlStr := "delete from user where id = ?"
ret, err := db.Exes(sqlStr,id)

n, err := ret.RowsAffected() // 操作影响的行数
```

## 事务处理
开启事务
```
Begin()
```

提交事务
```
Commit
```
回滚事务
```
Rplldack
```

##
```
drop
```

# GO操作Redis

## String类型
```
err := client.Set(ctx, k).Result // 设置键

value, err := client.Get(key).Result() // 获取key

err := client.MGet(ctx, k...) // 批量查询key的值

client.MSet(ctx, values) // 批量设置key的值

client.Del(ctx, k/...) // 删除单个或多个key

client.Expire(ctx, k, t) // 设置key的过期时间

val, err := client.Incr(ctx, "key").Result() // Incr函数每次加一
valBy, err := client.IncrBy(ctx, "key", 2).Result() // IncrBy函数，可以指定每次递增多少
valFloat, err := client.IncrByFloat(ctx, "key1", 2.2).Result() // IncrByFloat函数，可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数

val, err := client.Decr(ctx, "key").Result() // Decr函数每次减一
valBy, err := client.DecrBy(ctx, "key", 2).Result() // DecrBy函数，可以指定每次递减多少
```

## hash
```
client.HSet(hashKey, "name", "叶子")  // 存储
client.HSet(hashKey, "age", 18)

hashGet, _ := client.HGet(hashKey, "name").Result() // 获取某个属性

hashGetAll, _ := client.HGetAll(hashKey).Result() // 获取它的全部属性
```

## 集合
```
client.SAdd(setKey, "set1") // 往集合里添加数据

setList, _ := client.SMembers(setKey).Result() // 获取集合的所有的元素

client.SRem(setKey, "set1") // 删除集合里的set1

setFirst, _ := client.SPop(setKey).Result() // 移除并返回set的一个随机元素,因为set是无序的
```

## zset 有序集合
```
zsetKey := "go2zset"  
ranking := []*redis.Z{  
  &redis.Z{Score: 100.0, Member: "钟南山"},  
  &redis.Z{Score: 80.0, Member: "林医生"},  
  &redis.Z{Score: 70.0, Member: "王医生"},  
  &redis.Z{Score: 75.0, Member: "张医生"},  
  &redis.Z{Score: 59.0, Member: "叶医生"},  
}  
client.ZAdd(zsetKey, ranking...)  // 往zset里加入集合数据，数据是[]*redis.Z类型，里面包含Score和Member2个属性
 
newScore, err := client.ZIncrBy(zsetKey, 5.0, "钟南山").Result()  // 给钟南山加上5分，返回钟南山的最新热度分值

fmt.Println("钟南山加5分后的最新分数", newScore)  
  
zsetList2, _ := client.ZRevRangeWithScores(zsetKey, 0, 1).Result()  // 获取前2名热度的医生，前2名，所以索引是从0到1

fmt.Println("zset前2名热度的医生", zsetList2)
```


# NSQ
nsq_stat
```
用于获取指定topic，channel下系统信息
```

nsq_tail
```
用于将指定topic，channel中的信息打印到控制台
```

nsq_to_file
```
用于将指定topic，channel中的信息导出到文件
```

nsq_to_http
```
用于将指定topic，channel中的信息发送到http服务器
```

nsq_to_nsq
```
用于将指定topic，channel中的信息发送到其他http服务器
```

to_nsq
```
用于将控制台信息发送到指定的topic
```

创建唯一数据库
```
create table users(
  id int auto_increment primary key,
  username varchar(50) not null unique,
  password varchar(100) not null,
  created int
);


CREATE TABLE movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    created INT
);

create table comments(
  id int auto_increment primary key,
  username varchar(50) not null,
  movieid int not null,
  score int not null,
  comment varchar(200),
  created int
);
```


```
create table tickek(
  id int auto_increment primary key,
  ciname varchar(20) not null unique,
  movie varchar(10) not null,
  type varchar(10) not null,
  seat varchar(30) not null,
  price varchar(10) not null,
  created int
);
```


```
performer
sql增加字段
```
alter table movies add column directo varchar(50) not null;
```

删除字段
```
alter table users drop column nicname;
```

## 查看表的结构
```
DESC 表名;
DESCRIBE 表名;
SHOW COLUMNS FROM 表名;
EXPLAIN 表名;
```

删除字段
```
alter table comments drop directo;
```
Directo  string `json:"directo"`

# go操作redis
创造一个redis客户端
```
func RedisInit(addr string) (*redis.Client, error) {
	reidsdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		return nil, err
	} else {
		return reidsdb, nil
	}
}
```

## 字符串类型
添加
```
set key value

mset key value [key value...]增加多条
```

Del(key)
```
_, err = rdb.Del(ctx, "name").Result()

flushdb 删除当前数据库所有kye

fiuhall 删除所有

改
```
set key value
```

查ste(key, value, expiration)
```
err := redisdb.Set("key1","hello", 5*time.Second).Err()

get key

keys * 查看所有kye
```

get(key)
```
val, err := rdb.Get(ctx, "name").Result()
```



```

## 集合
添
```
err := rdb.SAdd(ctx, "tags", "redis", "go", "database").Err())
```

删
```
srem 集合名称 元素1 元素2.....
srem user tom jerry
```

查
```
tags, err := rdb.SMembers(ctx, "tags").Result()
```

有序集合
ZAdd
```
members := []redis.Z{
{Score: 1.0, Member: "member1"},
{Score: 2.0, Member: "member2"},
{Score: 3.0, Member: "member3"},

result := rdb.ZAdd(ctx, "myzset", members...)
```

## Hash
添加数据/改
```
hest key field value  添加一条

hmest key field value[field value...]
```

删除数据
```
hael key field
```

查
```
hget key field  获取value

hmget key field[fiels...] 获取多个value

hvals key 获取全部value

hkeys key 获取全部field

hgetall key 获取全部field和value

hlen key 查看有几个键值对
```


# Nsq