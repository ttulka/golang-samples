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
    
    req, err := read(conn)
    if err != nil {
      fmt.Println(err)
      return
    }    
    res, err := proxy(req)
    if err != nil {
      fmt.Println(err)
      return
    }
    conn.Write([]byte(res))
    conn.Close()
  }
}

func proxy(request string) (string, error) {
  conn, err := net.Dial("tcp", "127.0.0.1:4040")
  if err != nil {
    return "", err
  }
  defer conn.Close()
  
  if err := write(conn, request); err != nil {
    return "", err
  }
  return read(conn)
}

func read(conn net.Conn) (string, error) {
  buffer := make([]byte, 128)
  n, err := conn.Read(buffer)
  if err != nil {
    return "", err
  }
  return string(buffer[:n]), nil
}

func write(conn net.Conn, msg string) error {
  if _, err := conn.Write([]byte(msg)); err != nil {
    return err
  }
  return nil
}