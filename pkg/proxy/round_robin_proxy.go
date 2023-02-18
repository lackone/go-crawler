package proxy

import (
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

type RoundRobinProxy struct {
	proxyUrls []*url.URL
	index     uint32
}

func NewRoundRobinProxy(proxyUrls ...string) (ProxyFunc, error) {
	if len(proxyUrls) < 1 {
		return nil, errors.New("proxyUrls is empty")
	}
	urls := make([]*url.URL, len(proxyUrls))
	for k, v := range proxyUrls {
		parse, err := url.Parse(v)
		if err != nil {
			return nil, err
		}
		urls[k] = parse
	}
	return (&RoundRobinProxy{proxyUrls: urls, index: 0}).GetProxy, nil
}

func (r *RoundRobinProxy) GetProxy(req *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&r.index, 1) - 1
	u := r.proxyUrls[index%uint32(len(r.proxyUrls))]
	return u, nil
}
