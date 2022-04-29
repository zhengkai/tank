package db

import (
	"fmt"
	"project/pb"
	"project/zj"
)

// UpdatePercent ...
func UpdatePercent(id uint32, date, percent, percentNum int) (err error) {

	defer zj.Watch(&err)

	col := fmt.Sprintf(`p%d`, percent)
	query := `INSERT INTO percent SET tank_id = ?, date = ?, %s = ? ON DUPLICATE KEY UPDATE %s = ?`
	query = fmt.Sprintf(query, col, col)

	_, err = d.Exec(query, id, date, percentNum, percentNum)
	return
}

// GetPercent ...
func GetPercent(id uint32) (p1, p2, p3 uint32, err error) {

	query := `SELECT p1, p2, p3 FROM percent WHERE tank_id = ? ORDER BY date DESC LIMIT 1`

	row := d.QueryRow(query, id)
	err = row.Scan(&p1, &p2, &p3)

	return
}

// GetPercentRange ...
func GetPercentRange(id uint32) (m map[uint32]*pb.TankStats) {

	query := `SELECT date, p1, p2, p3 FROM percent WHERE tank_id = ? ORDER BY date DESC LIMIT 101`

	row, err := d.Query(query, id)
	if err != nil {
		return
	}
	defer row.Close()

	m = make(map[uint32]*pb.TankStats)

	for row.Next() {

		var d uint32
		st := &pb.TankStats{}

		err = row.Scan(&d, &st.P1, &st.P2, &st.P3)
		if err != nil {
			return
		}

		m[d] = st
	}
	return
}
