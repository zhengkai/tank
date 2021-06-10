package db

import (
	"project/pb"

	"google.golang.org/protobuf/proto"
)

// TankBase ...
func TankBase(tb *pb.TankBase) (err error) {

	ab, err := proto.Marshal(tb)
	if err != nil {
		return
	}

	query := `INSERT INTO tank SET id = ?, bin = ? ON DUPLICATE KEY UPDATE bin = ?`
	_, err = d.Exec(query, tb.ID, ab, ab)
	return
}

// LoadTankBase ...
func LoadTankBase() (list []*pb.TankBase, err error) {
	query := `SELECT bin FROM tank`
	r, err := d.Query(query)
	if err != nil {
		return
	}
	defer r.Close()

	for r.Next() {
		var bin []byte
		err = r.Scan(&bin)
		if err != nil {
			return
		}

		tb := &pb.TankBase{}
		err = proto.Unmarshal(bin, tb)
		if err != nil {
			return
		}
		list = append(list, tb)
	}
	return
}
