package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

// 正向代理
func TestForwardProxy(t *testing.T) {
	s := &http.Server{
		Addr: ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleHttp(w, r)
		}),
	}
	err := s.ListenAndServe()
	log.Fatalln(err)
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	req := new(http.Request)
	*req = *r

	fmt.Println(req)

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for k, v := range res.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}
