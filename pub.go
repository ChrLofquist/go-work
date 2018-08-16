package main

import (
	"fmt"
	"github.com/go-redis/redis"
//	"github.com/go-redis/redis/internal/pool"
//  "github.com/go-redis/redis/internal/proto"
//	"sync"
	"time"
  "strconv"
  "log"
)

func main() {
//	var comment []string

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
		client.Set("key", "any_value", 0).Err()
		val, err := client.Get("key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)

    // This is the value that has an expire time to it
		seconds := 100
		client.SetNX("HeartBeat", "On", time.Duration(seconds)*time.Second).Err()

		val, err = client.Get("count").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)

		// This next section is to try out the publish function
		//pubsub := client.Subscribe("A_Channel")
		err = client.Publish("A_Channel", "Hello, is there anybody outthere").Err()
    if err != nil {
      panic(err)
    }

		// Go channel which receives messages.
		pubsub := client.Subscribe("A_Channel")
//		ch := pubsub.Channel()

    // Wait for confirmation that subscription is created before publishing anything.
    theChannel, err := pubsub.Receive()
    if err != nil {
      panic(err)
    } else {
			fmt.Println("Information: ", theChannel)
		}

		// Compose messages.
		sum := 0
		for i := 0; i < 10; i++ {
		  sum += i
			comment := "PI: "+ strconv.Itoa(i) + "-->" + strconv.Itoa(sum)
			time.Sleep(1 * time.Second)
			err = client.Publish("A_Channel", comment).Err()
	    if err != nil {
	      panic(err)
	    }

	  }


    // Importantly, use defer to ensure the connection is always properly
    // closed before exiting the main() function.
    defer client.Close()

    // Send our command across the connection. The first parameter to Cmd()
    // is always the name of the Redis command (in this example HMSET),
    // optionally followed by any necessary arguments (in this example the
    // key, followed by the various hash fields and values).
//     resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
    // Check the Err field of the *Resp object for any errors.
//    if resp.Err != nil {
//        log.Fatal(resp.Err)
//    }

    fmt.Println("The eagle has landed")
}
