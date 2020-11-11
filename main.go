package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"github.com/RussellLuo/rpubsub"
)

func main() {

	snap := rpubsub.NewRedisSnapshotter(
		redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
		&rpubsub.RedisSnapshotterOpts{
			KeyPrefix:  "node1:",
			Expiration: 24 * time.Hour,
			SavePoint: &rpubsub.SavePoint{
				Duration: time.Second,
				Changes:  1,
			},
		},
	)

	sub := rpubsub.NewSubscriber(&rpubsub.SubOpts{
		NewRedisClient: func() rpubsub.RedisClient {
			return redis.NewClient(&redis.Options{
				Addr: "127.0.0.1:6379",
			})
		},
		Count:       10,
		Snapshotter: snap,
	})

	streams := make(chan rpubsub.Stream)
	sub.Subscribe(streams, "myst")
	defer sub.Unsubscribe()

	for stream := range streams {
		fmt.Printf("Received messages %+v from topic %+v\n", stream.Messages, stream.Topic)
	}
}
