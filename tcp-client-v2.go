package main

import  (
  "net"
  "fmt"
  "bufio"
  "os"
  "flag"
  "log"
  "time"
)

var TargetFile string = ".\\MastSimulatedRaw\\Scenario-High.txt"

func showHelp() {
  fmt.Println(`
Usage: readline textfile [OPTIONS]

Options:
  -i, --ip 192.168.0.40      IP address to transmit (default is localhost)
  -p, --port 8000            transmit port (default is 8081)
  -h, --help                 prints this usage guide
    `)
}

func checkFile(path string) (countLines int) {
  inFile, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
    log.Fatal(err)
    return
  }
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)
  countLines = 0
  for scanner.Scan() {
    countLines++
  }
  inFile.Close()
  return countLines
}

// func checkConnection() {
//   // connect to this socket
//   conn, err := net.Dial("tcp", "127.0.0.1:8081")
//   if err != nil {
//     fmt.Println("DialTCP:", err.Error())
//     return
//   }
// }

func main() {
  var sName string
  flag.StringVar(&sName, "", TargetFile, "")

  var sIP string
  flag.StringVar(&sIP, "i", "127.0.0.1", "")
  flag.StringVar(&sIP, "IP", "127.0.0.1", "")

  var sPort string
  flag.StringVar(&sPort, "p", "8081", "")
  flag.StringVar(&sPort, "port", "8081", "")

  var sHelp bool
  flag.BoolVar(&sHelp, "h", false, "")
  flag.BoolVar(&sHelp, "help", false, "")

  flag.Parse()

  if sName !="" {
    fmt.Println("Default test file is : ", sName)
    // something not right with the transfer of the textfile name
    // requires a little more work here
  }

  if sIP != "" {
      fmt.Println("Default IP address is: ", sIP)
  }
  if sPort != "" {
      fmt.Println("Default Port number is: ", sPort)
  }
  if sHelp {
    showHelp()
    return
  }

  findMax := checkFile(TargetFile)
  fmt.Println("Transmitting : ", findMax, " packets")

  // open the input text file again and get ready for transmission
  inFile, err := os.Open(TargetFile)
  scanner := bufio.NewScanner(inFile)
  scanner.Split(bufio.ScanLines)

  // connect to this socket
  // checkConnection()
  conn, err := net.Dial("tcp", "127.0.0.1:8081")
  if err != nil {
    fmt.Println("DialTCP:", err.Error())
    return
  }

  // tranmit one packet per line in the text file
  j := 0
  start := time.Now()
  for scanner.Scan() {
//    fmt.Println(scanner.Text())

    // send to socket
    fmt.Fprintf(conn, scanner.Text() + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    if j%1000==0 {
      fmt.Print("Message from server: "+message)
    }
    j++
    if j == findMax {
      fmt.Fprintf(conn, "STOP")
      fmt.Println("TCP client closed.")
      conn.Close()
      end := time.Now()
      fmt.Println("Transmission time: ", end.Sub(start))
      return
    }
  }

}
