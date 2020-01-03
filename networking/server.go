package main

import (
  "fmt"
  "net"
)

func main() {
  listener, err := net.Listen("tcp", ":4040")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer listener.Close()
  
  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println("Client accepted", conn.RemoteAddr())
    
    conn.Write([]byte("Nice to meet you!"))
    conn.Close()
  }
}