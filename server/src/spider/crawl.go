package spider

import (
	"fmt"
	"io/ioutil"

	"github.com/zhengkai/zu"
)

// CrawlAll ...
func CrawlAll(simulate bool) (err error) {

	for lv := 1; lv <= 10; lv++ {
		for _, higher := range []bool{false, true} {
			for _, ty := range []int{0, 1, 2} {
				err = Crawl(lv, higher, ty, simulate)
				if err != nil {
					break
				}
			}
		}
	}
	return
}

// Crawl ...
func Crawl(tier int, higher bool, ty int, simulate bool) (err error) {

	t := `default`
	if higher {
		t = `higher`
	}
	var ab []byte
	var next bool
	var page int
	sty := `heavyTank`

	switch ty {
	case 1:
		sty = `SPG,mediumTank`
	case 2:
		sty = `lightTank,AT-SPG`
	}

	for {
		page++
		if page > 10 {
			break
		}

		file := fmt.Sprintf(`/www/tank/tmp/%d-%d-%d-%s.json`, tier, page, ty, t)

		if simulate {
			ab, err = ioutil.ReadFile(file)
			if err != nil {
				return
			}
		} else {
			url := `https://tbox.wot.360.cn/rank/more?rank_type=%s&page=%d&size=30&type=%s&tier=%d&sort=damage_dealt_avg&&tank_sort=1,2,3`
			url = fmt.Sprintf(url, t, page, sty, tier)

			ab, err = zu.FetchURL(url)
			if err != nil {
				return
			}
			ioutil.WriteFile(file, ab, 0666)
		}

		next, err = Parse(ab)
		if err != nil {
			return
		}

		if !next {
			break
		}
	}
	return
}
