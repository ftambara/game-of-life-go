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
	states := make([][]CellState, ySize)
	underlying := make([]CellState, xSize*ySize)
	for y := range states {
		states[y], underlying = underlying[:xSize], underlying[xSize:]
	}
	for y := range states {
		for x := range states[y] {
			states[y][x] = fn(x, y)
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
	return b.states[y][x]
}

func (b *Board) set(x, y int, s CellState) {
	b.states[y][x] = s
}

func (b *Board) advance() *Board {
	nextBoard := NewBoard(b.width, b.height, func(_, _ int) CellState {
		return Off
	})
	for y := range b.states {
		for x := range b.states[y] {
			count := b.CountNeighbors(x, y)
			nextBoard.set(x, y, next(b.at(x, y), count))
		}
	}
	return nextBoard
}

func (b *Board) print() {
	for y := range b.states {
		for x := range b.states[y] {
			fmt.Printf("%v  ", b.at(x, y))
		}
		fmt.Print("\n\n")
	}
}

func (b *Board) CountNeighbors(x, y int) int {
	count := 0
	// Bound r and c with board edges
	for xx := max(x-1, 0); xx <= min(x+1, b.width-1); xx++ {
		for yy := max(y-1, 0); yy <= min(y+1, b.height-1); yy++ {
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
