package tank

import (
	"fmt"
	"project/db"
	"project/pb"
	"project/zj"
	"strconv"
	"time"
)

var expiredDate uint32

// BuildHistory ...
func BuildHistory() (err error) {
	Mux.Lock()
	defer Mux.Unlock()

	i, _ := strconv.Atoi(time.Now().Add(-30 * 24 * time.Hour).Format(`20060102`))
	expiredDate = uint32(i)

	var removeID []uint32
	for _, tb := range mapBase {
		remove := historyOne(tb.ID)
		if remove {
			removeID = append(removeID, tb.ID)
			// zj.J(tb.ID, tb.Name)
		}
	}
	for _, id := range removeID {
		delete(mapBase, id)
	}

	return
}

func historyOne(id uint32) (remove bool) {

	m := db.GetPercentRange(id)
	li, err := db.LoadTankRecentStats(id)
	if err != nil {
		zj.F(`db.LoadTankRecentStats %d fail: %s`, id, err.Error())
		return false
	}

	// 超过 30 天没数据的 id 会被永久隐藏
	size := len(li)
	if size > 0 {
		i := li[size-1].Date
		if i < expiredDate {
			return true
		}
	}

	historyOneList(m, li)

	d := &pb.TankStatHistory{
		TankId: id,
		List:   li,
	}

	file := fmt.Sprintf(`history/%d.pb`, id)
	WritePB(file, d)
	return false
}

func historyOneList(m map[uint32]*pb.TankStats, list []*pb.TankStatDate) {

	for _, v := range list {

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
}
