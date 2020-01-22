package main

import (
  "net"
  "bufio"
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
      firstLine = false
    }    
  }
  
  body := `<!DOCTYPE html>
           <html lang="en">
           <head>
              <meta charset="UTF-8">
              <title>dir2http</title>
           </head>
           <body>
            <h1>Hello from dir2http!</h1>
           </body>
           </html>`
           
  fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/html\r\n")
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, body)
}