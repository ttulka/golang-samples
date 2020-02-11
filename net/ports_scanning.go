package main
import (
  "fmt"
  "net"
  "sync"
  "os"
)
func main() {
  var host string
  if len(os.Args[1:]) > 0 {
    host = os.Args[1:][0];
  } else {
    host = "localhost"
  }
  var wg sync.WaitGroup
  for i := 1; i <= 1024; i++ {
    wg.Add(1)
    go func(j int) {
      defer wg.Done()
      address := fmt.Sprintf("%v:%d", host, j)
      conn, err := net.Dial("tcp", address)
      if err != nil {
        return
      }
      conn.Close()
      fmt.Printf("%d open\n", j)
    }(i)
  }
  wg.Wait()
}