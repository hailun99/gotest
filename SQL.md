选择库
```
use RUNOOB;
```

设置字符集
```
set naems utf8;
```

读取表格信息
```
select * from Websites;
```

# SQL命令
```
SELECT - 从数据库中提取数据
UPDATE - 更新数据库中的数据
DELETE - 从数据库中删除数据
INSERT INTO - 向数据库中插入新数据
CREATE DATABASE - 创建新数据库
ALTER DATABASE - 修改数据库
CREATE TABLE - 创建新表
ALTER TABLE - 变更（改变）数据库表
DROP TABLE - 删除表
CREATE INDEX - 创建索引（搜索键）
DROP INDEX - 删除索引
```

# 查询数据表中的数据[升序或降序]
```
SELECT column_name FROM table_name WHERE condition ORDER BY column_name [ASC | DESC]
```

多行排序
```
SELECT column_name(s), aggregate_function(column_name)
FROM table_name
WHERE condition
GROUP BY column_name(s)
```

插入数据
```
INSERT INTO table_name (column1, column2, ...)
VALUES (value1, value2, ...)
```

更新数据
```
UPDATE table_name
SET column1 = value1, column2 = value2, ...
WHERE condition
```

删除数据
```
DELETE FROM table_name WHERE condition
```

# 创建表
```
CREATE TABLE table_name (
    column1 data_type constraint,
    column2 data_type constraint,
    ...
)
```

修改表数据
```
ALTER TABLE table_name ADD column_name data_tybe

ALTER TABLE table_name DROP COLUMN column_name
```

删除数据库
```
DROP TABLE table_name
```

创建索引
```
CREATE INDEX index_name
ON table_name (column_name)
```

删除索引
```
DROP INDEX index_name NO table_name
```

多表结合
```
SELECT column_name(s)
FROM table_name1
JOIN table_name2
ON table_name1.column_name = table_name2.column_name
```

返回不同的值
```
SELECT DISTINCT column_name(s)
FROM table_name
```

# Redis
Windis下载地址
```
https://github.com/tporadowski/redis/releases
```

Linux下载地址
```
http://redis.io/download
```

redis启动
```
redis-server.exe redis.windows.conf
```

建立新窗口
```
redis-cli.exe -h 192.168.3.32 -p 6379
```

## Redis命令
链接本地Redis服务
```
redis-cli
```

检测Redis服务是否启动
```
PING
```

远程服务器链接
```
redis-cli -h host -p port -a password
```

停止redis
```
redis-cli shutdown

kill redis-pid
```

## 切换数据库
数据库没有名称，默认为16个，0-15
选择编号为1的库数据库
```
select 1
```

## key
*任意字符,？表示单个字符
查看全部
```
keys *
```

### 设置(修改)key
```
set runoobkey redis
```

删除key
```
del runoobk
```

设置键值过期时间(默认单位为秒)
```
setex user2 3 aaa
```

设置多个键值对
```
mset user2 aaa user3 bbb user4 ccc user5 ddd
```

追加值
```
append user2 hahha
```

### 获取值
获取值
```
get user2
```

获取多个值
```
mget user2 user3
```

删除键
```
del user1
```

删除多个键
```
del user2 user3
```

### 键命名
查找键
```
keys *
```

判断键是否存在(存在返回1,不存在返回0)
```
exists user4

exists user4 user5
```

查看键的类型
```
type user4
```

设置过期时间(大于0，代表有效时间,单位为秒;
            返回-1为永久有效;
            返回-2为键不存在)
```
expire user4 100
```

## hash(表， 字段，  值)
        key  field   value
添加/修改
```
hset huser1 name tom
```

设置多个值
```
hmset huser2 name marry sex female
```

type
```
type huser2
```

获取某个字段值
```
hget huser2 name
```

获取多个字段值
```
hmget huser2 name sex
```

获取所有值
```
hvals huser2
```

获取所有字段包括值
```
hgetall huser2
```



