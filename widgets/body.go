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

// Body is a controlling structure for "body" HTML head element.
type Body struct {
	metas.Generic
}

// NewBody creates new "body" HTML element controlling structure.
func NewBody() *Body {
	b := &Body{}
	b.initialize()

	return b
}

// Build builds child elements.
func (b *Body) Build() *js.Object {
	// This function is special - we should try to get <body> element only
	// after DOMContentLoaded event was fired, otherwise there is no
	// guarantee that <body> will ever exist.
	b.Object = js.Global.Get(common.HTMLElementDocument).Get(common.HTMLElementBody)

	b.BuildChilds(b.Object)

	return b.Object
}

// Initializes controlling structure.
func (b *Body) initialize() {
	b.InitializeGeneric()
}