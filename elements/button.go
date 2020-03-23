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
	buttonDefaultClass = "button"
)

// ButtonOptions is a "button" HTML element configuration structure.
type ButtonOptions struct {
	Class          string
	ID             string
	OnClickHandler func(e *js.Object)
	Text           string
}

// Button is a controlling structure for "button" HTML element.
type Button struct {
	metas.Generic

	options *ButtonOptions
}

// NewButton creates new "button" HTML element controlling structure.
func NewButton(buttonOptions *ButtonOptions) *Button {
	b := &Button{}
	b.initialize(buttonOptions)

	return b
}

// Build builds element and calls Build() for all child elements.
func (b *Button) Build() *js.Object {
	b.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementButton)

	if b.options == nil {
		b.AddClassesFromString(buttonDefaultClass)
		b.BuildChilds(b.Object)

		return b.Object
	}

	// Button should always have "button" class.
	if !strings.Contains(b.options.Class, "button") {
		b.options.Class += " button"
	}

	if b.options.Class != "" {
		b.AddClassesFromString(b.options.Class)
	}

	if b.options.ID != "" {
		b.Object.Set("id", b.options.ID)
	}

	if b.options.Text != "" {
		b.Object.Set("textContent", b.options.Text)
	}

	if b.options.OnClickHandler != nil {
		b.Object.Set("onclick", b.options.OnClickHandler)
	}

	b.BuildChilds(b.Object)

	return b.Object
}

// Initializes controlling structure with necessary parameters.
func (b *Button) initialize(buttonOptions *ButtonOptions) {
	b.options = buttonOptions

	b.InitializeGeneric()
}
