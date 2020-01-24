package main

import (
  "net"
  "bufio"
  "os"
  "io"
  "strings"
  "log"
  "fmt"
  "strconv"
  "path"
  "path/filepath"
)

const INDEX_FILE = "index.html"

var root string

func main() {
  args := os.Args[1:]
  if len(args) < 1 {
    fmt.Println("Usage: <port> [path/to/root:.]")
    os.Exit(1)
  }
  port, err := strconv.Atoi(args[0])
  if err != nil {
    log.Fatalln("Wrong port!")
  }
  if len(args) > 1 {
    root = args[1]
  } else {
    root = "."
  }
  root = filepath.Clean(root)
  
  info, err := os.Stat(root)
  if os.IsNotExist(err) || !info.IsDir() {
    fmt.Println(root, "is not a directory!")
    os.Exit(1)
  }
  
  startServer(port)
}

func startServer(port int) {
  li, err := net.Listen("tcp", ":" + strconv.Itoa(port))
  if err != nil {
    log.Fatalln("Cannot start the server:", err.Error())
  }
  defer li.Close()
  
  log.Println("Starting dir2http server...")
  for {
    conn, err := li.Accept()
    if err != nil {
      log.Fatalln("Cannot accept the client connection:", err.Error())
    }
    go handleRequest(conn)
  }
}

func handleRequest(conn net.Conn) {
  defer conn.Close()
  log.Printf("Serving a request %v...\n", conn.RemoteAddr())
  defer log.Printf("Diconnectiong %v...\n", conn.RemoteAddr())
  
  // parse and clean the requested path
  r := requestedPath(conn)
  if r == "" {
    r = INDEX_FILE
  }
  p := path.Clean(r)
  for strings.HasPrefix(p, "/") {         // remove the leading slash
    p = p[1:]
  }
  if p != r && p + "/" != r {             // redirect to the new cleaned URL
    redirection(conn, "/" + p)
  }
  
  // find the requested file
  fp := filepath.Join(root, p)  
  info, err := os.Stat(fp)
  if os.IsNotExist(err) {
    notfound(conn)
    return
  }
  if info.IsDir() {
    if !strings.HasSuffix(r, "/") {       // a directory ends with a slash
      redirection(conn, "/" + p + "/")
      return
    }
    fp = filepath.Join(fp, INDEX_FILE)
  }
  info, err = os.Stat(fp)
  if os.IsNotExist(err) {
    forbidden(conn)
    return
  }
  
  // read the file
  f, err := os.Open(fp)
  if err != nil {
    servererror(conn, err.Error())
    return
  }
  
  fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", info.Size())
  fmt.Fprint(conn, "\r\n") 
  
  // stream the file
  for {
    b := make([]byte, 1024)
    n, err := f.Read(b)
    if err != nil {
      if err != io.EOF {
        log.Println("Error by reading the file", fp, ":", err.Error())
        return
      }
      break
    }
    fmt.Fprint(conn, string(b[:n]))
  }
}

func requestedPath(conn net.Conn) string {
  scann := bufio.NewScanner(conn)
  var p string
  first := true
  for scann.Scan() {
    ln := scann.Text()
    if ln == "" {
      break
    }
    if first {
      p = strings.Fields(ln)[1][1:]
      first = false
    }    
  }
  return p
}

func redirection(conn net.Conn, loc string) {
  fmt.Fprint(conn, "HTTP/1.1 302 Found\r\n")
  fmt.Fprint(conn, "Content-Length: 0\r\n")
  fmt.Fprintf(conn, "Location: %v\r\n", loc)
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