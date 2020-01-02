package main

import "fmt"

type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

func main() {
	buffer := make([]byte, 16)
  p := processor{[]byte("data")}
  read, err := p.Read(buffer)
	p.Close()
  
  if err == nil {
    fmt.Println("Buffer:", buffer[:read])
  }
}

type processor struct {
	data []byte
}

func (p *processor) Read(buffer []byte) (n int, err error) {
	copy(buffer, p.data) 
	l := len(p.data)

	fmt.Println("Read", l, "bytes")
	return l, nil
}

func (p *processor) Close() error {
	fmt.Println("Closed.")
	return nil
}
