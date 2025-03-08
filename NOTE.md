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

复制
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


sql增加字段
```
alter table movies add column directo varchar(50) not null;

alter table tickek add column consumer varchar(50) not null; //添加顾客
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

删除数据库
```
drop database [if exists] 数据库名称;
```

删除字段
```
alter table comments drop directo;
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















