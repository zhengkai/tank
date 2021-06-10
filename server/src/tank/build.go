package tank

import (
	"fmt"
	"project/config"
	"project/db"
	"project/pb"
	"project/zj"

	"google.golang.org/protobuf/proto"
)

// Build ...
func Build() (err error) {

	defer zj.Watch(&err)

	tl := &pb.TankList{}

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

	if len(tl.List) < config.AssertCount {
		err = fmt.Errorf(`tank list count: %d`, len(tl.List))
		return
	}

	ab, err := proto.Marshal(tl)
	if err != nil {
		return
	}

	zj.J(`build`, len(ab), len(tl.List))

	return
}
