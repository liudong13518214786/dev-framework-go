package main

import "fmt"
import "github.com/garyburd/redigo/redis"

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn error: ", err)
	}

	defer conn.Close()

	// hset
	//_, err = conn.Do("hset", "website", "google", "www.google.com")
	//if err != nil {
	//    fmt.Println("hset error : ", err)
	//}
	//_, err = conn.Do("hset", "website", "123", "www.baidu.com")
	//if err != nil {
	//    fmt.Println("hset error : ", err)
	//}
	//// hmset
	//_, err = conn.Do("hmset", "website", "baidu", "www.baidu.com", "qq", "www.qq.com")
	//if err != nil {
	//    fmt.Println("hmset error:", err)
	//}
	//
	//// hgetall
	//res, err := redis.StringMap(conn.Do("hgetall", "website"))
	//if err == nil {
	//    fmt.Println(res)
	//}else {
	//    fmt.Println("hgetall error: ", err)
	//}
	//
	//// hvals
	//rs, err := redis.Strings(conn.Do("hvals", "website"))
	//if err == nil {
	//    fmt.Println(rs)
	//}else {
	//    fmt.Println(err)
	//}
	//
	// hkeys
	rk, err := redis.Strings(conn.Do("hkeys", "website"))
	if err == nil {

		for _, v := range rk {
			fmt.Println(v)
			_, xx := conn.Do("HDEL", "website", v)
			fmt.Println(xx)

		}
	} else {
		fmt.Println(err)
	}
	rk1, _ := redis.Strings(conn.Do("hkeys", "website"))
	fmt.Println(rk1)
	//
	//// hstrlen
	//l, err := redis.Int(conn.Do("hstrlen", "website", "google"))
	//if err == nil {
	//    fmt.Println(l)
	//}else {
	//    fmt.Println(err)
	//}
	//// hlen
	//l, err = redis.Int(conn.Do("hlen", "website"))
	//if err == nil {
	//    fmt.Println(l)
	//}else {
	//    fmt.Println(err)
	//}
}
