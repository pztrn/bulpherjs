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
	"go.dev.pztrn.name/bulpherjs/metas"
)

// POptions is a "p" HTML element configuration structure.
type POptions struct {
	Class string
	ID    string
	Text  string
}

// P is a controlling structure for "p" HTML element.
type P struct {
	metas.Generic
	options *POptions
}

// NewP creates new "p" HTML element controlling structure.
func NewP(opts *POptions) *P {
	p := &P{}
	p.initialize(opts)

	return p
}

// Build builds element and calls Build() for all child elements.
func (p *P) Build() *js.Object {
	p.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementP)

	if p.options == nil {
		return p.Object
	}

	if p.options.Class != "" {
		p.AddClassesFromString(p.options.Class)
	}

	if p.options.ID != "" {
		p.Object.Set(common.HTMLElementParameterID, p.options.ID)
	}

	if p.options.Text != "" {
		p.Object.Set("textContent", p.options.Text)
	}

	p.BuildChilds(p.Object)

	return p.Object
}

// Initializes controlling structure with necessary parameters.
func (p *P) initialize(opts *POptions) {
	p.options = opts

	p.InitializeGeneric()
}
