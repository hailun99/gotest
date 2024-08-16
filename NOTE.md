# GO语言操作sql
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

Slice
```
var a []string
a := []string
a := make([]string, 10)
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