package main

import (
	"fmt"

	myredis "github.com/go-redis/redis"

	"github.com/RussellLuo/rpubsub"
)

func main() {

	sub := rpubsub.NewSubscriber(&rpubsub.SubOpts{
		NewRedisClient: func() rpubsub.RedisClient {
			return myredis.NewClient(&myredis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
			})
		},
		Count: 100,
	})

	streams := make(chan rpubsub.Stream)
	sub.Subscribe(streams, "myfirst")
	defer sub.Unsubscribe()

	for stream := range streams {
		fmt.Printf("Received messages %+v from topic %+v\n", stream.Messages, stream.Topic)
	}

}
