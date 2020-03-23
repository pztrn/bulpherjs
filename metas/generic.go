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
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
)

// Generic is a structure embedded into other structures and provides
// some generic things like working with classes, IDs and others.
// Also this is the struct that holds real object.
type Generic struct {
	Buildable
	InnerItems
	Object *js.Object
}

// AddClassesFromString adds classes to object from string. Classes
// should be specified like "class1 class2 ... classN".
func (g *Generic) AddClassesFromString(classes string) {
	classesSplitted := strings.Split(strings.Trim(classes, " "), " ")

	for _, class := range classesSplitted {
		g.Object.Get(common.HTMLElementParameterClassList).Call(common.JSCallAdd, class)
	}
}

// InitializeGeneric initializes itself and all embedded things. This
// function SHOULD NOT be called manually by end-user, it should be
// called only for elements.
func (g *Generic) InitializeGeneric() {
	g.initializeInnerItems()
}

// SetTextContent sets text content for passed object. Note that not
// all HTML elements able to display it.
func (g *Generic) SetTextContent(text string) {
	g.Object.Set("textContent", text)
}
