package main

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

func (c *Cell) Next() CellState {
	if c.state == On {
		if c.neighborCount == 2 || c.neighborCount == 3 {
			return On
		} else {
			return Off
		}
	} else {
		if c.neighborCount == 3 {
			return On
		} else {
			return Off
		}
	}
}

type Cell struct {
	state         CellState
	neighborCount int
}

func (c Cell) String() string {
	return c.state.String()
}
