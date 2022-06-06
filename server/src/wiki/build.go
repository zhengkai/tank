package wiki

import (
	"project/config"
	"project/db"
	"project/pb"
	"project/util"
	"project/zj"
)

func build() (err error) {

	defer zj.Watch(&err)

	defaultSID()

	m, err := getTgMap()
	if err != nil {
		return
	}

	wm, err := db.SIDWikiList()
	if err != nil {
		return
	}

	for _, v := range wm {

		sid := v.Wiki

		t, ok := m[sid]
		if !ok {
			zj.W(`skip`, sid, t)
			continue
		}

		if v.Tanksgg != t.Slug {
			db.SIDTg(v.ID, t.Slug)
			v.Tanksgg = t.Slug
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
	return
}
