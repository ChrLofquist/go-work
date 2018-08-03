package main

import "net"
import "fmt"
//import "bufio"
import "strings" // only needed below for sample processing

func main() {
//  var message []string

  buffer := make([]byte, 64)

  fmt.Println("Launching server...")

  // listen on all interfaces
  //ln, _ := net.Listen("tcp", ":8081")

  // listen to the specific localhost port
//  ln, _ := net.Listen("tcp", "127.0.0.1:8081")
  ln, err := net.Listen("tcp", "0.0.0.0:8081")
  if err != nil {
    fmt.Println("Listen:", err.Error())
    return
  }
  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  countPackets := 0
  for {

    _, err := conn.Read(buffer)
    message := string(buffer)

    if strings.Count(message,"STOP")>0 {
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
    if countPackets%100==0{
      fmt.Printf("Packets (%v): %s \n", countPackets , message)
    }
  }
}