package proxy

import (
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
)

type roundRobinSwitcher struct {
	proxyURLs []*url.URL
	index     uint32
}

func (r *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RoundRobinProxySwitcher(ProxyURLs ...string) (colly.ProxyFunc, error) {
	_ = "STUB: not implemented"
	return *new(colly.ProxyFunc), nil
}
