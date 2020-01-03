package main

import (
  "fmt"
  "net"
  "time"
  "math/rand"
)

func init() {
  rand.Seed(time.Now().UnixNano())
}

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
    conn.SetDeadline(time.Now().Add(time.Second * 20))
    go process(conn)
  }
}

func process(conn net.Conn) {
  fmt.Println("Proxying a client", conn.RemoteAddr())
      
  req, err := read(conn)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  sleepLong()
  
  res, err := proxy(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  conn.Write(res)
  conn.Close()
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

func sleep() {
  n := rand.Intn(100)
  time.Sleep(time.Duration(n)*time.Millisecond)
}

func sleepLong() {
  n := rand.Intn(10) + 5
  time.Sleep(time.Duration(n)*time.Second)
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