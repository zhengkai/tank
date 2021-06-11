package tank

import (
	"fmt"
	"io/ioutil"
	"project/config"
	"project/db"
	"project/pb"
	"project/zj"

	"github.com/zhengkai/zu"
	"google.golang.org/protobuf/proto"
)

// Build ...
func Build() (err error) {

	defer zj.Watch(&err)

	tl := &pb.TankList{}

	Mux.Lock()
	for _, tb := range mapBase {
		t := &pb.Tank{
			Base: tb,
		}

		t.Stats, t.StatsHigher, err = db.LoadTankStats(tb.ID)
		if err != nil {
			zj.J(`skip`, tb.ID)
			continue
		}

		tl.List = append(tl.List, t)
	}
	Mux.Unlock()

	if len(tl.List) < config.AssertCount {
		err = fmt.Errorf(`tank list count: %d`, len(tl.List))
		return
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
	return config.FilePath + `/list.pb`
}
