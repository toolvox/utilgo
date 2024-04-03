// Package tmplz provides a templating engine with nested variables and dynamic variable assignment.
// It includes utilities for handling templates.
package tmplz

import (
	"fmt"

	"github.com/toolvox/utilgo/pkg/errs"
)

// ParseTemplates is similar to [ParseTemplates] but includes logging support.
func ParseTemplates(templateMap map[string]string) *MultiTemplate {
	var res MultiTemplate = MultiTemplate{
		Templates: make(map[string]*Template, len(templateMap)),
	}
	for k, v := range templateMap {
		res.Templates[k] = ParseTemplate(v)
	}

	return &res
}

// ParseTemplate is similar to [ParseTemplate] but includes logging support for parsing.
func ParseTemplate(s string) *Template {
	state := ParseState{
		Tokens: Lex(String(s)),
		Cursor: 0,
	}

	tmpl := Template{Text: String(s)}
	var varStart int
	for {
		if state.NextToken().Rune != ALPHA_RUNE {
			var ok bool
			varStart, ok = state.PeekNext(TokenAlpha)
			if !ok {
				break
			}
		}

		state.Cursor = varStart
		root, cursor, err := state.Parse(varStart, nil)
		if err != nil {
			state.Cursor++
			continue
		}
		tmpl.Children = append(tmpl.Children, root)
		state.Cursor = cursor
		varStart = cursor
	}

	return &tmpl
}

// ParseNode is similar to [ParseNode] but with added logging functionality.
func ParseNode(s String) (*Node, string, error) {
	if s.Len() == 0 || s[0] != ALPHA_RUNE {
		return nil, s.String(), errs.Newf("node must start with '%s'", ALPHA)
	}

	state := ParseState{
		Tokens: Lex(String(s)),
		Cursor: 0,
	}

	root, cursor, err := state.Parse(0, nil)
	if err != nil || cursor == 0 {
		return nil, s.String(), errs.Wrap(fmt.Sprintf("'%s' does not begin a valid node", s), err)
	}

	return root, s.Slice(cursor, -1).String(), nil
}
