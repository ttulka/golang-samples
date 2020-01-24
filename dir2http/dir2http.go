package main

import (
  "net"
  "bufio"
  //"io"
  "io/ioutil"
  "strings"
  "log"
  "fmt"
)

func main() {
  li, err := net.Listen("tcp", ":1234")
  if err != nil {
    log.Fatalln(err.Error())
  }
  defer li.Close()
  
  for {
    conn, err := li.Accept()
    if err != nil {
      log.Fatalln(err.Error())
    }
    go handle(conn)
  }  
}

func handle(conn net.Conn) {
  defer conn.Close()
  log.Println("Serving a request...")
    
  path := request(conn)
  
  dat, err := ioutil.ReadFile(path)
  if err != nil {
    log.Println(err.Error())
    notfound(conn)
    return
  }
             
  fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(dat))
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, string(dat))
}

func request(conn net.Conn) string {
  scann := bufio.NewScanner(conn)
  var path string
  firstLine := true
  for scann.Scan() {
    ln := scann.Text()
    if ln == "" {
      log.Println("Disconnecting...")
      break
    }
    if firstLine {
      path = strings.Fields(ln)[1][1:]
      firstLine = false
    }    
  }
  return path
}

func notfound(conn net.Conn) {
  fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
  fmt.Fprintf(conn, "Content-Length: 0\r\n")
  fmt.Fprint(conn, "\r\n")
}