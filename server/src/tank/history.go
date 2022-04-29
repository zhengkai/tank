package tank

import (
	"fmt"
	"project/config"
	"project/db"
	"project/pb"
	"project/zj"
	"time"

	"google.golang.org/protobuf/proto"
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

func historyOne(id uint32) {

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

	ab, err := proto.Marshal(d)
	if err != nil {
		return
	}

	file := fmt.Sprintf(`%s/history/%d.pb`, config.OutputPath, id)
	buildWriteFile(file, ab)
	zj.J(file)
}
