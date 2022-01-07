package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
)

// MetaOptions is a "meta" HTML head element configuration structure.
type MetaOptions struct {
	Charset string
	Content string
	Name    string
}

// Meta is a controlling structure for "meta" HTML head element.
type Meta struct {
	options *MetaOptions
}

// NewMeta creates "meta" HTML head controlling structure.
func NewMeta(metaOptions *MetaOptions) *Meta {
	m := &Meta{}
	m.initialize(metaOptions)

	return m
}

// Build builds element.
func (m *Meta) Build() *js.Object {
	meta := js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementMeta)

	if m.options == nil {
		return meta
	}

	if m.options.Charset != "" {
		meta.Set(common.HTMLMetaParameterCharset, m.options.Charset)
	}

	if m.options.Content != "" {
		meta.Set(common.HTMLMetaParameterContent, m.options.Content)
	}

	if m.options.Name != "" {
		meta.Set(common.HTMLMetaParameterName, m.options.Name)
	}

	return meta
}

// Initializes controlling structure with necessary parameters.
func (m *Meta) initialize(metaOptions *MetaOptions) {
	m.options = metaOptions
}
