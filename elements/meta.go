// BulpherJS - bulma for GopherJS.
// Copyright (c) 2020 by Stanislav Nikitin <pztrn@pztrn.name>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
