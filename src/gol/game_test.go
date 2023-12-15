package main

import (
	"fmt"
	"testing"
)

func newTestBoard(states [][]CellState) *Board {
	return NewBoard(len(states[0]), len(states), func(x, y int) CellState {
		return states[y][x]
	})
}

func assertBoard(t *testing.T, expected, actual *Board) {
	if !expected.Equals(actual) {
		fmt.Println("Expected:")
		expected.Print()
		fmt.Println("Got:")
		actual.Print()
		t.Error("Boards did not match")
	}
}

func TestNeighborCount(t *testing.T) {
	b := newTestBoard(
		[][]CellState{
			{On, Off, Off},
			{Off, On, Off},
			{Off, Off, Off},
		},
	)
	count := b.CountNeighbors(0, 0)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
	count = b.CountNeighbors(1, 1)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
	count = b.CountNeighbors(1, 2)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestEquals(t *testing.T) {
	b := newTestBoard(
		[][]CellState{
			{On, Off, Off},
			{Off, On, Off},
			{Off, Off, Off},
		},
	)
	if !b.Equals(b) {
		t.Errorf("Expected %v to equal %v", b, b)
	}
	b2 := newTestBoard(
		[][]CellState{
			{On, Off, Off},
			{Off, On, Off},
			{Off, Off, On},
		},
	)
	if b.Equals(b2) {
		t.Errorf("Expected %v to not equal %v", b, b2)
	}
}

func TestCrossAdvance(t *testing.T) {
	b := newTestBoard(
		[][]CellState{
			{Off, On, Off},
			{On, On, On},
			{Off, On, Off},
		},
	)
	next := b.Advance()
	expected := newTestBoard(
		[][]CellState{
			{On, On, On},
			{On, Off, On},
			{On, On, On},
		},
	)
	assertBoard(t, expected, next)
	next = next.Advance()
	expected = newTestBoard(
		[][]CellState{
			{On, Off, On},
			{Off, Off, Off},
			{On, Off, On},
		},
	)
	assertBoard(t, expected, next)
}

func TestBlockAdvance(t *testing.T) {
	b := newTestBoard(
		[][]CellState{
			{Off, Off, Off, Off},
			{Off, On, On, Off},
			{Off, On, On, Off},
			{Off, Off, Off, Off},
		},
	)
	next := b.Advance()
	assertBoard(t, b, next)
}

func TestBlinkerAdvance(t *testing.T) {
	b := newTestBoard(
		[][]CellState{
			{Off, Off, Off, Off, Off},
			{Off, Off, On, Off, Off},
			{Off, Off, On, Off, Off},
			{Off, Off, On, Off, Off},
			{Off, Off, Off, Off, Off},
		},
	)
	next := b.Advance()
	expected := newTestBoard(
		[][]CellState{
			{Off, Off, Off, Off, Off},
			{Off, Off, Off, Off, Off},
			{Off, On, On, On, Off},
			{Off, Off, Off, Off, Off},
			{Off, Off, Off, Off, Off},
		},
	)
	assertBoard(t, expected, next)
}
