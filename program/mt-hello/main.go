package main

import (
	"os"
	"runtime"
)

func main() {
	done := make(chan struct{})
	go func() {
		_, _ = os.Stdout.Write([]byte("hello world!\n"))
		close(done)
	}()
	<-done

	runtime.GC()
	_, _ = os.Stdout.Write([]byte("completed!\n"))
}
