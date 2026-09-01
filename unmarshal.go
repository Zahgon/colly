package colly

import (
	"reflect"

	"github.com/PuerkitoBio/goquery"
)

func (h *HTMLElement) Unmarshal(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (h *HTMLElement) UnmarshalWithMap(v interface{}, structMap map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalHTML(v interface{}, s *goquery.Selection, structMap map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalSelector(s *goquery.Selection, attrV reflect.Value, selector string) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalAttr(s *goquery.Selection, attrV reflect.Value, attrT reflect.StructField) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalStruct(s *goquery.Selection, selector string, attrV reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalPtr(s *goquery.Selection, selector string, attrV reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarshalSlice(s *goquery.Selection, selector, htmlAttr string, attrV reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func getDOMValue(s *goquery.Selection, attr string) string { _ = "STUB: not implemented"; return "" }
