package main

import (
  "fmt"
  "net"
  "time"
)

func main() {
  start := time.Now()
  conn, err := net.Dial("tcp", "127.0.0.1:4040")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer conn.Close()
  
  buffer := make([]byte, 128)
  n, err := conn.Read(buffer)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  elapsed := time.Now().Sub(start)
  
  fmt.Printf("Server reponse: '%v' took %v ms\n", string(buffer[:n]), elapsed)
}