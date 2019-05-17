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

func (rpb *RedisPubSub) Publish(queue string, value string) error {
	rpb.client.Publish(queue, value)
	return nil
}

// Subscribe subcribes the client to a certain queue, and uses a channel to
// indicate something on that subscription channel has sent a message
func (rpb *RedisPubSub) Subscribe(queueid string, msg chan string) error {
	sub := rpb.client.Subscribe(queueid)
	if _, err := sub.Receive(); err != nil {
		// handle error
		return err
	}

	ch := sub.Channel()

	go func() {
		for {
			data := <-ch
			msg <- data.Payload
		}
	}()

	return nil
}
