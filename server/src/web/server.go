package web

import (
	"fmt"
	"net/http"
	"project/task"
	"project/zj"
	"time"

	"github.com/arl/statsviz"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server ...
func Server(port int) {

	addr := fmt.Sprintf(`:%d`, port)

	mux := http.NewServeMux()
	mux.HandleFunc(`/task/crawl`, taskCrawlHandle)
	mux.HandleFunc(`/task/build`, taskBuildHandle)
	mux.Handle(`/metrics`, promhttp.Handler())
	mux.HandleFunc(`/`, failbackHandle)

	statsviz.Register(mux)

	s := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	zj.J(`start web server`, addr)

	s.ListenAndServe()
}

func taskCrawlHandle(w http.ResponseWriter, r *http.Request) {
	go task.Crawl()
	w.Write([]byte(`task crawl`))
}

func taskBuildHandle(w http.ResponseWriter, r *http.Request) {
	go task.Build()
	w.Write([]byte(`task build`))
}

func failbackHandle(w http.ResponseWriter, r *http.Request) {
	zj.J(`failback handle`, r.URL.String())
}
