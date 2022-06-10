package wiki

import (
	"encoding/json"
	"io/ioutil"
	"project/config"
	"project/db"
	"project/pb"
	"project/util"
	"project/zj"
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

	file := config.OutputPath + `/id.pb`
	util.WriteFile(file, out)

	ab, err := json.MarshalIndent(out, ``, "\t")
	if err == nil {
		ioutil.WriteFile(config.OutputPath+`/id.json`, ab, 0666)
	}

	return
}
