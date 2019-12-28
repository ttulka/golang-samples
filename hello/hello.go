package main

import "fmt"  

func main() {
	fmt.Printf("hello, world\n")
  fmt.Printf(c()) // 2
}

func c() (i int) {
    defer func() { i++ }()
    return 1
}
