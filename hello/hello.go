package main

import (
	"fmt"
	"github.com/ttulka/golang-samples/stringutil"
) 

func helloMsg() string {
  return "Hello, world!"
}

func reverse(s string) string {
  return stringutil.Reverse(s)
}

func main() {
	fmt.Println(reverse(helloMsg()))  
}	
