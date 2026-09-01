package queue

import (
	"sync"

	whatwgUrl "github.com/nlnwa/whatwg-url/url"

	"github.com/gocolly/colly/v2"
)

const stop = true

var urlParser = whatwgUrl.NewParser(whatwgUrl.WithPercentEncodeSinglePercentSign())

type Storage interface {
	Init() error

	AddRequest([]byte) error

	GetRequest() ([]byte, error)

	QueueSize() (int, error)
}

type Queue struct {
	Threads int
	storage Storage
	wake    chan struct{}
	mut     sync.Mutex
	running bool
}

type InMemoryQueueStorage struct {
	MaxSize int
	lock    *sync.RWMutex
	size    int
	first   *inMemoryQueueItem
	last    *inMemoryQueueItem
}

type inMemoryQueueItem struct {
	Request []byte
	Next    *inMemoryQueueItem
}

func New(threads int, s Storage) (*Queue, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *Queue) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *Queue) AddURL(URL string) error { _ = "STUB: not implemented"; return nil }

func (q *Queue) AddRequest(r *colly.Request) error { _ = "STUB: not implemented"; return nil }

func (q *Queue) storeRequest(r *colly.Request) error { _ = "STUB: not implemented"; return nil }

func (q *Queue) Size() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (q *Queue) Run(c *colly.Collector) error { _ = "STUB: not implemented"; return nil }

func (q *Queue) Stop() { _ = "STUB: not implemented"; return }

func (q *Queue) loop(c *colly.Collector, requestc chan<- *colly.Request, complete <-chan struct{}, errc chan<- error) {
	_ = "STUB: not implemented"
	return
}

func independentRunner(requestc <-chan *colly.Request, complete chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func (q *Queue) loadRequest(c *colly.Collector) (*colly.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *InMemoryQueueStorage) Init() error { _ = "STUB: not implemented"; return nil }

func (q *InMemoryQueueStorage) AddRequest(r []byte) error { _ = "STUB: not implemented"; return nil }

func (q *InMemoryQueueStorage) GetRequest() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *InMemoryQueueStorage) QueueSize() (int, error) { _ = "STUB: not implemented"; return 0, nil }
