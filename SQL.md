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

监控
```
monitor
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

### 删除键
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

删除字段
```
hdel huser2 sex
```

## list
### 添加
列表左边添加值
```
lpush luser1 aa bb cc
```

列表右边添加值
```
rpush luser1 dd
```

从值前/后插入值
```
linsert luser1 before cc 11

linsert luser1 after cc 22
```

### 获取(0表示开始值，-1表示最后一个值)
```
lrange luser1 0 -1
```

### 修改
```
lset luser1 1 ww
```

## 删除值
从右侧(末尾)删除(一个)
```
lrem luser2 -1 h1
```

从左侧(开头)删除(下面全部)
```
lrem luser2 0 h1
```

## 集合set(不允许有重复值,只能添加与删)
添加值
```
sadd suser1 aa bb cc
```

获取值
```
smembers suser1
```

删除值
```
srem suser1 aa
```

## zset(不允许有重复值,只能添加与删,分数可以为负数)
添加值(通过分数从大倒小排序)
```
zadd zuser1 1 aa 3 dd 6 cc 2 ff
```

获取值
```
zrange zuser1 0 -1
```

获取值和分数
```
zrange zuser1 0 -1 withscores
```

通过分数获取值
```
zrangebyscore zuser1 1 6
```

通过值获得分数
```
zscore zuser1 cc
```

# MySQL

## DDL
### 查看所有库
```
show databases;
```

### 查询当前数据库
```
select database();
```

### 创建数据库
```
create darabase [if not exists] 数据库名 [default character set 字符集] [collate 字符集排序规则];

create database test;
```

### 删除数据库
```
drop database [if exists] 数据库名;

drop database test; --直接删除数据库，不检查是否存在

drop database if exists test; 数据库存在才执行删除操作
```

### 选择数据库
```
use database_name;
```

## SQL
### 插入数据
```
insert into 表名 (字段1,字段2,字段3...) values (值1,值2,值3...); // 给指定字段插入数据
insert  into user(id, name, age) value (1,'张三',18);

insert into 表名 values (值1,值2,值3...); // 给所有字段插入数据
insert into user value (2,'李四',21);

insert into 表名(字段1,字段2...) values (值1,值2...)，(值1,值2...)，(值1,值2...); // 给指定字段插入多条数据
insert into 表名 values (值1,值2...), (值1,值2...), (值1,值2...); // 给所有字段插入多条数据
insert into user value (3,'王五',17), (4,'赵六',34);

alter table  emp add column idcard  char(18) after entrydata;
```

### 更新
```
update 表名 set 字段名 = 值 where 条件; // 更新指定数据
update user set name = '张亮' where id = 1;
```

### 删除
```
delete from 表名 where 条件; // 删除指定数据
delete from user where id = 2;
delete from user; // 清空表，仅删除数据，不删除表结构
```

## DQL
```
select 字段列表 from 表名列表 where 条件列表 group by 分组字段列表 having 分组后条件列表 order by 排序字段列表 limit 分页参数;
```

### 基本查询
```
1.查询多个字段
select 字段1, 字段2, 字段3... from 表名; 
select * from emp;
select name,workno,age from emp;

2.设置别名
select 字段1 [as 别名], 字段2 [as 别名]... from 表名;
select name as '姓名' from emp;

3.去重
select distinct 字段列表 from 表名;
select distinct name '姓名' from emp;
```

### 条件查询
```
1.语法
select 字段列表 from 表名 where 条件列表;

2.条件
> // 大于
>= // 大于等于
< // 小于
<= // 小于等于
= // 等于
<> 或 != // 不等于
between ... and ... // 区间查询
in (值1,值2,值3...) // 在指定的值中
like 占位符 // 模糊查询(' _ '查询单个字符，' % '查询多个字符)
is null // 为空
and 或 && // 并且
or 或 || // 或者
not 或 ! // 非

SELECT * FROM emp WHERE name like '__'; // 查询姓名为两个字符的
SELECT * FROM emp WHERE id like '%5'; // 查询id以5结尾的
```

### 聚合函数
```
count // 统计数量
max // 最大值
min // 最小值
avg // 平均值
sum // 求和

语法
select 函数名(字段名) from 表名;
select count(id) FROM emp;
select sum(age) from emp where gender ='男';
```

### 分组查询
```
where 与 having区别 
执行时间不同:where 是分组之前进行过滤，不满足where条件，不参与分组;而having是分组之后对结果进行过滤
判断条件不同: where 不能对聚合函数进行判断，而having可以
1.语法
select 字段列表 from 表名 [where 条件] group by 分组字段 [having 分组后条件];

select gender, count(*) from emp group by gender; // 统计区分男女员工数量

查询年龄小于30的，根据工作地址进行分组，获取员工大于等于3的工作地址
select workaddress, count(*) address_count from emp where age < 30 group by workaddress having address_count >= 3;
```

### 排序查询
```
1.语法
select 字段列表 from 表名 order by 字段1 排序方式1,字段2 排序方式2; 

排序方式
ASC // 升序(默认)
DESC // 降序
```

### 分页查询
```
1 语法
select 字段列表 from 表名 limit 起始索引,查询记录数;
select * from emp limit 0,10; // 查询第一页10条记录

```

## DCL
### 用户管理
```
1.查询用户
usr mysql;
select * from user;

2.创建用户 (% 表示所有主机)
create user '用户名@'主机名' identified by '密码';
create user 'qin'@'%' identified by '123456';
create user 'qin'@'localhost' identified by '123456';

3.修改用户密码
alter user '用户名@'主机名' identified with mysql_native_password by '新密码';
alter user 'qin'@'%' identified with mysql_native_password by '1234';

4.删除用户
drop user '用户名@'主机名;
drop user 'qin'@'%';
```

## 函数
### 字符串函数
```
concat(字段1,字段2,字段3...) // 字符串拼接
lower(字段) // 字符串转小写
upper(字段) // 字符串转大写
lpad(字段,长度,'填充字符') // 左填充,字符串填充
rpad(字段,长度,'填充字符') // 右填充,字符串填充
trim(字段) // 去除字符串两端的空格
substring(字段,起始位置,长度)
```


查看数据库字符集
```
show create database test1;
```

修改数据库的字符集
```
alter database test1 character set gbk;
```

## 表
查看表
```
show tables;
```

创建表格
```
create table qin(
    -> id int auto_increment primary key,
    -> user varchar(10) not null,
    -> sex varchar(10) not null
    -> );
```

### 删除表
```
drop table users;
```

### 查看表结构
```
DESC 表名;
```

## 在表中添加字段
```
alter table users add gende varchar(10);
alter table users add signature varchar(30);
alter table users add vip varchar(10);
alter table users add created timestamp default current_timestamp;
```

## 删除字段
```
alter table users drop created;
```

## 修改字段名
```
alter table users change gende gender varchar(5);
```

delete from users where id = 1; // 删除指定数据
delete from users; // 清空表，仅删除数据，不删除表结构


insert into users(username,password,gender,'singnature,vip) values('李四','123456','男','                          man','Svip');


