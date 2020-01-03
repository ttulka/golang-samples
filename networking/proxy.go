package main

import (
  "fmt"
  "net"
)

func main() {
  listener, err := net.Listen("tcp", ":4041")
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
    fmt.Println("Proxying a client", conn.RemoteAddr())
    
    response, err := proxy()
    if err != nil {
      fmt.Println(err)
      return
    }
    conn.Write([]byte(response))
    conn.Close()
  }
}

func proxy() (string, error) {
  conn, err := net.Dial("tcp", "127.0.0.1:4040")
  if err != nil {
    return "", err
  }
  defer conn.Close()
  
  buffer := make([]byte, 128)
  n, err := conn.Read(buffer)
  if err != nil {
    return "", err
  }  
  return string(buffer[:n]), nil
}