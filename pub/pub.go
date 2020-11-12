package main

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/RussellLuo/rpubsub"
)

func main() {
	pub := rpubsub.NewPublisher(redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}))

	id, _ := pub.Publish(&rpubsub.PubArgs{
		Topic: "myst",

		MaxLen: 1000,
		Values: map[string]interface{}{
			"greeting":   "deftouch test redis",
			"Hey":        "My value Vicky",
			"test_value": "testvalue for testing",
		},
	})
	// for k, v := range id {
	// 	fmt.Println(k, v)
	// 	fmt.Printf("Sent a message with ID: %s\n", id)
	// }

	fmt.Printf("Sent a message with ID: %s\n", id)

}
