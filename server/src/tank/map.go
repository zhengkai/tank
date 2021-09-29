package tank

import (
	"project/db"
	"project/pb"
	"project/zj"
	"sync"

	"google.golang.org/protobuf/proto"
)

var mapBase = make(map[uint32]*pb.TankBase)

// Mux ...
var Mux sync.Mutex

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

	Mux.Lock()
	old, ok := mapBase[tb.ID]
	Mux.Unlock()
	if ok && proto.Equal(old, tb) {
		return
	}

	mapBase[tb.ID] = tb
	err = db.TankBase(tb)
	zj.IO(`update`, tb.ID, tb.Name, tb.Shop, err)
	return
}
