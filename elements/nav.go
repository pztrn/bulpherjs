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
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

const (
	navDefaultClass = "navbar"
)

// NavOptions is a "nav" HTML element configuration structure.
type NavOptions struct {
	Class string
	Data  map[string]string
	Role  string
}

// Nav is a controlling structure for "nav" HTML element.
type Nav struct {
	metas.Generic
	options *NavOptions
}

// NewNav creates new "nav" HTML element controlling structure.
func NewNav(navOptions *NavOptions) *Nav {
	n := &Nav{}
	n.initialize(navOptions)

	return n
}

// Build builds element and calls Build() for all child elements.
func (n *Nav) Build() *js.Object {
	n.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementNav)

	if n.options == nil {
		n.AddClassesFromString(navDefaultClass)
		n.BuildChilds(n.Object)

		return n.Object
	}

	if !strings.Contains(n.options.Class, "navbar") {
		n.options.Class += " navbar"
	}

	if n.options.Class != "" {
		n.AddClassesFromString(n.options.Class)
	}

	if n.options.Role != "" {
		n.Object.Set(common.HTMLElementParameterRole, n.options.Role)
	}

	if n.options.Data != nil {
		for k, v := range n.options.Data {
			n.Object.Set(k, v)
		}
	}

	n.BuildChilds(n.Object)

	return n.Object
}

// Initializes controlling structure with necessary parameters.
func (n *Nav) initialize(navOptions *NavOptions) {
	n.options = navOptions

	n.InitializeGeneric()
}
