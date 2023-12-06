package main

import "testing"

func TestNeighborCount(t *testing.T) {
	b := Board{
		{On, Off, Off},
		{Off, On, Off},
		{Off, Off, Off},
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
