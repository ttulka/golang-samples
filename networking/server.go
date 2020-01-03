package main

import (
  "fmt"
  "net"
  "time"
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
    fmt.Printf("Client %v accepted\n", conn.RemoteAddr())
    conn.SetDeadline(time.Now().Add(time.Second * 20))
    
    req, err := readRequest(conn)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Printf("Client %v: '%v'\n", conn.RemoteAddr(), req)
    
    if err := writeResponse(conn, fmt.Sprintf("Hello, %v!", req)); err != nil {
      fmt.Println(err)
    }
    conn.Close()
  }
}

func readRequest(conn net.Conn) (string, error) {
  buffer := make([]byte, 128)
  n, err := conn.Read(buffer)
  if err != nil {
    return "", err
  }
  return string(buffer[:n]), nil
}

func writeResponse(conn net.Conn, response string) error {
  if _, err := conn.Write([]byte(response)); err != nil {
    return err
  }
  return nil
}