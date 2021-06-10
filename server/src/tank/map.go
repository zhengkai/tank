package tank

import (
	"project/db"
	"project/pb"
	"project/zj"

	"google.golang.org/protobuf/proto"
)

var mapBase = make(map[uint32]*pb.TankBase)

// InitMap ...
func InitMap() {
	list, err := db.LoadTankBase()
	if err != nil {
		return
	}

	for _, v := range list {
		mapBase[v.ID] = v
	}
}

func poolUpdate(tb *pb.TankBase) (err error) {

	old, ok := mapBase[tb.ID]
	if ok && proto.Equal(old, tb) {
		return
	}

	mapBase[tb.ID] = tb
	err = db.TankBase(tb)
	zj.J(`update`, tb.ID, tb.Name, err)
	return
}
