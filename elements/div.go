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

// DivOptions is a "div" HTML element configuration structure.
type DivOptions struct {
	Class string
	ID    string
}

// Div is a controlling structure for "div" HTML element.
type Div struct {
	metas.Generic
	options *DivOptions
}

// NewDiv creates new "div" HTML element controlling structure.
func NewDiv(divOptions *DivOptions) *Div {
	d := &Div{}
	d.initialize(divOptions)

	return d
}

// Build builds element and calls Build() for all child elements.
func (d *Div) Build() *js.Object {
	d.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementDiv)

	if d.options == nil {
		return d.Object
	}

	if d.options.Class != "" {
		d.AddClassesFromString(d.options.Class)
	}

	if d.options.ID != "" {
		d.Object.Set(common.HTMLElementParameterID, d.options.ID)
	}

	d.BuildChilds(d.Object)

	return d.Object
}

// Initializes controlling structure with necessary parameters.
func (d *Div) initialize(divOptions *DivOptions) {
	d.options = divOptions

	d.InitializeGeneric()
}
