package colly

import (
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/gobwas/glob"
)

type httpBackend struct {
	LimitRules []*LimitRule
	Client     *http.Client
	lock       *sync.RWMutex
}

type checkResponseHeadersFunc func(req *http.Request, statusCode int, header http.Header) bool
type checkRequestHeadersFunc func(req *http.Request) bool

type LimitRule struct {
	DomainRegexp string

	DomainGlob string

	Delay time.Duration

	RandomDelay time.Duration

	Parallelism    int
	waitChan       chan bool
	compiledRegexp *regexp.Regexp
	compiledGlob   glob.Glob
	lock           sync.Mutex
}

func (r *LimitRule) Init() error { _ = "STUB: not implemented"; return nil }

func (r *LimitRule) Clone() *LimitRule { _ = "STUB: not implemented"; return nil }

func (h *httpBackend) Init(jar http.CookieJar) { _ = "STUB: not implemented"; return }

func (r *LimitRule) Match(domain string) bool { _ = "STUB: not implemented"; return false }

func (h *httpBackend) GetMatchingRule(domain string) *LimitRule {
	_ = "STUB: not implemented"
	return nil
}

func (h *httpBackend) Cache(request *http.Request, bodySize int, checkRequestHeadersFunc checkRequestHeadersFunc, checkResponseHeadersFunc checkResponseHeadersFunc, cacheDir string, cacheExpiration time.Duration) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpBackend) Do(request *http.Request, bodySize int, checkRequestHeadersFunc checkRequestHeadersFunc, checkResponseHeadersFunc checkResponseHeadersFunc) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *httpBackend) Limit(rule *LimitRule) error { _ = "STUB: not implemented"; return nil }

func (h *httpBackend) Limits(rules []*LimitRule) error { _ = "STUB: not implemented"; return nil }
