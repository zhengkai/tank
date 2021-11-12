package spider

import (
	"fmt"
	"io/ioutil"
	"project/config"
	"project/zj"
	"time"

	"github.com/zhengkai/zu"
)

// CrawlAll ...
func CrawlAll() (err error) {
	return crawlAll(false)
}

// CrawlAllSimulate ...
func CrawlAllSimulate() (err error) {
	return crawlAll(true)
}

func crawlAll(simulate bool) (err error) {

	var cnt int
	var cntSum int

	for lv := 10; lv >= 3; lv-- {
		for _, higher := range []bool{false, true} {
			for _, ty := range []int{0, 1, 2} {
				t := time.Now()
				cnt, err = Crawl(lv, higher, ty, simulate)
				cntSum += cnt
				if cnt > 0 {
					zj.J(`crawl row`, lv, higher, ty, time.Now().Sub(t), cnt)
				}
			}
		}
	}
	zj.J(`crawl`, cntSum)
	return
}

// Crawl ...
func Crawl(tier int, higher bool, ty int, simulate bool) (cnt int, err error) {

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

		file := fmt.Sprintf(`%s/%d-%d-%d-%s.json`, config.TmpPath, tier, page, ty, t)

		var url string

		if simulate {
			ab, err = ioutil.ReadFile(file)
			if err != nil {
				return
			}
		} else {
			time.Sleep(time.Second)
			url = `https://tbox.wot.360.cn/rank/more?rank_type=%s&page=%d&size=30&type=%s&tier=%d&sort=damage_dealt_avg&tank_sort=1,2,3`
			url = fmt.Sprintf(url, t, page, sty, tier)
			zj.J(url)

			for range []int{0, 0, 0} {
				ab, err = zu.FetchURL(url)
				if err != nil {
					break
				}
			}
			if err != nil {
				zj.W(`fetch url fail`, err)
				return
			}
			ioutil.WriteFile(file, ab, 0666)
		}

		cnt++
		next, err = Parse(ab)
		if page > 1 && err == errEmptyList {
			next = false
			err = nil
		}
		if err != nil {
			zj.W(`prase`, url, err)
			break
		}

		zj.J(`url`, url)

		if !next {
			break
		}
	}
	return
}
