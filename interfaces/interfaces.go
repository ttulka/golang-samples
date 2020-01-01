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
	p := processor{}
	p.Read(make([]byte, 5))
	p.Close()
}

type processor struct {
	readBytes []byte
}

func (p *processor) Read(bytes []byte) (n int, err error) {
	p.readBytes = bytes
	l := len(bytes)

	fmt.Println("Read", l, "bytes")
	return l, nil
}

func (p *processor) Close() error {
	fmt.Println("Closed.")
	return nil
}
