package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// TFoot is a controlling structure for "tfoot" HTML element.
type TFoot struct {
	metas.Generic
}

// NewTFoot creates new "tfoot" HTML element controlling structure.
func NewTFoot() *TFoot {
	tfoot := &TFoot{}
	tfoot.initialize()

	return tfoot
}

// Build builds element and calls Build() for all child elements.
func (t *TFoot) Build() *js.Object {
	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *TFoot) initialize() {
	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTFoot)
}
