package tmplz

import (
	"fmt"
	"strings"

	"github.com/toolvox/utilgo/pkg/maputil"
	"github.com/toolvox/utilgo/pkg/stringutil"
)

var sortKeys = maputil.SortedKeys[map[string]*Template]

type MultiTemplate struct {
	Templates map[string]*Template
}

var ioIn = stringutil.IndentOption{
	IndentPrefix: "\t",
	IndentString: "\t",
}
var ioOut = stringutil.IndentOption{
	EmptyLines:   stringutil.EmptyLine_TrimHead | stringutil.EmptyLine_TrimTail,
	IndentPrefix: "",
	IndentString: "\t",
}

// String returns a string representation of all templates in MultiTemplate, formatted with indentation for readability.
// Each template's name and content are included.
func (mt MultiTemplate) String() string {
	var sb strings.Builder
	for _, tk := range sortKeys(mt.Templates) {
		t := mt.Templates[tk]
		sb.WriteString(tk + ":\n")
		sb.WriteString(ioIn.Indent(t.String()))
	}
	return ioOut.Indent(sb.String())
}

// Debug returns a detailed string representation of all templates in MultiTemplate, including template names and their debug information.
func (mt MultiTemplate) Debug() string {
	var sb strings.Builder
	for _, tk := range sortKeys(mt.Templates) {
		t := mt.Templates[tk]
		sb.WriteString(tk + ":\n")
		sb.WriteString(ioIn.Indent(t.Debug()))
	}
	return ioOut.Indent(sb.String())
}

// Dot generates a Graphviz dot graph of the MultiTemplate, illustrating the relationships between templates.
func (mt MultiTemplate) Dot() string {
	var sb strings.Builder
	for i, tk := range sortKeys(mt.Templates) {
		t := mt.Templates[tk]
		sb.WriteString(fmt.Sprintf("\n\t\"%s\" [shape=rect]\n\t\"%s\" -> \"tmpl root %d\" [color=red]\n", tk, tk, i))
		sb.WriteString(strings.ReplaceAll(t.dot(), "tmpl root", fmt.Sprintf("tmpl root %d", i)))
		i++
	}
	return fmt.Sprintf("digraph g {\n%s}\n", sb.String())

}

// Clone creates a deep copy of the MultiTemplate, including all its templates.
func (mt MultiTemplate) Clone() *MultiTemplate {
	clone := MultiTemplate{
		Templates: make(map[string]*Template, len(mt.Templates)),
	}
	for k, v := range mt.Templates {
		clone.Templates[strings.Clone(k)] = v.Clone()
	}
	return &clone
}

// Execute clones the [MultiTemplate] and executes the [Assignments] on the clone.
//
// Returns the result as well as the clone.
// By default, the result will be the value that has the key ".".
// If missing, an empty string is returned.
func (mt MultiTemplate) Execute(assignments Assignments) (string, *MultiTemplate) {
	clone := mt.Clone()
	clone.ExecuteInPlace(assignments)
	if root, ok := clone.Templates["."]; ok {
		return root.Text.String(), clone
	}
	return "", clone
}

// ExecuteInPlace modifies the [MultiTemplate] using the data from the [Assignments].
//
// To select how "@_" gets evaluated, set:
//
//	assignments.Assign("", "_") // default value
//
// Returns when no more assignments are left to assign.
func (mt *MultiTemplate) ExecuteInPlace(assignments Assignments) (anyChange bool) {
	for _, k := range sortKeys(mt.Templates) {
		v := mt.Templates[k]
		assignments.Assign(k, v.Text.String())
	}

	for _, k := range sortKeys(mt.Templates) {
		v := mt.Templates[k]
		if v.ExecuteInPlace(assignments) {
			mt.Templates[k] = v
			assignments.Assign(k, v.Text)
			anyChange = true
		}
	}

	if anyChange {
		return mt.ExecuteInPlace(assignments) || anyChange
	}
	return anyChange
}
