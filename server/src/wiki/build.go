package wiki

import (
	"encoding/json"
	"project/db"
	"project/pb"
	"project/util"
	"project/zj"

	"google.golang.org/protobuf/proto"
)

func build() (err error) {

	zj.J(`wiki build`)

	defer zj.Watch(&err)

	zj.J(`  wiki build load`)

	defaultSID()

	zj.J(`  wiki tg`)
	tg, _ := getTgMap()
	zj.J(`  wiki s4`)
	s4, _ := getS4Map()

	zj.J(`  wiki list`)
	wm, err := db.SIDWikiList()
	if err != nil {
		return
	}

	zj.J(`  wiki build loop`)
	for _, v := range wm {

		t, ok := tg[v.Wiki]
		if ok {
			if v.Tanksgg != t.Slug {
				db.SIDTg(v.ID, t.Slug)
				v.Tanksgg = t.Slug
			}
		}

		s, ok := s4[v.ID]
		if ok {
			v.Skill4Ltu = s
		}
	}

	zj.J(`wiki build write`)

	out := &pb.TankAliasList{
		List: make([]*pb.TankAlias, 0, len(wm)),
	}

	for _, v := range wm {
		out.List = append(out.List, v)
	}

	ab, err := proto.Marshal(out)
	if err != nil {
		zj.W(`proto marshal error:`, err)
	}
	file := `data/id.pb`
	util.WriteFile(file, ab)

	ab, err = json.MarshalIndent(out, ``, "\t")
	if err == nil {
		util.WriteFile(`data/id.json`, ab)
	}

	return
}
