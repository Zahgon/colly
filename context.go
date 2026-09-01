package colly

import (
	"sync"
)

type Context struct {
	contextMap map[string]interface{}
	lock       *sync.RWMutex
}

func NewContext() *Context { _ = "STUB: not implemented"; return nil }

func (c *Context) UnmarshalBinary(_ []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Context) MarshalBinary() (_ []byte, _ error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Context) Put(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (c *Context) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (c *Context) GetAny(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (c *Context) ForEach(fn func(k string, v interface{}) interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) Clone() *Context { _ = "STUB: not implemented"; return nil }
