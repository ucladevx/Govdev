package redis

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

const defaultKey = "value" // a default key needed to use RedisStreams

// RedisQueue implements the PublishSubscribe interface, but only publishes to
// one subscriber. Implemented by having Publishers push into a queue, and
// Subscribers read from the queue and pop off from the top.
type RedisQueue struct {
	client *redis.Client
}

// NewRedisQueue creates a new queue and a group associated with that queue.
// Underlying mecahnism uses Redis Streams.
func NewRedisQueue(client *redis.Client, queueID, groupID string) *RedisQueue {
	_, err := client.XGroupCreateMkStream(queueID, groupID, "$").Result()
	if err != nil {
		fmt.Println("Error creating redis group stream")
		os.Exit(1)
	}
	return &RedisQueue{
		client: client,
	}
}

// Subscribe creates a goroutine that subscribes to a RedisStream, based on
// queueID, groupID, consumerID. Sends data values to a msg []byte channel.
func (rq *RedisQueue) Subscribe(queueID, groupID, consumerID string, msg chan []byte) error {
	// Create Subscription

	// Read from stram (do in loop)
	// XREADGROUP GROUP queueGROUP ConsumerID COUNT 1 STREAMS queueID >

	// Acknowledge that message was processed
	// XACK queueID queueGROUP MSGID

	go func() {
		args := &redis.XReadGroupArgs{
			Group:    groupID,
			Consumer: consumerID,
			Streams:  []string{queueID},
			Count:    1,
			Block:    0,
			NoAck:    false,
		}

		for {
			xstreams, err := rq.client.XReadGroup(args).Result()
			if err != nil {
				// handle error, prob by logging
			}
			xstream := xstreams[0]                          // only asking for one stream
			message := xstream.Messages[0]                  // asking for one message
			msgID := message.ID                             // retrieve message ID
			value, err := getBytes(message.Values["value"]) // retrieve value using defaultKey, transform to bytes
			if err != nil {
				// log gob decoding error
			}

			msg <- value // send the value

			// Ack the read
			_, err = rq.client.XAck(queueID, groupID, msgID).Result()
			if err != nil {
				// log ack error
			}
		}
	}()

	return nil
}

// Publish value to a queue, based on queueID.
func (rq *RedisQueue) Publish(queueID string, value []byte) error {
	// XADD queueID * field value
	var m map[string]interface{}
	m["value"] = value
	args := &redis.XAddArgs{
		Stream:       queueID,
		MaxLen:       1, // MAXLEN N
		MaxLenApprox: 1, // MAXLEN ~ N
		ID:           "*",
		Values:       m,
	}

	err := rq.client.XAdd(args).Err()
	if err != nil {
		return err
	}
	return nil
}

func getBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
