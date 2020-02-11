package main

import (
  "net"
  "io"
  "os"
  "log"
)

func main() {
  if len(os.Args[1:]) <= 0 {
    log.Fatalln("Usage: <host>")
  }
  host := os.Args[1:][0]
  
  listener, err := net.Listen("tcp", "127.0.0.1:12345")
  if err != nil {
    log.Fatalln("Unable to bind to port")
  }
  log.Println("Listening on 127.0.0.1:12345")
  
  for {
    conn, err := listener.Accept()
    log.Println("Received connection")
    if err != nil {
      log.Fatalln("Unable to accept connection")
    }
    go handle(conn, host)
  }
}

func handle(src net.Conn, host string) {
  dst, err := net.Dial("tcp", host + ":80")
  if err != nil {
    log.Fatalln("Unable to connect to our unreachable host")
  }
  defer dst.Close()
  
  go func() {
    // source's output to the destination
    if _, err := io.Copy(dst, src); err != nil {
      log.Fatalln(err)
    }
  }()

  // destination's output back to the source
  if _, err := io.Copy(src, dst); err != nil {
    log.Fatalln(err)
  }
}