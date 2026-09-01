package colly

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocolly/colly/v2/debug"
	"github.com/gocolly/colly/v2/storage"
	whatwgUrl "github.com/nlnwa/whatwg-url/url"
	"github.com/temoto/robotstxt"
)

type CollectorOption func(*Collector)

type Collector struct {
	UserAgent string

	Headers *http.Header

	MaxDepth int

	AllowedDomains []string

	DisallowedDomains []string

	DisallowedURLFilters []*regexp.Regexp

	URLFilters []*regexp.Regexp

	AllowURLRevisit bool

	MaxBodySize int

	CacheDir string

	IgnoreRobotsTxt bool

	Async bool

	ParseHTTPErrorResponse bool

	ID uint32

	DetectCharset bool

	redirectHandler func(req *http.Request, via []*http.Request) error

	CheckHead bool

	TraceHTTP bool

	Context context.Context

	MaxRequests uint32

	store                    storage.Storage
	debugger                 debug.Debugger
	robotsMap                map[string]*robotstxt.RobotsData
	htmlCallbacks            []*htmlCallbackContainer
	xmlCallbacks             []*xmlCallbackContainer
	requestCallbacks         []RequestCallback
	responseCallbacks        []ResponseCallback
	responseHeadersCallbacks []ResponseHeadersCallback
	requestHeadersCallbacks  []RequestCallback
	errorCallbacks           []ErrorCallback
	scrapedCallbacks         []ScrapedCallback
	requestCount             atomic.Uint32
	responseCount            atomic.Uint32
	backend                  *httpBackend
	wg                       *sync.WaitGroup
	lock                     *sync.RWMutex

	CacheExpiration time.Duration
}

type RequestCallback func(*Request)

type ResponseHeadersCallback func(*Response)

type ResponseCallback func(*Response)

type HTMLCallback func(*HTMLElement)

type XMLCallback func(*XMLElement)

type ErrorCallback func(*Response, error)

type ScrapedCallback func(*Response)

type ProxyFunc func(*http.Request) (*url.URL, error)

type AlreadyVisitedError struct {
	Destination *url.URL
}

func (e *AlreadyVisitedError) Error() string { _ = "STUB: not implemented"; return "" }

type htmlCallbackContainer struct {
	Selector string
	Function HTMLCallback
	active   atomic.Bool
}

type xmlCallbackContainer struct {
	Query    string
	Function XMLCallback
	active   atomic.Bool
}

type cookieJarSerializer struct {
	store storage.Storage
	lock  *sync.RWMutex
}

var collectorCounter uint32

type key int

const (
	ProxyURLKey key = iota
	CheckRevisitKey
)

const envVariablePrefix = "COLLY_"

var (
	ErrForbiddenDomain = errors.New("Forbidden domain")

	ErrMissingURL = errors.New("Missing URL")

	ErrMaxDepth = errors.New("Max depth limit reached")

	ErrForbiddenURL = errors.New("ForbiddenURL")

	ErrNoURLFiltersMatch = errors.New("No URLFilters match")

	ErrRobotsTxtBlocked = errors.New("URL blocked by robots.txt")

	ErrNoCookieJar = errors.New("Cookie jar is not available")

	ErrNoPattern = errors.New("No pattern defined in LimitRule")

	ErrEmptyProxyURL = errors.New("Proxy URL list is empty")

	ErrAbortedAfterHeaders = errors.New("Aborted after receiving response headers")

	ErrAbortedBeforeRequest = errors.New("Aborted before Do Request")

	ErrQueueFull = errors.New("Queue MaxSize reached")

	ErrMaxRequests = errors.New("Max Requests limit reached")

	ErrRetryBodyUnseekable = errors.New("Retry Body Unseekable")
)

var envMap = map[string]func(*Collector, string){
	"ALLOWED_DOMAINS": func(c *Collector, val string) {
		c.AllowedDomains = strings.Split(val, ",")
	},
	"CACHE_DIR": func(c *Collector, val string) {
		c.CacheDir = val
	},
	"DETECT_CHARSET": func(c *Collector, val string) {
		c.DetectCharset = isYesString(val)
	},
	"DISABLE_COOKIES": func(c *Collector, _ string) {
		c.backend.Client.Jar = nil
	},
	"DISALLOWED_DOMAINS": func(c *Collector, val string) {
		c.DisallowedDomains = strings.Split(val, ",")
	},
	"IGNORE_ROBOTSTXT": func(c *Collector, val string) {
		c.IgnoreRobotsTxt = isYesString(val)
	},
	"FOLLOW_REDIRECTS": func(c *Collector, val string) {
		if !isYesString(val) {
			c.redirectHandler = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
		}
	},
	"MAX_BODY_SIZE": func(c *Collector, val string) {
		size, err := strconv.Atoi(val)
		if err == nil {
			c.MaxBodySize = size
		}
	},
	"MAX_DEPTH": func(c *Collector, val string) {
		maxDepth, err := strconv.Atoi(val)
		if err == nil {
			c.MaxDepth = maxDepth
		}
	},
	"MAX_REQUESTS": func(c *Collector, val string) {
		maxRequests, err := strconv.ParseUint(val, 0, 32)
		if err == nil {
			c.MaxRequests = uint32(maxRequests)
		}
	},
	"PARSE_HTTP_ERROR_RESPONSE": func(c *Collector, val string) {
		c.ParseHTTPErrorResponse = isYesString(val)
	},
	"TRACE_HTTP": func(c *Collector, val string) {
		c.TraceHTTP = isYesString(val)
	},
	"USER_AGENT": func(c *Collector, val string) {
		c.UserAgent = val
	},
}

var urlParser = whatwgUrl.NewParser(whatwgUrl.WithPercentEncodeSinglePercentSign())

func NewCollector(options ...CollectorOption) *Collector { _ = "STUB: not implemented"; return nil }

func UserAgent(ua string) CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func Headers(headers map[string]string) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func MaxDepth(depth int) CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func MaxRequests(max uint32) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func AllowedDomains(domains ...string) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func ParseHTTPErrorResponse() CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func DisallowedDomains(domains ...string) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func DisallowedURLFilters(filters ...*regexp.Regexp) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func URLFilters(filters ...*regexp.Regexp) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func AllowURLRevisit() CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func MaxBodySize(sizeInBytes int) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func CacheDir(path string) CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func IgnoreRobotsTxt() CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func TraceHTTP() CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func StdlibContext(ctx context.Context) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func ID(id uint32) CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func Async(a ...bool) CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func DetectCharset() CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func Debugger(d debug.Debugger) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func CheckHead() CollectorOption { _ = "STUB: not implemented"; return *new(CollectorOption) }

func CacheExpiration(d time.Duration) CollectorOption {
	_ = "STUB: not implemented"
	return *new(CollectorOption)
}

func (c *Collector) Init() { _ = "STUB: not implemented"; return }

func (c *Collector) Appengine(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Collector) Visit(URL string) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) HasVisited(URL string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Collector) HasPosted(URL string, requestData map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Collector) Head(URL string) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) Post(URL string, requestData map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) PostRaw(URL string, requestData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) PostMultipart(URL string, requestData map[string][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) Request(method, URL string, requestData io.Reader, ctx *Context, hdr http.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) SetDebugger(d debug.Debugger) { _ = "STUB: not implemented"; return }

func (c *Collector) UnmarshalRequest(r []byte) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Collector) scrape(u, method string, depth int, requestData io.Reader, ctx *Context, hdr http.Header, checkRevisit bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) fetch(u, method string, depth int, requestData io.Reader, ctx *Context, hdr http.Header, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) requestCheck(parsedURL *url.URL, method string, getBody func() (io.ReadCloser, error), depth int, checkRevisit bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) checkFilters(URL, domain string) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) isDomainAllowed(domain string) bool { _ = "STUB: not implemented"; return false }

func (c *Collector) checkRobots(u *url.URL) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) String() string { _ = "STUB: not implemented"; return "" }

func (c *Collector) Wait() { _ = "STUB: not implemented"; return }

func (c *Collector) OnRequest(f RequestCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) OnResponseHeaders(f ResponseHeadersCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) OnRequestHeaders(f RequestCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) OnResponse(f ResponseCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) OnHTML(goquerySelector string, f HTMLCallback) {
	_ = "STUB: not implemented"
	return
}

func (c *Collector) OnXML(xpathQuery string, f XMLCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) OnHTMLDetach(goquerySelector string) { _ = "STUB: not implemented"; return }

func (c *Collector) OnXMLDetach(xpathQuery string) { _ = "STUB: not implemented"; return }

func (c *Collector) OnError(f ErrorCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) OnScraped(f ScrapedCallback) { _ = "STUB: not implemented"; return }

func (c *Collector) SetClient(client *http.Client) { _ = "STUB: not implemented"; return }

func (c *Collector) WithTransport(transport http.RoundTripper) { _ = "STUB: not implemented"; return }

func (c *Collector) DisableCookies() { _ = "STUB: not implemented"; return }

func (c *Collector) SetCookieJar(j http.CookieJar) { _ = "STUB: not implemented"; return }

func (c *Collector) SetRequestTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (c *Collector) SetStorage(s storage.Storage) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) SetProxy(proxyURL string) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) SetProxyFunc(p ProxyFunc) { _ = "STUB: not implemented"; return }

func createEvent(eventType string, requestID, collectorID uint32, kvargs map[string]string) *debug.Event {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) handleOnRequest(r *Request) { _ = "STUB: not implemented"; return }

func (c *Collector) handleOnResponse(r *Response) { _ = "STUB: not implemented"; return }

func (c *Collector) handleOnResponseHeaders(r *Response) { _ = "STUB: not implemented"; return }

func (c *Collector) handleOnRequestHeaders(r *Request) { _ = "STUB: not implemented"; return }

func (c *Collector) handleOnHTML(resp *Response) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) handleOnXML(resp *Response) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) handleOnError(response *Response, err error, request *Request, ctx *Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) cleanupCallbacks() { _ = "STUB: not implemented"; return }

func (c *Collector) handleOnScraped(r *Response) { _ = "STUB: not implemented"; return }

func (c *Collector) Limit(rule *LimitRule) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) Limits(rules []*LimitRule) error { _ = "STUB: not implemented"; return nil }

func (c *Collector) SetRedirectHandler(f func(req *http.Request, via []*http.Request) error) {
	_ = "STUB: not implemented"
	return
}

func (c *Collector) SetCookies(URL string, cookies []*http.Cookie) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) Cookies(URL string) []*http.Cookie { _ = "STUB: not implemented"; return nil }

func (c *Collector) Clone() *Collector { _ = "STUB: not implemented"; return nil }

func (c *Collector) checkRedirectFunc() func(req *http.Request, via []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collector) parseSettingsFromEnv() { _ = "STUB: not implemented"; return }

func (c *Collector) checkHasVisited(URL string, requestData map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func SanitizeFileName(fileName string) string { _ = "STUB: not implemented"; return "" }

func createFormReader(data map[string]string) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func createMultipartReader(boundary string, data map[string][]byte) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

var multipartFieldNameEscaper = strings.NewReplacer(
	`\`, `\\`,
	`"`, `\"`,
	"\r", "%0D",
	"\n", "%0A",
)

func escapeMultipartFieldName(s string) string { _ = "STUB: not implemented"; return "" }

func randomBoundary() string { _ = "STUB: not implemented"; return "" }

func isYesString(s string) bool { _ = "STUB: not implemented"; return false }

func createJar(s storage.Storage) http.CookieJar {
	_ = "STUB: not implemented"
	return *new(http.CookieJar)
}

func (j *cookieJarSerializer) SetCookies(u *url.URL, cookies []*http.Cookie) {
	_ = "STUB: not implemented"
	return
}

func (j *cookieJarSerializer) Cookies(u *url.URL) []*http.Cookie {
	_ = "STUB: not implemented"
	return nil
}

func isMatchingFilter(fs []*regexp.Regexp, d []byte) bool { _ = "STUB: not implemented"; return false }

func normalizeURL(u string) string { _ = "STUB: not implemented"; return "" }

func requestHash(url string, body io.Reader) uint64 { _ = "STUB: not implemented"; return 0 }
