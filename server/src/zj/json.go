package zj

import "encoding/json"

// JSON ...
func JSON(v interface{}) (s string) {

	ab, err := json.Marshal(v)
	if err != nil {
		return
	}

	return string(ab)
}
