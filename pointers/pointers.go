package main

import "fmt"

type person struct {
  name string
}

func (p person) updateName1(name string) {
  p.name = name
}

func (p *person) updateName2(name string) {
  (*p).name = name
}

func (p *person) updateName3(name string) {
  p.name = name
}

func updateSlice(a []int, v int) {
  a[0] = v
}

func updateArray(arr *[3]int, v int) {
  arr[0] = v
}

func updateString(s *string, v string) {
  *s = v
}

func printPointer(p *string) {
   fmt.Println(&p)
}

func main() {
  jim := person{"Jim"}
  
  jim.updateName1("Jimmy1");  
  fmt.Println(jim)  // no change, passed by value
  
  (&jim).updateName2("Jimmy2A");  
  fmt.Println(jim)
  
  jim.updateName2("Jimmy2B");  
  fmt.Println(jim)
  
  (&jim).updateName3("Jimmy3A");  
  fmt.Println(jim)
  
  jim.updateName3("Jimmy3B");  
  fmt.Println(jim)
  
  
  a := []int{1, 2, 3}
  
  updateSlice(a, 999)
  fmt.Println(a)  
  
  arr := [3]int{11, 22, 33}
  
  updateArray(&arr, 888)
  fmt.Println(arr)
  
  s := "abc"
  
  updateString(&s, "def")
  fmt.Println(s)
  
  sPointer := &s
 
  fmt.Println(sPointer)
  fmt.Println(&sPointer)
  printPointer(sPointer)
}

