package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// TD is a controlling structure for "td" HTML element.
type TD struct {
	metas.Generic
}

// NewTD creates new "td" HTML element controlling structure.
func NewTD() *TD {
	td := &TD{}
	td.initialize()

	return td
}

// Build builds element and calls Build() for all child elements.
func (t *TD) Build() *js.Object {
	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *TD) initialize() {
	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTD)
}
