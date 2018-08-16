package main

import (
	"fmt"
	"github.com/go-redis/redis"
//	"github.com/go-redis/redis/internal/pool"
//  "github.com/go-redis/redis/internal/proto"
//	"sync"
//	"time"

  "log"
)

func main() {
    // Establish a connection to the Redis server listening on port 6379 of the
    // local machine. 6379 is the default port, so unless you've already
    // changed the Redis configuration file this should work.
//		conn, err := redis.Dial("tcp", "localhost:6379") // does not work
		client := redis.NewClient(&redis.Options{
		  Addr:     "localhost:6379",
		  Password: "", // no password set
		  DB:       0,  // use default DB
		})
		pong, err := client.Ping().Result()
    if err != nil {
        log.Fatal(err)
    } else {
			fmt.Println(pong)
		}

		// This next section is to try out the publish function
		// Go channel which receives messages.
		pubsub := client.Subscribe("A_Channel")
		ch := pubsub.Channel()

    // Wait for confirmation that subscription is created before publishing anything.
    theChannel, err := pubsub.Receive()
    if err != nil {
      panic(err)
    } else {
			fmt.Println("Information: ", theChannel)
		}

		// Consume messages.
		for {
		    msg, ok := <-ch
		    if !ok {
		        break
		    }
		    fmt.Println(msg.Channel, msg.Payload)
		}


    defer client.Close()

    fmt.Println("The listning task is completed")
}
