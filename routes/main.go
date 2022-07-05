package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func main() {
	//	http.HandleFunc("/", restHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8080", nil)
}

func restHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "GET called")
	} else if r.Method == "POST" {
		fmt.Fprintln(w, "POST called")
	} else if r.Method == "PUT" {
		fmt.Fprintln(w, "PUT called")
	} else if r.Method == "DELETE" {
		fmt.Fprintln(w, "DELETE called")
	}

}

func pingHandler(w http.ResponseWriter, _ *http.Request) {

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
				Addr: "redis:6379",
		//Addr: "localhost:6379",
		DB:   0,
	})

	pong, err := rdb.Ping(ctx).Result()

	fmt.Println(pong, err)
}

func job_testHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "successful job queuing")
	// todo: クエリでジョブを受け取れるようにする
}

func ExampleClient() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		//		Addr: "redis:6379",
		Addr: "localhost:6379",
		DB:   0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
