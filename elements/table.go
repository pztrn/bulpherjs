// BulpherJS - bulma for GopherJS.
// Copyright (c) 2020 by Stanislav Nikitin <pztrn@pztrn.name>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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

	Header []string
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

	if t.options != nil {
		if len(t.options.Header) != 0 {
			theader := NewTHead()
			t.AddChild(theader)

			trow := NewTR()
			theader.AddChild(trow)

			for _, header := range t.options.Header {
				th := NewTH()
				th.SetTextContent(header)
				trow.AddChild(th)
			}
		}
	}

	t.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementTable)
}
