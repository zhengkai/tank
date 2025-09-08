package wiki

import (
	"encoding/json"
	"io"
	"os"
	"project/config"
	"project/pb"
	"project/util"
	"project/zj"
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

	file := config.OutputPath + `/tg.json`

	defer zj.Watch(&err)

	rsp, err := util.HTTPNoProxyGet(tgURL)
	if err == nil {
		ab, err = io.ReadAll(rsp.Body)
		if err == nil {
			os.WriteFile(file, ab, 0666)
			return
		}
	}
	return os.ReadFile(file)
}
