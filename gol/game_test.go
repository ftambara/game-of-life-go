package main

import "testing"

func TestNeighborCount(t *testing.T) {
	b := Board{
		states: [][]CellState{
			{On, Off, Off},
			{Off, On, Off},
			{Off, Off, Off},
		},
		width:  3,
		height: 3,
	}
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
	b := Board{
		states: [][]CellState{
			{On, Off, Off},
			{Off, On, Off},
			{Off, Off, Off},
		},
		width:  3,
		height: 3,
	}
	if !b.Equals(&b) {
		t.Errorf("Expected %v to equal %v", b, b)
	}
	b2 := Board{
		states: [][]CellState{
			{On, Off, Off},
			{Off, On, Off},
			{Off, Off, On},
		},
		width:  3,
		height: 3,
	}
	if b.Equals(&b2) {
		t.Errorf("Expected %v to not equal %v", b, b2)
	}
}

func TestCrossAdvance(t *testing.T) {
	b := Board{
		states: [][]CellState{
			{Off, On, Off},
			{On, On, On},
			{Off, On, Off},
		},
		width:  3,
		height: 3,
	}
	next := b.advance()
	expected := Board{
		states: [][]CellState{
			{On, On, On},
			{On, Off, On},
			{On, On, On},
		},
		width:  3,
		height: 3,
	}
	if !next.Equals(&expected) {
		t.Error("Board did not advance as expected")
	}
	next = next.advance()
	expected = Board{
		states: [][]CellState{
			{On, Off, On},
			{Off, Off, Off},
			{On, Off, On},
		},
		width:  3,
		height: 3,
	}
	if !next.Equals(&expected) {
		t.Error("Board did not advance as expected")
	}
}

func TestBlockAdvance(t *testing.T) {
	b := Board{
		states: [][]CellState{
			{Off, Off, Off, Off},
			{Off, On, On, Off},
			{Off, On, On, Off},
			{Off, Off, Off, Off},
		},
		width:  4,
		height: 4,
	}
	next := b.advance()
	if !next.Equals(&b) {
		t.Error("Board did not advance as expected")
	}
}

func TestBlinkerAdvance(t *testing.T) {
	b := Board{
		states: [][]CellState{
			{Off, Off, Off, Off, Off},
			{Off, Off, On, Off, Off},
			{Off, Off, On, Off, Off},
			{Off, Off, On, Off, Off},
			{Off, Off, Off, Off, Off},
		},
		width:  5,
		height: 5,
	}
	next := b.advance()
	expected := Board{
		states: [][]CellState{
			{Off, Off, Off, Off, Off},
			{Off, Off, Off, Off, Off},
			{Off, On, On, On, Off},
			{Off, Off, Off, Off, Off},
			{Off, Off, Off, Off, Off},
		},
		width:  5,
		height: 5,
	}
	if !next.Equals(&expected) {
		t.Error("Board did not advance as expected")
	}
}
