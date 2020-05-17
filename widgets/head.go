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
	"go.dev.pztrn.name/bulpherjs/elements"
	"go.dev.pztrn.name/bulpherjs/metas"
)

const (
	bulmaVersion = "0.8.0"
)

// HeadOptions is a "head" HTML head element configuration structure.
type HeadOptions struct {
}

// Head is a controlling structure for "head" HTML head element.
type Head struct {
	metas.Generic

	options *HeadOptions
}

// NewHeader creates new "head" HTML element controlling structure.
func NewHead(opts *HeadOptions) *Head {
	hw := &Head{}
	hw.initialize(opts)

	return hw
}

// Build builds header.
func (hw *Head) Build() *js.Object {
	hw.BuildChilds(hw.Object)

	return hw.Object
}

// Initializes controlling structure.
func (hw *Head) initialize(opts *HeadOptions) {
	hw.options = opts

	hw.InitializeGeneric()

	hw.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallGetElementsByTagName, common.HTMLElementHead).Index(0)

	// Add default header elements.
	hw.AddChild(elements.NewLink(&elements.LinkOptions{
		Href: "//cdn.jsdelivr.net/npm/bulma@" + bulmaVersion + "/css/bulma.min.css",
		ID:   "",
		Rel:  "stylesheet",
	}))

	hw.AddChild(elements.NewMeta(&elements.MetaOptions{
		Content: "width=device-width, initial-scale=1",
		Name:    "viewport",
	}))
}

// SetTitle sets title for whole page.
func (h *Head) SetTitle(title string) {
	js.Global.Get(common.HTMLElementDocument).Set(common.HTMLHeadElementTitle, title)
}
