package main

import (
	"flag"
	"github.com/golang/glog"
	"net/http"
	"os"
	"strings"
)

func main() {
	_ = flag.Set("logtostderr", "true")
	_ = flag.Set("v", "4")
	flag.Parse()
	defer glog.Flush()

	glog.Info("Starting http server...")
	mux := http.NewServeMux()

	envs := loadEnv()
	mux.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range envs {
			w.Header().Set(k, v)
		}
		code := http.StatusOK
		w.WriteHeader(code)
		glog.Infof("url:%s ip:%s code:%d", "/env", r.RemoteAddr, code)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusOK
		w.WriteHeader(code)
		glog.Infof("url:%s ip:%s code:%d", "/env", r.RemoteAddr, code)
	})

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		for k, vs := range r.Header {
			for _, v := range vs {
				w.Header().Set(k, v)
			}
		}
		code := http.StatusOK
		w.WriteHeader(code)
		glog.Infof("url:%s ip:%s code:%d", "/env", r.RemoteAddr, code)
	})

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		panic(err)
	}

}

func loadEnv() map[string]string {
	raws := os.Environ()

	envs := make(map[string]string, len(raws))

	for _, r := range raws {
		ss := strings.SplitN(r, "=", 2)
		var k, v string
		k = ss[0]
		if len(ss) > 1 {
			v = ss[1]
		}
		envs[k] = v
	}
	return envs
}
