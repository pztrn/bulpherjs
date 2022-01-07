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
