package db

import (
	"fmt"
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
