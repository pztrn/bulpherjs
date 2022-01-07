package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// TR is a controlling structure for "tr" HTML element.
type TR struct {
	metas.Generic
}

// NewTR creates new "tr" HTML element controlling structure.
func NewTR() *TR {
	tr := &TR{}
	tr.initialize()

	return tr
}

// Build builds element and calls Build() for all child elements.
func (t *TR) Build() *js.Object {
	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *TR) initialize() {
	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTR)
}
