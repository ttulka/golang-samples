package main

import (
	"testing"
) 

func TestHelloMsg(t *testing.T) {
	msg := helloMsg()
  
  if msg != "Hello, world!" {
    t.Errorf("Expected 'Hello, world!', but got '%v'", msg)
  }
}	

func TestReverse(t *testing.T) {
	s := reverse("abc")
  
  if s != "cba" {
    t.Errorf("Expected 'cba', but got '%v'", s)
  }
}	