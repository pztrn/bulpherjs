package elements

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

// NavOptions is a "nav" HTML element configuration structure.
type NavOptions struct {
	Class string
	Data  map[string]string
	Role  string
}

// Nav is a controlling structure for "nav" HTML element.
type Nav struct {
	metas.Generic
	options *NavOptions
}

// NewNav creates new "nav" HTML element controlling structure.
func NewNav(navOptions *NavOptions) *Nav {
	n := &Nav{}
	n.initialize(navOptions)

	return n
}

// Build builds element and calls Build() for all child elements.
func (n *Nav) Build() *js.Object {
	n.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementNav)

	if n.options == nil {
		return n.Object
	}

	if n.options.Class != "" {
		n.AddClassesFromString(n.options.Class)
	}

	if n.options.Role != "" {
		n.Object.Set(common.HTMLElementParameterRole, n.options.Role)
	}

	if n.options.Data != nil {
		for k, v := range n.options.Data {
			n.Object.Set(k, v)
		}
	}

	n.BuildChilds(n.Object)

	return n.Object
}

// Initializes controlling structure with necessary parameters.
func (n *Nav) initialize(navOptions *NavOptions) {
	n.options = navOptions

	n.InitializeGeneric()
}
