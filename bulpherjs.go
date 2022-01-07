package bulpherjs

import (
	"github.com/gopherjs/gopherjs/js"
	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/elements"
)

// Application is a controlling structure for whole application.
type Application struct {
	options   *ApplicationOptions
	startFunc func(*Application)

	HTML *elements.HTML
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

	a.HTML = elements.NewHTML()
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
