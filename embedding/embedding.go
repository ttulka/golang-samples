package main

import (
	"fmt"
) 

type ReaderData struct {
    Read int
}

type WriterData struct {
    Written int
}

type ProcessorData struct {
    *ReaderData
    *WriterData
}

func main() {
	proc := ProcessorData{&ReaderData{123}, &WriterData{456}}
	
	fmt.Println(proc.Read, proc.Written)
}