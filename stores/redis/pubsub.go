package redis

import "github.com/go-redis/redis"

type RedisPubSub struct {
	client *redis.Client
}

func NewRedisPubSub(client *redis.Client) *RedisPubSub {
	return &RedisPubSub{
		client: client,
	}
}

func (rpb *RedisPubSub) Publish(queue string, value []byte) error {
	rpb.client.Publish(queue, value)
	return nil
}

func (rpb *RedisPubSub) Subscribe(queue string, msg chan []byte) error {
	sub := rpb.client.Subscribe(queue)
	if _, err := sub.Receive(); err != nil {
		// handle error
	}

	ch := sub.Channel()

	go func() {
		for {
			data := <-ch
			msg <- []byte(data.Payload)
		}
	}()

	return nil
}
