package storage

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

type Storage interface {
	Init() error

	Visited(requestID uint64) error

	IsVisited(requestID uint64) (bool, error)

	Cookies(u *url.URL) string

	SetCookies(u *url.URL, cookies string)
}

type InMemoryStorage struct {
	visitedURLs map[uint64]bool
	lock        *sync.RWMutex
	jar         *cookiejar.Jar
}

func (s *InMemoryStorage) Init() error { _ = "STUB: not implemented"; return nil }

func (s *InMemoryStorage) Visited(requestID uint64) error { _ = "STUB: not implemented"; return nil }

func (s *InMemoryStorage) IsVisited(requestID uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *InMemoryStorage) Cookies(u *url.URL) string { _ = "STUB: not implemented"; return "" }

func (s *InMemoryStorage) SetCookies(u *url.URL, cookies string) { _ = "STUB: not implemented"; return }

func (s *InMemoryStorage) Close() error { _ = "STUB: not implemented"; return nil }

func StringifyCookies(cookies []*http.Cookie) string { _ = "STUB: not implemented"; return "" }

func UnstringifyCookies(s string) []*http.Cookie { _ = "STUB: not implemented"; return nil }

func ContainsCookie(cookies []*http.Cookie, name string) bool {
	_ = "STUB: not implemented"
	return false
}
