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

func TestEquals(t *testing.T) {
	b := Board{
		{On, Off, Off},
		{Off, On, Off},
		{Off, Off, Off},
	}
	if !b.Equals(&b) {
		t.Errorf("Expected %v to equal %v", b, b)
	}
	b2 := Board{
		{On, Off, Off},
		{Off, On, Off},
		{Off, Off, On},
	}
	if b.Equals(&b2) {
		t.Errorf("Expected %v to not equal %v", b, b2)
	}
}

func TestCrossAdvance(t *testing.T) {
	b := Board{
		{Off, On, Off},
		{On, On, On},
		{Off, On, Off},
	}
	next := b.advance()
	expected := Board{
		{On, On, On},
		{On, Off, On},
		{On, On, On},
	}
	if !next.Equals(&expected) {
		t.Error("Board did not advance as expected")
	}
	next = next.advance()
	expected = Board{
		{On, Off, On},
		{Off, Off, Off},
		{On, Off, On},
	}
	if !next.Equals(&expected) {
		t.Error("Board did not advance as expected")
	}
}

func TestBlockAdvance(t *testing.T) {
	b := Board{
		{Off, Off, Off, Off},
		{Off, On, On, Off},
		{Off, On, On, Off},
		{Off, Off, Off, Off},
	}
	next := b.advance()
	if !next.Equals(&b) {
		t.Error("Board did not advance as expected")
	}
}

func TestBlinkerAdvance(t *testing.T) {
	b := Board{
		{Off, Off, Off, Off, Off},
		{Off, Off, On, Off, Off},
		{Off, Off, On, Off, Off},
		{Off, Off, On, Off, Off},
		{Off, Off, Off, Off, Off},
	}
	next := b.advance()
	expected := Board{
		{Off, Off, Off, Off, Off},
		{Off, Off, Off, Off, Off},
		{Off, On, On, On, Off},
		{Off, Off, Off, Off, Off},
		{Off, Off, Off, Off, Off},
	}
	if !next.Equals(&expected) {
		t.Error("Board did not advance as expected")
	}
}
