package metas

import "github.com/gopherjs/gopherjs/js"

// Buildable is an interface which is used internally to ensure
// that item is buildable. Used mostly for child items.
type Buildable interface {
	Build() *js.Object
}
