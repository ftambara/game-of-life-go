package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	cells  [][]*Cell
	width  int
	height int
}

func NewEmptyBoard(xSize, ySize int) *Board {
	// Use padding to avoid bounds checking
	realXSize := xSize + 2
	realYSize := ySize + 2
	cells := make([][]*Cell, realYSize)
	underlying := make([]*Cell, realXSize*realYSize)
	for y := range cells {
		cells[y], underlying = underlying[:realXSize], underlying[realXSize:]
	}
	// Fill padding with off cells
	// First and last rows
	for x := range cells[0] {
		cells[0][x] = &Cell{state: Off}
		cells[ySize+1][x] = &Cell{state: Off}
	}
	// First and last columns
	for y := range cells {
		cells[y][0] = &Cell{state: Off}
		cells[y][xSize+1] = &Cell{state: Off}
	}

	return &Board{cells, xSize, ySize}
}

func NewBoard(xSize, ySize int, fn func(int, int) CellState) *Board {
	board := NewEmptyBoard(xSize, ySize)
	for y := range board.cells {
		for x := range board.cells[y] {
			if x == 0 || y == 0 || x == xSize+1 || y == ySize+1 {
				board.cells[y][x] = &Cell{state: Off}
				continue
			}
			board.cells[y][x] = &Cell{state: fn(x-1, y-1)}
		}
	}
	// Count neighbors
	for y := range board.cells {
		for x := range board.cells[y] {
			// Count manually at the edges
			var count int
			if x == 0 || y == 0 || x == xSize+1 || y == ySize+1 {
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dx == 0 && dy == 0 {
							continue
						}
						if x+dx < 0 || x+dx > xSize+1 || y+dy < 0 || y+dy > ySize+1 {
							continue
						}
						if board.cells[y+dy][x+dx].state == On {
							count++
						}
					}
				}
				board.cells[y][x].neighborCount = count
				continue
			}

			board.cells[y][x].neighborCount = board.CountNeighbors(x-1, y-1)
		}
	}
	return board
}

func RandomBoard(xSize, ySize int) *Board {
	return NewBoard(xSize, ySize, func(_, _ int) CellState {
		if rand.Intn(2) == 0 {
			return On
		}
		return Off
	})
}

func (b *Board) At(x, y int) *Cell {
	return b.cells[y+1][x+1]
}

func (b *Board) Set(x, y int, s *Cell) {
	b.cells[y+1][x+1] = s
}

func (b *Board) Advance() *Board {
	nextBoard := NewEmptyBoard(b.width, b.height)
	// Create next board
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			oldCell := b.At(x, y)
			nextBoard.Set(x, y, &Cell{
				state:         oldCell.Next(),
				neighborCount: oldCell.neighborCount,
			})
		}
	}

	// Update neighbor counts
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			newState := nextBoard.At(x, y).state
			if b.At(x, y).state == newState {
				// No change, skip
				continue
			}

			// Do conditional outside of loop
			var change int
			if newState == On {
				change = 1
			} else {
				change = -1
			}

			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					nextBoard.At(x+dx, y+dy).neighborCount += change
				}
			}
		}
	}

	return nextBoard
}

func (b *Board) Print() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			fmt.Printf("%v  ", b.At(x, y))
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
			if b.At(x+dx, y+dy).state == On {
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
	// Compare all but padding
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.At(x, y).state != other.At(x, y).state {
				return false
			}
			if b.At(x, y).neighborCount != other.At(x, y).neighborCount {
				return false
			}
		}
	}
	return true
}
