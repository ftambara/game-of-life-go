package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	states [][]CellState
	width  int
	height int
}

func NewBoard(xSize, ySize int, fn func(int, int) CellState) *Board {
	// Use padding to avoid bounds checking
	realXSize := xSize + 2
	realYSize := ySize + 2
	states := make([][]CellState, realYSize)
	underlying := make([]CellState, realXSize*realYSize)
	for y := range states {
		states[y], underlying = underlying[:realXSize], underlying[realXSize:]
	}
	for y := range states {
		for x := range states[y] {
			if x == 0 || x == realXSize-1 || y == 0 || y == realYSize-1 {
				continue
			}
			states[y][x] = fn(x-1, y-1)
		}
	}
	return &Board{states, xSize, ySize}
}

func RandomBoard(xSize, ySize int) *Board {
	return NewBoard(xSize, ySize, func(_, _ int) CellState {
		if rand.Intn(2) == 0 {
			return On
		}
		return Off
	})
}

func (b *Board) at(x, y int) CellState {
	return b.states[y+1][x+1]
}

func (b *Board) set(x, y int, s CellState) {
	b.states[y+1][x+1] = s
}

func (b *Board) advance() *Board {
	nextBoard := NewBoard(b.width, b.height, func(x, y int) CellState {
		count := b.CountNeighbors(x, y)
		return next(b.at(x, y), count)
	})
	return nextBoard
}

func (b *Board) print() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			fmt.Printf("%v  ", b.at(x, y))
		}
		fmt.Print("\n\n")
	}
}

func (b *Board) CountNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if b.at(x+dx, y+dy) == On {
				count++
			}
		}
	}
	return count
}

func (b *Board) Equals(other *Board) bool {
	if b.width != other.width || b.height != other.height {
		return false
	}
	for y := range b.states {
		for x := range b.states[y] {
			if b.at(x, y) != other.at(x, y) {
				return false
			}
		}
	}
	return true
}
