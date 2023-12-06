package main

import "fmt"

func main() {
	board := Board{
		{Off, Off, Off, Off},
		{Off, Off, On, Off},
		{Off, Off, On, Off},
		{Off, Off, On, Off},
		{Off, Off, Off, Off},
	}
	board.print()
	fmt.Print("\n\n")
	for i := 0; i < 3; i++ {
		nextBoard := board.advance()
		nextBoard.print()
		fmt.Print("\n\n")
		board = *nextBoard
	}
}

type Board [][]CellState

func NewBoard(xSize, ySize int) *Board {
	board := make(Board, ySize)
	underlying := make([]CellState, xSize*ySize)
	for y := range board {
		board[y], underlying = underlying[:xSize], underlying[xSize:]
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
	nextBoard := NewBoard(len((*b)[0]), len(*b))
	for y := range *b {
		for x := range (*b)[y] {
			count := b.CountNeighbors(x, y)
			nextBoard.set(x, y, next(b.at(x, y), count))
		}
	}
	return nextBoard
}

func (b *Board) print() {
	for y := range *b {
		for x := range (*b)[y] {
			fmt.Printf("%v  ", b.at(x, y))
		}
		fmt.Print("\n\n")
	}
}

func (b *Board) CountNeighbors(x, y int) int {
	count := 0
	// Bound r and c with board edges
	for xx := max(x-1, 0); xx <= min(x+1, len((*b)[0])-1); xx++ {
		for yy := max(y-1, 0); yy <= min(y+1, len(*b)-1); yy++ {
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

func (b *Board) Equals(other *Board) bool {
	if len(*b) != len(*other) {
		return false
	}
	for y := range *b {
		if len((*b)[y]) != len((*other)[y]) {
			return false
		}
		for x := range (*b)[y] {
			if (*b)[y][x] != (*other)[y][x] {
				return false
			}
		}
	}
	return true
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
