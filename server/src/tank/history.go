package tank

import (
	"fmt"
	"project/db"
	"project/pb"
	"project/zj"
	"time"
)

// BuildHistory ...
func BuildHistory() (err error) {
	Mux.Lock()
	defer Mux.Unlock()
	size := len(mapBase)
	i := 0
	for _, tb := range mapBase {
		i++
		zj.F(`%3d/%3d %5d %s`, i, size, tb.ID, tb.Name)
		historyOne(tb.ID)
		time.Sleep(time.Second / 20)
	}
	return
}

func historyOne(id uint32) (err error) {

	defer zj.Watch(&err)

	m := db.GetPercentRange(id)

	li, err := db.LoadTankRecentStats(id)
	for _, v := range li {

		p, ok := m[v.Date]
		if !ok {
			continue
		}

		s := v.Stats
		if s != nil {
			s.P1 = p.P1
			s.P2 = p.P2
			s.P3 = p.P3
		}
		s = v.StatsHigher
		if s != nil {
			s.P1 = p.P1
			s.P2 = p.P2
			s.P3 = p.P3
		}

	}

	d := &pb.TankStatHistory{
		TankId: id,
		List:   li,
	}

	file := fmt.Sprintf(`history/%d.pb`, id)
	WritePB(file, d)
	return
}
