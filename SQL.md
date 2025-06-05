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



# MySQL

## DDL（Data Definition Language，数据定义语言）
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

## DML（Data Manipulation Language，数据操作语言）
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

## DQL（Data Query Language，数据查询语言）
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

## DCL（Data Control Language，数据控制语言）
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

select concat('Hello','MySql'); // HelloMySql
select lower('Hello'); // hello
select upper('Hello'); // HELLO
select lpad('01',5,'-'); // ---01
select rpad('01',5,'-'); // 01---
select trim(' Hello MySQL '); // Hello MySQL
select substring('Hello MySQL',1,5); // Hello
update emp set workno = rpad(workno, 5, '0') where id = 1; // Hello
```

### 数值函数
```
CEIL(x) // 向上取整
FLOOR(x) // 向下取整
MOD(x,y) // 取余
RAND() // 随机数0~1
ROUND(x,y) // 四舍五入,保留小数

-- 数值函数
-- ceil
select ceil(1.5); // 2

-- floor
select floor(1.9); // 1

-- mod
select mod(5,3); // 2

-- rand
select rand(); // 0~1

-- round
select round(2.2453,2); // 2.24

-- 案例，通过函数生成一个六位数的验证码
select lpad(round( rand()*1000000, 0), 6, '0'); // 数据六个数
```

### 时间函数
```
curdate(); // 当前日期
curtime(); // 当前时间
now();// 当前日期和时间
year(data); // data年份
month(data); // data月份
day(data); // data日期
date_add(data, interval expr day); // 返回一个日期/时间值加上一个时间间隔expr后的时间值
datediff(data1, data2); // 返回两个日期之间的天数差

-- 日期函数
-- curdate()
select curdate();

-- curtime
select curtime();

-- now
select now();

-- year, month, day
select year(now());
select month(now());
select day(now());

-- date_add
select date_add(now(), interval 70 day );
select date_add(now(), interval 70 month );
select date_add(now(), interval 70 year );

-- datediff
select datediff('2025-12-01','2025-10-01')
select datediff('2025-8-01','2025-10-01')

-- 案例,查找员工的入职天数,根据入职天数倒序排序
select name, datediff(curdate(),entrydata) as 'entrydays' from emp order by entrydays desc;
```

### 流程函数
```
if(value, t, f) // 如果value为true,则返回t,否则返回f
ifnull(value1,value2) // 如果value1不为空,返回value1,否则返回value2\
case when [val1] then [res1] ...else [default] end // 如果val1为true,返回res1, ...否则返回default默认值
case [expr] when [val1] then [res1] ...else [default] end // 如果ecpr的值等于val1,返回res1, ...否则返回default默认值

-- 流程图控制函数
-- if
select if(true, 'ok', 'Error');  -- ok
select if(false, 'ok', 'Error'); -- Error

-- ifnull
select ifnull('ok','default'); -- ok
select ifnull('','default'); -- default
select ifnull(null,'default'); -- default

-- case when then else end
select
    name,
    (case wokaddress when '北京' then '一线城市' when '上海' then '一线城市' else '二线城市' end ) as '工作地址'
from emp;

-- 案例,统计成绩
-- >=85,优秀
-- >=60,及格
-- 否则,不及格
create table score(
    id int comment 'ID',
    name varchar(20) comment '姓名',
    math int comment '数学',
    english int comment '英语',
    chinese int comment '语文'
) comment '学员成绩表';
insert into score(id, name, math, english, chinese) VALUES (1, 'Tom', 67, 88, 95 ), (2, 'Rose' , 23, 66, 90),(3, 'Jack', 56, 98, 76);

--
select
    id,
    name,
    ( case when math >= 85 then '优秀' when math >= 60 then '及格' else '不及格' end) '数学',
    ( case when english >= 85 then '优秀' when english >= 60 then '及格' else '不及格' end) '英语',
    ( case when chinese >= 85 then '优秀' when chinese >= 60 then '及格' else '不及格' end) '语文'
from score;
```

## 约束
关键字
```
NOT NULL // 非空约束,限制该字段的数据不能为null
UNIQUE // 唯一约束,保证该字段的数据唯一性,不能重复
PRIMARY KEY // 主键约束,保证该字段的数据唯一性,不能重复,并且不能为null
DEFAUT // 默认值约束,保存数据时,为字段指定值,采取默认值
CHECK // 检查约束,保证字段的数据满足指定的条件(8.0.16版本之后)
FOREIGN KEY // 外键约束,用来让两张表数据之间建立链接,保证数据的一致性和完整性
```

### 外键
**添加外键**
```
create table 表名(
    字段名 数据类型
    ......
    [constraint] [外键名] foreign key (外键字段名称) references 主表名(主表字段名称)
)

alter table emp add constraint fk_emp_dept_id foreign key (dept_id) references dept(id);
```

**删除外键**
```
alter table emp drop foreign key  fk_emp_dept_id;
```

**删除/更新**
```
no action // 在父表中删除或更新，检测是否有外键,有测不允许删除或更新
restrict // 在父表中删除或更新，删除或更新数据时,外键字段的数据必须存在,否则不允许删除或更新
cascade // 在父表中删除或更新，检查是否有外键,有，也删除或更新外键在子表中的记录
set null // 在父表中删除或更新，检查是否有外键,有，则将外键字段的数据设置为null
set default // 父表有变更时,子表将外键设置成一个默认的值

-- 外键的删除/更新行为
alter table emp add constraint fk_emp_dept_id foreign key (dept_id) references dept(id) on update cascade on delete cascade;

alter table emp add constraint fk_emp_dept_id foreign key (dept_id) references dept(id) on update set null on delete set null;
```

## 多表查询
**多对多**
```
create table student(
    id int auto_increment primary key comment '主键ID',
    name varchar(10) comment '姓名',
    no varchar(10) comment '学号'
);
insert into student values (null,'拉丁语','2000100101'),(null,'海口','2000100102'),(null,'段海','2000100103'),(null,'王阳明','2000100104')

create table course(
    id int auto_increment primary key comment '主键ID',
    name varchar(10) comment '课程名称'
) comment '课程表';

insert into course values (null,'java'),(null,'PHP'),(null,'MySQL'),(null,'hadoop')


create table student_course(
    id int auto_increment comment '主键' primary key ,
    studentid int not null comment '学生ID',
    courseid int not null comment '课程ID',
    constraint fk_courseid foreign key (courseid) references course (id),
    constraint fk_studentid foreign key (studentid) references student (id)
)comment '学生课程中间表';

insert into student_course values (null,1,1),(null,1,2),(null,1,3),(null,2,2),(null,2,3),(null,3,4)
```

**一对一**
```
create table tb_user(
    id int auto_increment primary key comment '主键ID',
    name varchar(10) comment '姓名',
    age int comment '年龄',
    gender char(1) comment '1: 男, 2: 女',
    phone char(11) comment '手机号'
) comment '用户信息表';

create table tb_user_edu(
    id int auto_increment primary key comment '主键ID',
    degree varchar(20) comment '学历',
    major varchar(50) comment '专业',
    primaryschool varchar(50) comment '小学',
    middleschool varchar(50) comment '中学',
    university varchar(50) comment '大学',
    userid int unique comment '用户ID',
    constraint fk_userid foreign key (userid) references tb_user(id)
) comment '用户教育信息表';

```

### 内连接
**隐式内连接**()
去重:(distinct)
```
select 字段列表 from 表1,表2 where 条件;

select emp.name, dept.name from emp, dept where emp.dept_id = dept.id;

select e.name, d.name from emp e, dept d where e.dept_id = d.id;
```

**显示内连接**(inner join / join)
```
select 字段列表 from 表1 [inner] join 表2 on 链接条件;

select e.name, d.name from emp e, dept d inner join emp e2 on d.id = e2.dept_id

select e.name, d.name from emp e, dept d join emp e2 on d.id = e2.dept_id
```

### 外连接
**左外连接**(一般都用这个)(left)
```
select 字段列表 from 表1 left join 表2 on 条件;

select e.*, d.name from emp e left  join dept d on e.dept_id=d.id;

```

**右外连接**(right)
```
select 字段列表 from 表1 right [left] join 表2 on 条件;

select d.*, e.name from emp e right join dept d on e.dept_id=d.id;

select d.*, e.name from dept d left join emp e on e.dept_id=d.id;
```

### 自连接(一定要起别名)
```
select 字段列表 from 表A 别名A join 表A 别名B on 条件;

select a.name,b.name from emp a,emp b where a.manager_id= b.id;

select a.name '员工', b.name '领导' from emp a left join emp b on a.manager_id = b.id;
```

### 联合查询(union，union all)
```
select 字段列表 from 表1 ... 
union [all]
select 字段列表 from 表2 ...;

select * from emp where salary < 5000
union
select * from emp where age > 50;
```

### 子查询(嵌套查询)
```
select * from t1 where columm1 = (select column1 from t2);
```

#### 标量子查询
**常用的操作符**
(=, !=, <, >, <=, >=)
```
查询销售部ID
select id from dept where name = '销售部';

select * from emp where dept_id = (select id from dept where name = '销售部');
```

#### 列子查询
**常用的操作符**
```
in // 在指定范围内，多选一
not in // 不在指定范围内
any // 子查询返回列表中，右任意一个满足即可
some // 与any相同，使用some的地方都可以使用any
all // 子查询返回列表中，所有都满足才行

-- 根据ID，查询员工信息
select * from emp where dept_id in (select id from dept where name = '研发部' or name = '销售部' or name = '市场部');


-- 查询比财务部工资高的人
select * from emp where salary > all (select salary from emp where dept_id = (select id from dept where name = '财务部')
);

-- 查询比研发任意一个人工资高
select * from emp where salary > any (select salary from emp where dept_id = (select id from dept where name = '研发部')
);
```

#### 行子查询
**常用的操作符**
```
(=、<>、in、not in)
-- 查询展示的薪资与直系领导
select * from emp where (salary, manager_id) = (select salary, manager_id from emp where name = '展示');
```

#### 表子查询
**常用的操作符**
(in)
```
-- 查询与王五、赵六职位和薪资相同的
select * from emp where (job, salary) in (select job, salary from emp where name = '王五' or name = '赵六');
```

## 事务
**查看**
```
--查看事务情况
select @@autocommit;

-- 自动提交
set @@autocommit = 1;

-- 手动提交
set @@autocommit = 0;
```

**开启事务**
```
start transaction 或 begin;
```

**提交**
```
commit;
```

**回滚**
```
rollback;
```

### 事务四大特性(ACID)
**原子性**
```
Atomicity // 事务是不可分割的，要么全部成功，要么全部失败
```

**一致性**
```
Consistency // 事务完成时，必须使用所有的数据的保持一致状态
```

**隔离性**
```
Isolation // 数据库系统提供的隔离机制，保证事务在不受外部并发操作的独立操作影响的独立环境下运行
```

**持久性**
```
durability // 事务一旦提交或回滚，它对数据库中的数据的改变就是永久的。
```

### 并发事务问题
**脏读**:
一个事务读到另一个事务还没有提交的数据

**不可重复读**:
一个事务先后读取到同一条记录，但两次读取的数据不一致，称之为不可重复读

**幻读**:
一个事务查询时，没有对应的数据，但在插入数据时,又发现这行的数据已经存在

### 事务隔离级别
read uncommitted // 读未提交
可能出现情况
脏读、不可重复读、幻读

read committed // 读已提交
不可重复读、幻读

repeatable read(mysql默认) // 可重复读
幻读

serializable // 串行化

#### 查看运行级别
```
select @@transaction_isolation;
```

#### 设置事务隔离级别
set [session | global] transaction isolation level [read uncommitted | read committed | repeatable read | serializable];
selssion // 代表当前会话，仅对当前窗口有效
global // 代表全局，对所有用户有效


## 存储引擎(LIUNX版)
### MySQL体系结构
1连接层 -->  2服务层 -->  3引擎层 -->  4数据层

### 存储引擎介绍
存储引擎就是存储数据、建立索引、更新/查询数据等技术的实现方式。基于表

-- 查询建表语句 ----默认存储引擎InnoDB
show create table account;

-- 查询当前数据库的存储引擎
show engines ;

### 存储引擎特点
#### InnoDB
**介绍**:
高可靠、高性能通用的存储引擎

**特点**:
DML(增删改)操作遵循ACID模型,支持事务
行级锁，高并发访问性能
支持外键FOREIGN KEY约束，保证数据的完整性和正常


## 索引
**索引结构**:
B+Tree索引  大部分引擎都支持，最常见的索引
hash索引    只有精确匹配索引才有效，不支持范围查询
R-Tree(空间索引)    是MyISAM引擎，用于地理空间类型，使用较少
Full-text(全文索引)   倒非索引，快速匹配文档。类似Lucene，Solr,ES

二叉树: 循序插入，形成列表，查询性能低
红黑树: 大数据情况下，层级深，查询性能慢

**Hash索引**:
只能用于对等比较(=, in)，不支持范围
无法利用索引完成排序操作
查询效率高(无Hash碰撞情况下)


### 索引分类:
主键索引 针对表中主键创建的索引 默认自动创建，只能有一个 primary
唯一索引 避免表数据重复 可以有多个 unique
常规索引 快速定位数据 可以有多个
全文索引 关键字 可以有多个 fulltext

**在innoDB存储中**:
聚集索引    将数据存储放到一块，索引结构的叶子节点保存了行数据   必须有，且只有一个
二级索引    将数据与索引分开存储，索引结构的叶子节点关联的是对应的主键  可以存在多个

规则:主键索引就是聚集索引
     不存在主键，将使用第一个(unique)索引作为聚集索引
     没主键，没合适的唯一索引，自动生成一个rowid作为隐藏的聚集索引


n * 8+(n + 1)* 6=16*1024,指针比key多一个

### 索引语法
**创建索引**:
create [unique | fulltext] index index_name on table_name(index_column_name);

**查看索引**:
show index from table_name;

**删除索引**:
drop index index_name on table_name;

#### 启动mysql
mysql -uroot -p
密码: Itcast@2022


#### 性能分析
**查询频率**:
show global status like 'com_______' // 查看执行频率

**慢查询日志**:
show variabless like 'slow_query_log'; // 查看慢查询是否开启

配置文件(vi /etc/my.cnf)(默认慢查询时间10秒)
show_query_log=1; // 开启慢查询日志
long_query_time=2; // 设置慢查询日志的时间为2秒，超过的是为慢查询

查看慢日志文件记录信息
cat /var/log/mysql/mysql-slow.log


**profile详细**:
show profile 能够在做SQL优化时帮助我们了解时间都消耗到哪里去了
select @@have_profiling; // 查看是否支持profile操作
set profiling=1; // 开启profile，默认为关闭

show profile; // 查看每条SQL的耗时时间
show profile for query_id; // 查看指定query_id的SQl语句各个阶段的耗时情况
show profile cpu for query_id; // 查看指定query_id的SQl语句cpu的使用情况

**explain执行计划**:
explain/desc 字段列表 from 表名 wherr 条件;

**explain执行计划各字段含义**:

***id***:表查询中执行select字句或者是操作表的顺序(id相同，执行顺序从上到下;id不同，值越大执行)

**select_type**: 表示查询的类型，simple表示简单查询，primary表示主查询，union表示并集查询，derived表示派生查询，subquery表示子查询

***type***: 表示链接类型，性能有好到差NULL、system、const、eq_ref、ref、rang、index、all

***possible_keys***: 显示可以应该用在这张表上的索引，一个或多个

key 实际使用的索引，如果为null，则表示没有使用索引
key_len 表示索引中使用的字节数，该值为索引字段最大可能长度，不精确，长度越短越好
rows 表示mysql执行该查询时使用的行数，可能并精确的
fillered 表示返回结果的行数占读取行数的百分比，fillered的值越大越好

#### 索引使用规则
**最左前缀法则**
索引了多列(联合索引),要遵守最左前缀法则。最左前缀法则是查询从索引的最左列开始，且不跳过索引中的列。如果跳跃某一列，后面的索引失效

**范围查询**
联合索引中，出现范围查询(>, <),范围查询右侧的索引失效

#### 索引失效情况
**索引列运算**
不要在索引列上进行运算操作，索引将失效。

**字符串不加引号**
字符串类型字段使用时，不加引号，索引将失效。

**模糊查询**
如果仅仅是尾部模糊匹配,索引不会失效。如果头部模糊匹配，索引失效。

**or链接条件**
用or分开条件，如果or前的条件中的列有索引，而后面的列中没有索引，那么涉及的索引都不会被用到
explain select * from user where id = 1 or age = 21;
解决办法，age建立索引
create index index__user_age on user(age);

**数据分布影响**
如果mysql评估使用索引比全表慢，侧不使用索引

#### SQl提示
explain select * from user use index(idx__user_age) where age = 21;
建议
use index
不要
ignore index
强制
force index

#### 覆盖索引&回表查询
drop index 索引 on 表名; // 删除索引

**覆盖索引**
尽量使用覆盖索引，即索引中包含查询的所有列，避免回表查询
using index condition 查找使用了索引，但需要回表查询数据
using where; using index 查找使用了索引，且不需要回表查询数据

**回表查询**
两层索引

##### 前缀索引
语法
create index ind_xxxx on 表名(column(n));

操作
select count(*) from user; // 查看有多少数据
slelct count(email) from user; // 求取email不为空的值
select coumt(distinct email) from user; //去重
select coumt(distinct email)/ count(*) from user; // 选择性为1
select coumt(distinct emailemail,1,5)/ count(*) from user;
create index ind_email_5 on user(email(5)); //提取5个前缀

##### 单列&联合索引
单列索引: 即一个索引只包含单个列
联合索引: 即一个索引包含多个列









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

















