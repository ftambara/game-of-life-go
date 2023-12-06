package main

import "fmt"

func main() {
	board := Board{
		// x
		// 0   1   2        y
		{Off, Off, Off}, // 0
		{On, On, Off},   // 1
		{Off, On, Off},  // 2
	}
	board.print()
	for i := 0; i < 3; i++ {
		nextBoard := board.advance()
		nextBoard.print()
		fmt.Print("\n\n")
	}
}

type Board [][]CellState

func newBoard(size int) *Board {
	board := make(Board, size)
	underlying := make([]CellState, size*size)
	for x := range board {
		board[x], underlying = underlying[:size], underlying[size:]
	}
	return &board
}

func (b *Board) at(x, y int) CellState {
	return (*b)[y][x]
}

func (b *Board) set(x, y int, s CellState) {
	(*b)[y][x] = s
}

func (b *Board) advance() *Board {
	nextBoard := newBoard(len(*b))
	for y := range *b {
		for x := range (*b)[y] {
			count := b.countNeighbors(x, y)
			nextBoard.set(x, y, next(b.at(x, y), count))
		}
	}
	return nextBoard
}

func (b *Board) print() {
	for y := 0; y < len(*b); y++ {
		for x := 0; x < len((*b)[0]); x++ {
			fmt.Printf("%v  ", b.at(x, y))
		}
		fmt.Print("\n\n")
	}
}

func (b *Board) countNeighbors(x, y int) int {
	count := 0
	// Bound r and c with board edges
	for xx := max(x-1, 0); xx <= min(x+1, len(*b)-1); xx++ {
		for yy := max(y-1, 0); yy <= min(y+1, len((*b)[y])-1); yy++ {
			if xx == x && yy == y {
				// Given cell, skip
				continue
			}
			if b.at(xx, yy) == On {
				count++
			}
		}
	}
	return count
}

type CellState int

const (
	Off CellState = iota
	On
)

func (c CellState) String() string {
	if c == On {
		return "██"
	} else {
		return "░░"
	}
}

func next(state CellState, neighbor_count int) CellState {
	if state == On {
		if neighbor_count == 2 || neighbor_count == 3 {
			return On
		} else {
			return Off
		}
	} else {
		if neighbor_count == 3 {
			return On
		} else {
			return Off
		}
	}
}