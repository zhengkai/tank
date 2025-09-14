package tank

import (
	"fmt"
	"project/config"
	"project/util"
	"project/zj"

	"google.golang.org/protobuf/proto"
)

const dirPrefix = `data`

func fileLocal(file string) string {
	return fmt.Sprintf(`%s/%s/%s`, config.StaticDir, dirPrefix, file)
}

func WriteFile(file string, ab []byte) error {
	file = fmt.Sprintf(`%s/%s`, dirPrefix, file)
	return util.WriteFile(file, ab)
}

func WritePB(file string, d proto.Message) (err error) {

	defer zj.Watch(&err)
	ab, err := proto.Marshal(d)
	if err != nil {
		return
	}
	return WriteFile(file, ab)
}
