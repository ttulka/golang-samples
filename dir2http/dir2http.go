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
    
  scann := bufio.NewScanner(conn)
  var path string
  firstLine := true
  for scann.Scan() {
    ln := scann.Text()
    fmt.Println("REQ:", ln)
    if ln == "" {
      log.Println("Disconnecting...")
      break
    }
    if firstLine {
      r := strings.Fields(ln)
      m := r[0]
      p := r[1]
      fmt.Println("***METHOD:", m)
      fmt.Println("***PATH:", p)
      path = p[1:]
      firstLine = false
    }    
  }
  
  var body string
  
  dat, err := ioutil.ReadFile(path)
  if err != nil {
    log.Println(err.Error())
    notfound(conn)
    return
  }
  
  body = string(dat)
           
  fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, body)
}

func notfound(conn net.Conn) {
  fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
  fmt.Fprintf(conn, "Content-Length: 0\r\n")
  fmt.Fprint(conn, "\r\n")
}