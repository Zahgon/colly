package colly

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type HTMLElement struct {
	Name       string
	Text       string
	attributes []html.Attribute

	Request *Request

	Response *Response

	DOM *goquery.Selection

	Index int
}

func NewHTMLElementFromSelectionNode(resp *Response, s *goquery.Selection, n *html.Node, idx int) *HTMLElement {
	_ = "STUB: not implemented"
	return nil
}

func (h *HTMLElement) Attr(k string) string { _ = "STUB: not implemented"; return "" }

func (h *HTMLElement) ChildText(goquerySelector string) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *HTMLElement) ChildTexts(goquerySelector string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (h *HTMLElement) ChildAttr(goquerySelector, attrName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *HTMLElement) ChildAttrs(goquerySelector, attrName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (h *HTMLElement) ForEach(goquerySelector string, callback func(int, *HTMLElement)) {
	_ = "STUB: not implemented"
	return
}

func (h *HTMLElement) ForEachWithBreak(goquerySelector string, callback func(int, *HTMLElement) bool) {
	_ = "STUB: not implemented"
	return
}
