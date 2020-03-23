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
	sectionDefaultClass = "section"
)

// SectionOptions is a "section" HTML element configuration structure.
type SectionOptions struct {
	Class string
}

// Section is a controlling structure for "section" HTML element.
type Section struct {
	metas.Generic

	options *SectionOptions
}

// NewSection creates new "section" HTML element controlling structure.
func NewSection(sectionOptions *SectionOptions) *Section {
	s := &Section{}
	s.initialize(sectionOptions)

	return s
}

// Build builds element and calls Build() for all child elements.
func (s *Section) Build() *js.Object {
	s.Object = js.Global.Get(common.HTMLElementDocument).Call(common.JSCallCreateElement, common.HTMLElementSection)

	if s.options == nil {
		s.AddClassesFromString(sectionDefaultClass)
		s.BuildChilds(s.Object)

		return s.Object
	}

	if !strings.Contains(s.options.Class, "section") {
		s.options.Class += " section"
	}

	if s.options.Class != "" {
		s.AddClassesFromString(s.options.Class)
	}

	s.BuildChilds(s.Object)

	return s.Object
}

// Initializes controlling structure with necessary parameters.
func (s *Section) initialize(sectionOptions *SectionOptions) {
	s.options = sectionOptions

	s.InitializeGeneric()
}
