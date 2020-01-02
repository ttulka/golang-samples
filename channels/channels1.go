package main

import (
  "fmt"
  "strconv"
  "time"
)

func main() {
  const times = 5
  ch := make(chan string)
  
  for i := 0; i < times; i++ {
    go process(strconv.Itoa(i + 1), ch)
  } 
  
  for {
    select {
    case m := <-ch:
      fmt.Println(m)
    case <- time.After(time.Second):
      fmt.Println("Closing the channel after timeout was reached")
      close(ch)
      return
    }
  }
}

func process(msg string, ch chan<- string) {
  ch <- msg + " is processed"
}