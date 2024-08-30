package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// // redisdb,err := RedisInit("192.168.10:1234")

// 创建一个客户端
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

// string 操作
func StringTest(redisdb *redis.Client) {
	// 在客户端里设置一个key ,value, 与一个过期时间
	err := redisdb.Set("key1", "Hello", 5*time.Second).Err
	if err != nil {
		fmt.Println(err)
		return
	}
	err = redisdb.Set("key1", "Hello", 0*time.Second).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 获取kye1的value
func StringTest1(redisdb *redis.Client) {
	result, err := redisdb.Get("key1").Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
}

// 给kye1赋值value
func StringTest2(redisdb *redis.Client) {
	result, err := redisdb.GetSet("key1", "world").Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
} // 返回上一次的值hello,设置key1新值world

// 返回多个value
func StringTest4(redisdb *redis.Client) {
	result, err := redisdb.MGet("key1", "key2").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range result {
		fmt.Println(value)
	}
} //hello1
//hello2

// 获取key ，0表示不过时
func StringTest5(redisdb *redis.Client) {
	// 设置key1
	result, err := redisdb.SetNX("key1", "hello3", 0).Result
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
} // false 因为kye1已经存在

// 设置多个key valeu
func StringTest6(redisdb *redis.Client) {
	result, err := redisdb.MSet("key1", "h1", "key2", "h2", "key3", "h3").Result
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// 设置key1-1 增加1
func StringTest7(redisdb *redis.Client) {
	redisdb.Set("key1", "1", 0)
	result, err := redisdb.Incr("key1").Result
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
} // 2

// 增加key1
func StringTest8(redisdb *redis.Client) {
	redisdb.Set("key1", "1", 0)
	result, err := redisdb.IncrBy("key1", 10).Result
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
} // kye1 变为11

// 减1操作
func StringTest9(redisdb *redis.Client) {
	redisdb.Set("key1", "1", 0)
	result, err := redisdb.Decr("kye1").Result
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
} // 9

// 减 多个
func StringTest10(redisdb *redis.Client) {
	redisdb.Set("key1", "10", 0)
	result, err := redisdb.DecrBy("kye1", 4).Result
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
} // 6

// 增加
func StringTest11(redisdb *redis.Client) {
	redisdb.Set("key1", "abc", 0)
	result, err := redisdb.Append("key1", "def").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
} // 返回6,append后字符串"abcdef"长度为6

func StringTest10086(redisdb *redis.Client) {

}
