package wiki

import (
	"fmt"
	"os"
	"project/db"
	"project/pb"
	"project/tank"
	"project/zj"

	"google.golang.org/protobuf/proto"
)

// Test ...
func Test() {

	m, err := getTgMap()

	list, _ := db.SIDWikiList()

	for _, v := range list {

		sid := v.Wiki

		t, ok := m[sid]
		if !ok {
			zj.W(`skip`, sid, t)
			continue
		}
		fmt.Printf("%d: \"%s\",\n", v.ID, t.Slug)
		// zj.Raw(fmt.Sprintf("%d: \"%s\",\n", v.ID, t.Slug))
	}

	zj.J(len(m))
	zj.J(len(list))

	return

	ab, err := os.ReadFile(tank.File())
	if err != nil {
		return
	}

	tl := &pb.TankList{}
	err = proto.Unmarshal(ab, tl)
	if err != nil {
		return
	}

	for _, v := range tl.GetList() {
		id := v.GetBase().GetID()
		if id < 1 {
			continue
		}
		sid, err := SID(id)
		zj.J(id, sid, err)
	}
}
