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

type Cell struct {
	state         CellState
	neighborCount int
}

func (c Cell) String() string {
	return c.state.String()
}
