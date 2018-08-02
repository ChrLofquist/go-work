package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
  //ln, _ := net.Listen("tcp", ":8081")

  // listen to the specific localhost port
  ln, _ := net.Listen("tcp", "127.0.0.1:8081")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  countPackets := 0
  for {

    // will listen for message to process ending in newline (\n)
    message, err := bufio.NewReader(conn).ReadString('\n')
    if strings.TrimSpace(string(message))== "STOP" {
      fmt.Println("Exiting TCP Server - bye")
      ln.Close()
      return
    }
    // Error checking should be before the above, but ...
    if err != nil {
      fmt.Println(err)
      return
    }

    // output message received
    countPackets++
    if countPackets%1000==0{
      fmt.Printf("Packets (%v): %s", countPackets, string(message))
    }
    // send new string back to client
    conn.Write([]byte(message + "\n"))

  }

}
