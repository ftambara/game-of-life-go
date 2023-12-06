package main

import (
	"fmt"
	"time"
)

func main() {
	defer timer("main")()
	board := RandomBoard(10000, 10000)
	for i := 0; i < 20; i++ {
		nextBoard := board.advance()
		board = nextBoard
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
