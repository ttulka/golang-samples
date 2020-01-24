package main

import (
  "net"
  "bufio"
  "os"
  "io"
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
  
  log.Println("Starting dir2http server...")
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
  
  if path == "" {
    path = "index.html"
  }
  pathChanged := false
  for strings.HasPrefix(path, "/") {
    path = path[1:len(path)]
    pathChanged = true
  }
  for strings.HasSuffix(path, "/") {
    path = path[:len(path) - 1]
    pathChanged = true
  }  
  if pathChanged {
    redirection(conn, "/" + path)
  }
  
  info, err := os.Stat(path)
  if os.IsNotExist(err) {
    notfound(conn)
    return
  }
  if info.IsDir() {
    path += "/index.html"
  }
  info, err = os.Stat(path)
  if os.IsNotExist(err) {
    forbidden(conn)
    return
  }
  
  f, err := os.Open(path)
  if err != nil {
    servererror(conn, err.Error())
    return
  }
  
  fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", info.Size())
  fmt.Fprint(conn, "\r\n")
  
  for {
    b := make([]byte, 1024)
    n, err := f.Read(b)
    if err != nil {
      if err != io.EOF {
        log.Println(err.Error())
        return
      }
      break
    }
    fmt.Fprint(conn, string(b[:n]))
  }
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

func redirection(conn net.Conn, loc string) {
  fmt.Fprint(conn, "HTTP/1.1 302 Found\r\n")
  fmt.Fprint(conn, "Content-Length: 0\r\n")
  fmt.Fprintf(conn, "Location: %s\r\n", loc)
  fmt.Fprint(conn, "\r\n")
}

func forbidden(conn net.Conn) {
  fmt.Fprint(conn, "HTTP/1.1 403 Forbidden\r\n")
  fmt.Fprint(conn, "Content-Length: 0\r\n")
  fmt.Fprint(conn, "\r\n")
}

func notfound(conn net.Conn) {
  fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
  fmt.Fprint(conn, "Content-Length: 0\r\n")
  fmt.Fprint(conn, "\r\n")
}

func servererror(conn net.Conn, msg string) {
  fmt.Fprint(conn, "HTTP/1.1 500 Internal Server Error\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(msg))
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, msg)
}