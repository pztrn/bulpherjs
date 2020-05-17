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

package widgets

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// HTML is a controlling structure for HTML page.
type HTML struct {
	metas.Generic

	Head *Head
	Body *Body
}

// NewHTML creates new HTML page.
func NewHTML() *HTML {
	h := &HTML{}
	h.initialize()

	return h
}

// Build builds HTML and all it's childs. For HTML object it'll return
// nil.
func (h *HTML) Build() *js.Object {
	h.BuildChilds(h.Object)

	js.Global.Get(common.HTMLElementDocument).Set(common.HTMLElementHTML, h.Object)

	return nil
}

// Initializes controlling structure.
func (h *HTML) initialize() {
	h.InitializeGeneric()

	h.Object = js.Global.Get(common.HTMLElementDocument).Get(common.HTMLElementDocumentElement)

	h.Head = NewHead(&HeadOptions{})
	h.Body = NewBody()

	h.AddChild(h.Head)
	h.AddChild(h.Body)

	// Set default title. Use Body.SetTitle to override in your application.
	h.Head.SetTitle("BulpherJS default application")
}
