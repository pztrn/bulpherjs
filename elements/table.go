package elements

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"go.dev.pztrn.name/bulpherjs/common"
	"go.dev.pztrn.name/bulpherjs/metas"
)

const (
	tableDefaultClass = "table"
)

// TableOptions is a "table" HTML element configuration structure.
type TableOptions struct {
	Class       string
	IsBordered  bool
	IsFullWidth bool
	IsHoverable bool
	IsNarrow    bool
	IsStriped   bool
}

// Table is a controlling structure for "table" HTML element.
type Table struct {
	metas.Generic

	options *TableOptions
}

// NewTable creates new "table" HTML element controlling structure.
func NewTable(tableOpts *TableOptions) *Table {
	t := &Table{}
	t.initialize(tableOpts)

	return t
}

// Build builds element and calls Build() for all child elements.
func (t *Table) Build() *js.Object {
	if t.options == nil {
		t.AddClassesFromString(tableDefaultClass)
		t.BuildChilds(t.Object)

		return t.Object
	}

	if !strings.Contains(t.options.Class, "table") {
		t.options.Class += " table"
	}

	if t.options.Class != "" {
		t.AddClassesFromString(t.options.Class)
	}

	if t.options.IsBordered {
		t.AddClassesFromString("is-bordered")
	}

	if t.options.IsFullWidth {
		t.AddClassesFromString("is-fullwidth")
	}

	if t.options.IsHoverable {
		t.AddClassesFromString("is-hoverable")
	}

	if t.options.IsNarrow {
		t.AddClassesFromString("is-narrow")
	}

	if t.options.IsStriped {
		t.AddClassesFromString("is-striped")
	}

	t.BuildChilds(t.Object)

	return t.Object
}

// Initializes controlling structure with necessary parameters.
func (t *Table) initialize(tableOpts *TableOptions) {
	t.options = tableOpts

	t.InitializeGeneric()

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTable)
}
