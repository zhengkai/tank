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

	defer zj.Watch(&err)

	defaultSID()

	tg, _ := getTgMap()
	s4, _ := getS4Map()

	wm, err := db.SIDWikiList()
	if err != nil {
		return
	}

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
