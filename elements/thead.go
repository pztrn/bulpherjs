package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// THead is a controlling structure for "thead" HTML element.
type THead struct {
	metas.Generic
}

// NewTHead creates new "thead" HTML element controlling structure.
func NewTHead() *THead {
	thead := &THead{}
	thead.initialize()

	return thead
}

// Build builds element and calls Build() for all child elements.
func (t *THead) Build() *js.Object {
	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *THead) initialize() {
	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTHead)
}
