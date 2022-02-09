package db

import "sort"

// Date ...
func Date() (list []int, err error) {

	query := `SELECT date FROM stat GROUP BY date`
	row, err := d.Query(query)
	if err != nil {
		return
	}
	defer row.Close()

	list = make([]int, 0, 300)
	var date int
	for row.Next() {
		err = row.Scan(&date)
		if err != nil {
			break
		}
		list = append(list, date)
	}

	sort.Ints(list)

	return
}

// ByDate ...
func ByDate() (list []int, err error) {
	return
}
