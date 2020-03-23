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

// AOptions is an "a" HTML element configuration structure.
type AOptions struct {
	Class string
	Href  string
	ID    string
	Text  string
}

// A is a controlling structure that describes everything for "a" HTML
// element.
type A struct {
	metas.Generic
	options *AOptions
}

// NewA creates new "a" element controlling structure.
func NewA(opts *AOptions) *A {
	a := &A{}
	a.initialize(opts)

	return a
}

// Build builds element and calls Build() for all child elements.
func (a *A) Build() *js.Object {
	a.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementA)

	if a.options == nil {
		return a.Object
	}

	if a.options.Class != "" {
		a.AddClassesFromString(a.options.Class)
	}

	if a.options.Href != "" {
		a.Object.Set(common.HTMLElementParameterHref, a.options.Href)
	}

	if a.options.ID != "" {
		a.Object.Set(common.HTMLElementParameterID, a.options.ID)
	}

	if a.options.Text != "" {
		a.Object.Set("textContent", a.options.Text)
	}

	a.BuildChilds(a.Object)

	return a.Object
}

// Initializes controlling structure with necessary parameters.
func (a *A) initialize(opts *AOptions) {
	a.options = opts

	a.InitializeGeneric()
}
