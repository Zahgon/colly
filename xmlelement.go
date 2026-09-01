package colly

import (
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html"
)

type XMLElement struct {
	Name       string
	Text       string
	attributes interface{}

	Request *Request

	Response *Response

	DOM    interface{}
	isHTML bool

	Index int
}

func NewXMLElementFromHTMLNode(resp *Response, s *html.Node) *XMLElement {
	_ = "STUB: not implemented"
	return nil
}

func NewXMLElementFromXMLNode(resp *Response, s *xmlquery.Node) *XMLElement {
	_ = "STUB: not implemented"
	return nil
}

func (h *XMLElement) Attr(k string) string { _ = "STUB: not implemented"; return "" }

func (h *XMLElement) ChildText(xpathQuery string) string { _ = "STUB: not implemented"; return "" }

func (h *XMLElement) ChildAttr(xpathQuery, attrName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (h *XMLElement) ChildAttrs(xpathQuery, attrName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (h *XMLElement) ChildTexts(xpathQuery string) []string { _ = "STUB: not implemented"; return nil }
