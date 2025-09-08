package tank

import (
	"encoding/json"
	"fmt"
	"os"
	"project/config"
	"project/db"
	"project/pb"
	"project/util"
	"project/zj"
	"strconv"
)

// Build ...
func Build() (err error) {

	defer zj.Watch(&err)

	tl := &pb.TankList{}

	var dateMin int
	var dateMax int
	Mux.Lock()
	for _, tb := range mapBase {
		t := &pb.Tank{
			Base: tb,
		}
		var date int
		t.Stats, t.StatsHigher, date, err = db.LoadTankStats(tb.ID)
		if err != nil {
			zj.J(`skip`, tb.ID)
			continue
		}
		if t.Stats == nil && t.StatsHigher == nil {
			continue
		}

		p1, p2, p3 := defaultPercent.Get(tb.ID)
		for _, s := range []*pb.TankStats{t.Stats, t.StatsHigher} {
			if s != nil {
				s.P1 = p1
				s.P2 = p2
				s.P3 = p3
			}
		}

		tl.List = append(tl.List, t)

		if dateMax < date {
			dateMax = date
		}
		if dateMin > date || dateMin == 0 {
			dateMin = date
		}
	}
	Mux.Unlock()

	if len(tl.List) < config.AssertCount {
		err = fmt.Errorf(`tank list count: %d`, len(tl.List))
		return
	}

	sdx := strconv.Itoa(dateMax)
	sdx = fmt.Sprintf(`%s.%s.%s`, sdx[:4], sdx[4:6], sdx[6:])
	tl.BuildTime = sdx
	if dateMax != dateMin {
		sdi := strconv.Itoa(dateMin)
		sdi = fmt.Sprintf(`%s.%s.%s`, sdi[:4], sdi[4:6], sdi[6:])
		tl.BuildTime = sdi + `-` + sdx
	}

	util.WriteFile(File(), tl)

	ab, err := json.MarshalIndent(tl, ``, "\t")
	if err == nil {
		os.WriteFile(config.OutputPath+`/list.json`, ab, 0666)
	}

	return
}

// File ...
func File() string {
	return config.OutputPath + `/list.pb`
}
