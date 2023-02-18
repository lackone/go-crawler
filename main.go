package main

import (
	"fmt"
	"github.com/lackone/go-crawler/pkg/collect"
	"github.com/lackone/go-crawler/pkg/proxy"
)

func main() {
	robinProxy, _ := proxy.NewRoundRobinProxy("http://127.0.0.1:8888")

	fetch := collect.BrowserFetch{
		Proxy: robinProxy,
	}
	get, _ := fetch.Get("https://book.douban.com/subject/1007305/")
	fmt.Println(string(get))
}
