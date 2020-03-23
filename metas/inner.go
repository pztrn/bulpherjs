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

package metas

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
)

// InnerItems is a meta structure that should be embedded in all other
// structures (notably element controlling ones). It provide ability to
// add child elements.
type InnerItems struct {
	innerItems []Buildable
}

// AddChild adds child to element.
func (it *InnerItems) AddChild(object Buildable) {
	it.innerItems = append(it.innerItems, object)
}

// BuildChilds build child elements and adds them to parent.
func (it *InnerItems) BuildChilds(parent *js.Object) {
	for _, item := range it.innerItems {
		parent.Call(common.JSCallAppendChild, item.Build())
	}
}

// Initializes child elements storage.
func (it *InnerItems) initializeInnerItems() {
	it.innerItems = make([]Buildable, 0)
}
