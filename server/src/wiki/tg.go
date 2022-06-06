package wiki

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"project/config"
	"project/pb"
	"project/zj"
)

var tgURL = `https://tanks.gg/api/list`

type tgMap map[string]*pb.TGRow

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

	defer zj.Watch(&err)

	file := config.OutputPath + `/tg.json`
	ab, err = os.ReadFile(file)
	if err == nil {
		return
	}

	rsp, err := http.Get(tgURL)
	if err != nil {
		return
	}

	ab, err = io.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	os.WriteFile(file, ab, 0666)
	return
}
