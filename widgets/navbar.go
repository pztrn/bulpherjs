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

	"go.dev.pztrn.name/bulpherjs/elements"
)

// NavBarOptions is a "nav" HTML body element configuration structure.
type NavBarOptions struct {
	IsDark bool
	Title  string
}

// NavBar is a controlling structure for "nav" HTML body element.
type NavBar struct {
	options *NavBarOptions

	navbar *elements.Nav
}

// NewNavBar creates new "nav" HTML element controlling structure.
func NewNavBar(opts *NavBarOptions) *NavBar {
	n := &NavBar{}
	n.initialize(opts)

	return n
}

// Build builds element and calls Build() for all child elements.
func (n *NavBar) Build() *js.Object {
	return n.navbar.Build()
}

// Initializes controlling structure.
func (n *NavBar) initialize(opts *NavBarOptions) {
	n.options = opts

	navopts := &elements.NavOptions{
		Class: "navbar",
		Data:  map[string]string{},
		Role:  "navigation",
	}

	n.navbar = elements.NewNav(navopts)

	if n.options.IsDark {
		n.navbar.AddClassesFromString("is-dark")
	}

	if n.options.Title != "" {
		brandDiv := elements.NewDiv(&elements.DivOptions{
			Class: "navbar-brand",
		})

		n.navbar.AddChild(brandDiv)

		navMainLink := elements.NewA(&elements.AOptions{
			Class: "navbar-item",
			Href:  "/",
			Text:  n.options.Title,
		})

		brandDiv.AddChild(navMainLink)
	}
}
