package tank

import (
	"fmt"
	"project/config"
	"project/db"
	"project/pb"
	"project/util"
	"project/zj"
	"time"
)

// BuildHistory ...
func BuildHistory() (err error) {
	Mux.Lock()
	defer Mux.Unlock()
	for _, tb := range mapBase {
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

	file := fmt.Sprintf(`%s/history/%d.pb`, config.OutputPath, id)
	util.WriteFile(file, d)
	return
}
