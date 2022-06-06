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
		if v.ID == 60529 {
			zj.J(v)
		}
		mapBase[v.ID] = v
		initPercentPool(v.ID)
	}
}

// LoadID ...
func LoadID() (li []uint32) {

	Mux.Lock()
	defer Mux.Unlock()

	li = make([]uint32, 0, len(mapBase))
	for i := range mapBase {
		li = append(li, i)
	}
	return
}

func initPercentPool(id uint32) {
	p1, p2, p3, err := db.GetPercent(id)
	if err != nil {
		return
	}
	UpdatePercent(id, p1, 1)
	UpdatePercent(id, p2, 2)
	UpdatePercent(id, p3, 3)
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
