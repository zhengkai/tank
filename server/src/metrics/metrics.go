package metrics

import "time"

var (
	crawlTime  = newSummary(`crawl_time`, `爬取时间`)
	crawlFail  = newCounter(`crawl_fail`, `爬取失败`)
	crawlBytes = newCounter(`crawl_bytes`, `爬取字节数`)
)

// CrawlFail ...
func CrawlFail() {
	crawlFail.Inc()
}

// CrawlTime ...
func CrawlTime(n time.Duration) {
	crawlTime.Observe(float64(n) / float64(time.Second))
}

// CrawlBytes ...
func CrawlBytes(ab []byte) {
	crawlBytes.Add(float64(len(ab)))
}
