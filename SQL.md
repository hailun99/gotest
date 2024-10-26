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


