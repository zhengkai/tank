package wiki

import (
	"encoding/json"
	"project/pb"
	"project/util"
	"project/zj"

	"github.com/zhengkai/zu"
)

var tgURL = `https://tanks.gg/api/list`

func getTgMap() (m map[string]*pb.TGRow, err error) {

	m = make(map[string]*pb.TGRow)

	d := &pb.TGList{}

	ab, err := tgFile()
	if err != nil {
		return
	}

	err = json.Unmarshal(ab, d)
	if err != nil {
		return
	}

	for _, v := range d.GetTanks() {
		m[v.Id] = v
	}

	return
}

func tgFile() (ab []byte, err error) {

	file := `data/tg.json`

	defer zj.Watch(&err)

	ab, err = zu.FetchURL(tgURL)
	if err == nil {
		util.WriteFile(file, ab)
		return
	}
	return util.ReadFile(file)
}
