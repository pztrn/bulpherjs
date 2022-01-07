package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// DivOptions is a "div" HTML element configuration structure.
type DivOptions struct {
	Class string
	ID    string
}

// Div is a controlling structure for "div" HTML element.
type Div struct {
	metas.Generic
	options *DivOptions
}

// NewDiv creates new "div" HTML element controlling structure.
func NewDiv(divOptions *DivOptions) *Div {
	d := &Div{}
	d.initialize(divOptions)

	return d
}

// Build builds element and calls Build() for all child elements.
func (d *Div) Build() *js.Object {
	d.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementDiv)

	if d.options == nil {
		return d.Object
	}

	if d.options.Class != "" {
		d.AddClassesFromString(d.options.Class)
	}

	if d.options.ID != "" {
		d.Object.Set(common.HTMLElementParameterID, d.options.ID)
	}

	d.BuildChilds(d.Object)

	return d.Object
}

// Initializes controlling structure with necessary parameters.
func (d *Div) initialize(divOptions *DivOptions) {
	d.options = divOptions

	d.InitializeGeneric()
}
