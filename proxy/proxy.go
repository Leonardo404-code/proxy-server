package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	fmt.Printf(
		"[reverse proxy server] received request at: %s\n",
		time.Now().Format(time.DateTime),
	)

	r.Host = originServerURL.Host
	r.URL.Host = originServerURL.Host
	r.URL.Scheme = originServerURL.Scheme
	r.RequestURI = ""

	time.Sleep(time.Second * 3)

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Copy(w, res.Body)
}
