package spider

import (
	"fmt"
	"os"
	"project/config"
	"project/metrics"
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

	for lv := 11; lv >= 3; lv-- {
		for _, higher := range []bool{false, true} {
			for _, ty := range []int{0, 1, 2} {
				t := time.Now()
				cnt, err = Crawl(lv, higher, ty, simulate)
				cntSum += cnt
				if cnt > 0 {
					zj.J(`crawl row`, lv, higher, ty, time.Since(t), cnt)
				}
			}
		}

		for _, percent := range []int{1, 2, 3} {
			t := time.Now()
			cnt, err = CrawlPercent(lv, percent, simulate)
			cntSum += cnt
			if cnt > 0 {
				zj.J(`crawl percent`, lv, percent, time.Since(t), cnt)
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

		url := `https://tbox.wot.360.cn/rank/more?rank_type=%s&page=%d&size=30&type=%s&tier=%d&sort=damage_dealt_avg&tank_sort=1,2,3`

		if simulate {
			ab, err = os.ReadFile(file)
			if err != nil {
				return
			}
		} else {
			time.Sleep(time.Second)
			url = fmt.Sprintf(url, t, page, sty, tier)
			zj.J(url)

			for range []int{0, 0, 0} {
				t := time.Now()
				ab, err = zu.FetchURL(url)
				metrics.CrawlTime(time.Now().Sub(t))
				metrics.CrawlBytes(ab)
				if err == nil {
					break
				}
				metrics.CrawlFail()
			}
			if err != nil {
				zj.W(`fetch url fail`, err)
				return
			}
			os.WriteFile(file, ab, 0666)
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

		if !next {
			break
		}
	}
	return
}

// CrawlPercent ...
func CrawlPercent(tier, percent int, simulate bool) (cnt int, err error) {

	percentNum := 65
	if percent == 2 {
		percentNum = 85
	} else if percent == 3 {
		percentNum = 95
	}

	var ab []byte
	var next bool
	var page int
	for {
		page++

		var url string
		file := fmt.Sprintf(`%s/percent-%d-%d-%d.json`, config.TmpPath, tier, percent, page)

		if simulate {
			ab, err = os.ReadFile(file)
			if err != nil {
				return
			}
		} else {
			time.Sleep(time.Second)

			url = `https://tbox.wot.360.cn/rank/more?percentile=%d&rank_type=default&page=%d&size=30&type=&tier=%d&sort=mastery&tank_sort=&nation=`
			url = fmt.Sprintf(url, percentNum, page, tier)

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
			os.WriteFile(file, ab, 0666)
		}

		cnt++
		next, err = ParsePercent(ab, percent)
		if page > 1 && err == errEmptyList {
			next = false
			err = nil
		}
		if err != nil {
			zj.W(`prase`, url, err)
			break
		}

		if !next {
			break
		}
	}
	return
}
