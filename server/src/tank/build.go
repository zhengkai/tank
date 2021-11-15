package tank

import (
	"fmt"
	"io/ioutil"
	"project/config"
	"project/db"
	"project/pb"
	"project/zj"
	"strconv"

	"github.com/zhengkai/zu"
	"google.golang.org/protobuf/proto"
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

	ab, err := proto.Marshal(tl)
	if err != nil {
		return
	}

	file := File()

	err = ioutil.WriteFile(file, ab, 0666)
	if err != nil {
		return
	}

	err = zu.Brotli(file)
	if err != nil {
		return
	}

	err = zu.Gzip(file)
	if err != nil {
		return
	}

	zj.J(`build`, file, len(ab), len(tl.List))

	return
}

// File ...
func File() string {
	return config.OutputPath + `/list.pb`
}
