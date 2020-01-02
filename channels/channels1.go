package main

import (
  "fmt"
  "strconv"
)

func main() {
  const times = 5
  ch := make(chan string)
  
  for i := 0; i < times; i++ {
    go process(strconv.Itoa(i + 1), ch)
  } 
  
  for i := 0; i < times; i++ {
    fmt.Println(<-ch)
  }
}

func process(msg string, ch chan<- string) {
  ch <- msg + " is processed"
}