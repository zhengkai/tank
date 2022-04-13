package tank

import "sync"

// PercentPool ...
type PercentPool struct {
	m   map[uint32]*PercentRow
	mux sync.Mutex
}

// PercentRow ...
type PercentRow struct {
	P1 uint32
	P2 uint32
	P3 uint32
}

var defaultPercent = &PercentPool{
	m: make(map[uint32]*PercentRow),
}

// UpdatePercent ...
var UpdatePercent = defaultPercent.Update

// Update ...
func (pp *PercentPool) Update(id, num uint32, p int) {

	pp.mux.Lock()
	row, _ := pp.m[id]
	if row == nil {
		row = &PercentRow{}
		pp.m[id] = row
	}
	pp.mux.Unlock()

	switch p {
	case 2:
		row.P2 = num
	case 3:
		row.P3 = num
	default:
		row.P1 = num
	}
}

// Get ...
func (pp *PercentPool) Get(id uint32) (p1, p2, p3 uint32) {
	pp.mux.Lock()
	row, ok := pp.m[id]
	pp.mux.Unlock()
	if !ok {
		return
	}
	p1 = row.P1
	p2 = row.P2
	p3 = row.P3
	return
}
