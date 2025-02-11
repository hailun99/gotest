# 添加go mod
```
go mod init demo
go mod tidy
```

# 链接redis
```
var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.66.137:6379", // Redis地址
		Password: "123456",              // 密码（如果没有则留空）
		DB:       0,                     // 使用默认DB
	})
	// 测试连接是否成功
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
	fmt.Println("链接redis出现错误,错误信息: %v", err)
	}
	fmt.Println(pong)
}
```

# 关闭Redis
```
// 关闭Redis连接
err = rdb.Close()
if err != nil {
    panic(err)
}
```

# 基本命令
## Keys():根据正则获取keys
```
keys, err := rdb.Keys(ctx, "*").Result()
```

## Type():获取key对应值得类型
```
vType, err := rdb.Type(ctx, "key").Result()
```

## Del():删除缓存项
```
n, err := rdb.Del(ctx, "key1", "key2").Result()
```

## Exists():检测缓存项是否存在 
只传入一个key，判断是否大于零即可
```
n, err := rdb.Exists(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
```

## Expire(),ExpireAt():设置有效期
Expire()方法是设置某个时间段(time.Duration)后过期
ExpireAt()方法是在某个时间点(time.Time)过期失效
```
res, err := rdb.Expire(ctx, "key", time.Minute*2).Result()

res, err := rdb.ExpireAt(ctx, "key", time.Now()).Result()
```
## TTL(),PTTL():获取有效期
```
// 设置一分钟有效期
	rdb.Expire(ctx, "key", time.Minute)

	// //获取剩余有效期,单位:秒(s)
	ttl, err := rdb.TTL(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(ttl)

	// //获取剩余有效期,单位:毫秒(ms)
	pttl, err := rdb.PTTL(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pttl)
```

## DBSize():查看当前数据库key的数量
```
num, err := rdb.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("数据库有 %v 个缓存项\n", num)
```

## FlushDB():清空当前数据
```
res, err := rdb.FlushDB(ctx).Result()
```

## FlushAll():清空所有数据库
```
res, err := rdb.FlushAll(ctx).Result()
```

# 字符串类型
## Set()
```
// 0 表示永不过时
err = rdb.Set(ctx, "key1", "value1", 0).Err()

err = rdb.Set(ctx, "key2", "value2", time.Minute*10).Err()
```

## SetEX()
设置键的同时，设置过期时间。不管该key是否已经存在缓存中直接覆盖
```
err = rdb.SetEX(ctx, "key", "value", time.Hour*2).Err()
```

## SetNX()
当key不在的时候才设置
```
res, err := rdb.SetNX(ctx, "key", "value", time.Minute).Result()
```

## Get()
```
err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %v\n", val)

	val2, err := rdb.Get(ctx, "key-not-exist").Result()
	if err == redis.Nil {
		fmt.Println("key不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("key: %v\n", val2)
	}
```

## GetRange():字符串截取
```
err := rdb.Set(ctx, "key", "Hello World!", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.GetRange(ctx, "key", 1, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %v\n", val)  key:ell
	fmt.Printf("%q\n", val)  "ell"
```

## Incr():+1
对数字进行加1
```
val, err := rdb.Incr(ctx, "number").Result()
```

## IncrBy():减-1
```
val, err := rdb.IncrBy(ctx, "number", 1).Result()
```

## Decr():减-1
```
val, err := rdb.Decr(ctx, "number", 12).Result()
```

## DecrBy():指定步长减少
```
val, err := rdb.DecrBy(ctx, "number", 12).Result()
```

## Append:追加
```
err := rdb.Set(ctx, "key", "hello", 0).Err()
length, err := rdb.Append(ctx, "key", "world!").Result()
```

## strLen:获取长度
```
length, err := rdb.StrLen(ctx, "key").Result()
```

# 列表
## LPush:将元素从左侧压入链表
```
n, err := rdb.LPush(ctx, "list", 1, 2).Result()
```

## LPush:将元素从右侧压入链表
```
n, err := rdb.RPush(ctx, "list", 8, 9).Result()
```

## LInsert():在某个位置插入新元素
LInsertBefore() : 从前面插入
LInsertAfer() : 从后面插入
```
err := rdb.LInsert(ctx, "key", "before", "100", 123)
```

## LSet():设置某个元素的值
```
err := rdb.LSet(ctx, "list", 1, 100).Err()
```

## LLen():获取链表元素个数
```
length, err := rdb.LLen(ctx, "list").Result()
```

## LIndex():获取链表下标对应的元素
```
val, err := rdb.LIndex(ctx, "list", 1).Result()
```

## LRange():获取某个选定范围的元素集
```
vals, err := rdb.LRange(ctx, "list", 0, 2).Result()
```

## 弹出数据
LPop: 左边 
```
val, err := rdb.LPop(ctx, "list").Result()
```
RPop: 右边
```
val, err := rdb.RPop(ctx, "list").Result()
```

## LRem():根据值移除元素
```
n, err := rdb.LRem(ctx, "list", 2, "100").Result()
```

# 集合(set)类型
1. 元素不能重复，保持唯一性
2. 元素无序，不能使用索引(下标)操作
## SAdd: 添加元素
```
rdb.SAdd(ctx, "team", "kode", "jorden")
```

## SPop():随机获取一个元素
```
val, err := rdb.SPop(ctx, "team").Result()
```

## SRem():删除集合里指定的值
```
n, err := rdb.SRem(ctx, "team", "kobe", "v2").Result()
```

## SMembers():获取所有成员
```
vals, err := rdb.SMembers(ctx, "team").Result()
```

## SIsMember():判断元素是否在集合中
```
exists, err := rdb.SIsMember(ctx, "team", "jorden").Result()
```

## SCard():获取集合元素个数
```
total, err := rdb.SCard(ctx, "team").Result()
```

## SUnion():并集,SDiff():差集,SInter():交集
并集
```
union, err := rdb.SUnion(ctx, "setA", "setB").Result()
```
差集
```
diff, err := rdb.SDiff(ctx, "setA", "setB").Result()
```
交集
```
inter, err := rdb.SInter(ctx, "setA", "setB").Result()
```

# 有序集合(zset)类型
## ZAdd():添加元素
```
	zadd, err := rdb.ZAdd(ctx, "myzset", &redis.Z{
		Score:  1,
		Member: "one",
	}, &redis.Z{
		Score:  2,
		Member: "two",
	}).Result()
```

## ZRange获取元素
```
zr, err := rdb.ZRange(ctx, "myzset", 0, -1).Result()
```

## ZScore获取元素分数
```
zs, err := rdb.ZScore(ctx, "myzset", "one").Result()
```

## ZRangeWithScores获取有序集合的元素及其分数
```
zrs, err := rdb.ZRangeWithScores(ctx, "myzset", 0, -1).Result()
```

## 删除有序集合中的元素
```
zrem, err := rdb.ZRem(ctx, "myzset", "one").Result()
```

## 获取有序集合的元素数量
```
zca, err := rdb.ZCard(ctx, "myzset").Result()
```

## 获取有序集合中指定分数范围的元素
```
zrby, err := rdb.ZRangeByScore(ctx, "myzset", &redis.ZRangeBy{
		Min: "1",
		Max: "2",
	}).Result()
```

# 哈希
## HSet():设置
```
rdb.HSet(ctx, "user", "key", "value", "key2", "value2")
	rdb.HSet(ctx, "user", []string{"key3", "value3", "key4", "value4"})
	rdb.HSet(ctx, "user",map[string]interface{}{"key5":"value5", "key6":"value6"})
```

## HMset():批量设置
```
rdb.HMSet(ctx, "user", map[string]interface{}{"name": "kevin", "age": 18, "address": "北京"})
```

## HGet():获取某个元素
```
address, err := rdb.HGet(ctx, "user", "address").Result()
```

## HGetAll():获取全部元素
```
user, err := rdb.HGetAll(ctx, "user").Result()
```

## HDel():删除某个元素
```
res, err := rdb.HDel(ctx, "user", "name", "agae").Result()
```

## HExists():判断元素是否存在
```
res, err := rdb.HExists(ctx, "user", "address").Result()
```

## HLen():获取长度
```
res, err := rdb.HLen(ctx, "user").Result()
```



## 订阅
订阅一个频道
```
SUBSCRIBE [频道名称]
```

发布一个频道
```
publish [频道名称]  [发送的消息]
```

## stream流
* 表示自动生成消息的id
```
XADD geekhour * course redis

XADD geekhour 1-0 course doker
```

查看有多少条信息
```
XLEN geekhour
```

查看内容
-表示最后一位 +表示开始位
```
XRANGE geekhour - +
```

删除
```
XDEL geekhour [id]
```

删除所有的信息
0表示所有信息
```
XIPIM geekhour MAXLEN 0
```

读取指定几条数据
阻塞1000秒
从0开始
```
XREAD COUNT 2 BLOCK 1000 STREAMS geekhour 0
```

创建消费者组
```
XGROUP CREATE geekhour group1 [id]
```

查看消费者组
```
XIFNFO GROUPS geekhour
```

创建消费者
```
XGROUP CREATECONSUMER geekhour group1 consumer1
```

读取信息
> 表示读取这个消息中最新的消息
```
XREADGROUP GROUP group1 consumer1 COUNT 2 BLOCK 3000 STREAMS geekhour >
```

## Geospatial
```
GEOADD [地理信息名字] [经度纬度] [城市名字]

GEOADD city 116.405385 39.904989 beijing
```

获取某个地方的位置信息
```
GEOPOS city beijing
```

计算两地的距离
```
GEODIST city bejing shanghai
```

获取范围内的城市
```
GEOSEARCH city FROMMEMBER shanghai BYRADIUS 300 KM
```

## HyperLogLog
添加/创建
```
DFADD course git docker redis
```
查看
```
PFCOUNT course
```

合并
```
PFMERGE result course course2
```

## Bitmap
设置下标
```
SETBIT dianzan 0 1

SET dianzan "\xF0"
```
查看
```
GETBIT dianzan 0\1
```

统计有多少个1
```
BITCOUNT dianzan
```

查看某个数出现的位置
```
BITPOS dainzan 0\1
```
