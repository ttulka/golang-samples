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
    conn.Write(res)
    conn.Close()
  }
}

func proxy(request []byte) ([]byte, error) {
  conn, err := net.Dial("tcp", "127.0.0.1:4040")
  if err != nil {
    return nil, err
  }
  defer conn.Close()
  
  if err := write(conn, request); err != nil {
    return nil, err
  }
  return read(conn)
}

func read(conn net.Conn) ([]byte, error) {
  buffer := make([]byte, 128)
  n, err := conn.Read(buffer)
  if err != nil {
    return nil, err
  }
  return buffer[:n], nil
}

func write(conn net.Conn, data []byte) error {
  if _, err := conn.Write(data); err != nil {
    return err
  }
  return nil
}