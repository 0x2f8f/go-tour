package main

import "fmt"
import "github.com/go-redis/redis"
//import "reflect"

func main() {
	redisClient := GetRedisNewClient();
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}

func GetRedisNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client

	//fmt.Println(reflect.TypeOf(client))

//	pong, err := client.Ping().Result()
//	fmt.Println(pong, err)
	// Output: PONG <nil>
}