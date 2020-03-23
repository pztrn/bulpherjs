package metas

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
)

// Generic is a structure embedded into other structures and provides
// some generic things like working with classes, IDs and others.
// Also this is the struct that holds real object.
type Generic struct {
	Buildable
	InnerItems
	Object *js.Object
}

// AddClassesFromString adds classes to object from string. Classes
// should be specified like "class1 class2 ... classN".
func (g *Generic) AddClassesFromString(classes string) {
	classesSplitted := strings.Split(strings.Trim(classes, " "), " ")

	for _, class := range classesSplitted {
		g.Object.Get(common.HTMLElementParameterClassList).Call(common.JSCallAdd, class)
	}
}

// InitializeGeneric initializes itself and all embedded things. This
// function SHOULD NOT be called manually by end-user, it should be
// called only for elements.
func (g *Generic) InitializeGeneric() {
	g.initializeInnerItems()
}

// SetTextContent sets text content for passed object. Note that not
// all HTML elements able to display it.
func (g *Generic) SetTextContent(text string) {
	g.Object.Set("textContent", text)
}
