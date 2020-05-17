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

package bulpherjs

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/widgets"
)

// Application is a controlling structure for whole application.
type Application struct {
	options   *ApplicationOptions
	startFunc func(*Application)

	HTML *widgets.HTML
}

// NewApplication creates new BulpherJS application.
func NewApplication(opts *ApplicationOptions) *Application {
	a := &Application{}
	a.initialize(opts)

	return a
}

// Build builds HTML for application.
func (a *Application) Build() {
	a.HTML.Build()
}

// Initializes controlling structure.
func (a *Application) initialize(opts *ApplicationOptions) {
	a.options = opts

	a.HTML = widgets.NewHTML()
}

// SetStartFunction sets start function which will be executed on app's
// Start() call.
func (a *Application) SetStartFunction(f func(*Application)) {
	a.startFunc = f
}

// SetTitle sets passed title as page title.
func (a *Application) SetTitle(title string) {
	a.HTML.Head.SetTitle(title)
}

// Start starts application execution.
func (a *Application) Start() {
	document := js.Global.Get(common.HTMLElementDocument)
	document.Call(common.JSCallAddEventListener, common.JSEventDOMContentLoaded, func(event *js.Object) {
		a.startFunc(a)

		a.Build()
	})
}
