package main

import "testing"

func TestProcessorRead(t *testing.T) {
  const test = "test"
  
	buffer := make([]byte, 32)
  p := processor{[]byte(test)}
  read, err := p.Read(buffer)
	p.Close()
  
  if err != nil {
    t.Error("Cannot read bytes", err)
  }
  if read != len(test) {
    t.Errorf("Count of read bytes doesn't macht. Expected '%v', got '%v'", len(test), read)
  }
  got := string(buffer[:read])
  if got != test {
    t.Errorf("Read bytes doesn't macht. Expected '%v', got '%v'", test, got)
  }
}
