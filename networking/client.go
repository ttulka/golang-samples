package main

import (
  "fmt"
  "net"
  "os"
  "time"
)

var port, msg string

func init() {
  args := os.Args[1:]
  if len(args) < 2 {
    fmt.Println("Usage: <port> <message>")
    os.Exit(2)
  }
  port = args[0]
  msg  = args[1]
}

func main() {
  start := time.Now()
  
  conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", port))
  if err != nil {
    fmt.Println(err)
    return
  }
  defer conn.Close()
  
  conn.Write([]byte(msg))
  
  res, err := readResponse(conn)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  elapsed := time.Now().Sub(start)
  
  fmt.Printf("Server reponse: '%v' took %v ms\n", res, elapsed)
}

func readResponse(conn net.Conn) (string, error) {
  buffer := make([]byte, 128)
  n, err := conn.Read(buffer)
  if err != nil {
    return "", err
  }
  return string(buffer[:n]), nil
}