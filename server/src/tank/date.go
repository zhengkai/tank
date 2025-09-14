package tank

import (
	"fmt"
	"os"
	"project/db"
	"project/pb"
	"slices"
	"strconv"
	"strings"
)

var br = []byte{'\n'}

type dateRow struct {
	date  uint32
	count uint64
}

type dateData struct {
	dateMap map[uint32]bool
	tankMap map[uint32]map[uint32]uint64
	tankTop map[uint32]uint64
	baseMap map[uint32]*pb.TankBase
	higher  bool
}

// Date ...
func Date() {

	for _, isHigher := range []bool{false, true} {

		list, err := db.LoadTankBase()
		if err != nil {
			return
		}

		o := dateData{
			dateMap: make(map[uint32]bool),
			tankMap: make(map[uint32]map[uint32]uint64),
			tankTop: make(map[uint32]uint64),
			baseMap: make(map[uint32]*pb.TankBase),
			higher:  isHigher,
		}

		for _, v := range list {
			o.baseMap[v.ID] = v
			o.dateByTank(v.ID)
		}

		o.arrange()
	}

	// zj.J(o.tankTop)
}

func (dd *dateData) dateByTank(id uint32) (re []*dateRow) {
	var date uint32
	var battle uint64
	m := make(map[uint32]uint64)
	for {
		list, err := db.LoadTankDateStats(id, dd.higher, date)
		if err != nil {
			break
		}
		for _, v := range list {
			date = v.Date
			if date < 20211005 {
				continue
			}
			dd.dateMap[date] = true
			if v.Stat == nil {
				continue
			}
			count := v.Stat.Battles
			m[date] = count
			battle += count
		}
		if len(list) != 500 {
			break
		}
	}
	if battle > 0 {
		dd.tankMap[id] = m
		dd.tankTop[id] = battle
	}
	return
}

func (dd *dateData) arrange() {
	d := make([]uint32, 0, len(dd.dateMap))
	for k := range dd.dateMap {
		d = append(d, k)
	}
	slices.Sort(d)

	f, err := dd.file(d)
	if err != nil {
		return
	}

	for id, tm := range dd.tankMap {
		base := dd.baseMap[id]

		flag := fmt.Sprintf(`https://tank.9farm.com/assets/flag/%d.png`, base.Nation)
		name := strings.ReplaceAll(base.Name, `"`, ``)
		name = strings.ReplaceAll(name, `,`, ``)
		fmt.Fprintf(f, `%s,%s,%s`, name, base.Type, flag)

		for i, date := range d {
			count, ok := tm[date]
			if !ok && i > 0 {
				prev, ok := tm[d[i-1]]
				if ok {
					tm[date] = prev
					count = prev
				}
			}

			fmt.Fprintf(f, `,%d`, count)
		}
		f.Write(br)
	}
}

func (dd *dateData) file(date []uint32) (f *os.File, err error) {

	name := `tank`
	if dd.higher {
		name += `-higher`
	}

	f, err = os.OpenFile(fileLocal(name), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	f.WriteString(`Tank Name,Region,Icon`)

	for _, v := range date {
		s := strconv.Itoa(int(v))
		fmt.Fprintf(f, `,%s-%s-%s`, s[:4], s[4:6], s[6:])
	}
	f.Write(br)

	return
}
