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
}

