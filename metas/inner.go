package metas

import (
	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
)

// InnerItems is a meta structure that should be embedded in all other
// structures (notably element controlling ones). It provide ability to
// add child elements.
type InnerItems struct {
	innerItems []Buildable
}

// AddChild adds child to element.
func (it *InnerItems) AddChild(object Buildable) {
	it.innerItems = append(it.innerItems, object)
}

// BuildChilds build child elements and adds them to parent.
func (it *InnerItems) BuildChilds(parent *js.Object) {
	for _, item := range it.innerItems {
		parent.Call(common.JSCallAppendChild, item.Build())
	}
}

// Initializes child elements storage.
func (it *InnerItems) initializeInnerItems() {
	it.innerItems = make([]Buildable, 0)
}
