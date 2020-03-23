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

// LinkOptions is a "link" HTML head element configuration structure.
type LinkOptions struct {
	Href string
	ID   string
	Rel  string
}

// Link is a controlling structure for "link" HTML head element.
type Link struct {
	options *LinkOptions
}

// NewLink creates new "link" HTML head controlling structure.
func NewLink(linkOptions *LinkOptions) *Link {
	l := &Link{}
	l.initialize(linkOptions)

	return l
}

// Build builds element.
func (l *Link) Build() *js.Object {
	link := js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementLink)

	if l.options == nil {
		return link
	}

	if l.options.Href != "" {
		link.Set(common.HTMLElementParameterHref, l.options.Href)
	}

	if l.options.Rel != "" {
		link.Set(common.HTMLElementParameterRel, l.options.Rel)
	}

	return link
}

// Initializes controlling structure with necessary parameters.
func (l *Link) initialize(linkOptions *LinkOptions) {
	l.options = linkOptions
}
