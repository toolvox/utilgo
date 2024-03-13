package tmplz

import (
	"fmt"
	"slices"
	"strings"
)

type Template struct {
	Text     String
	Children []*Node
}

// String returns a string representation of the template, including its text and variables.
func (t Template) String() string {
	childrenString := t.render(func(n *Node) string { return n.String() })
	return fmt.Sprintf("Template:\n\t%s\nVariables:%s\n", t.Text, childrenString)
}

// Debug returns a detailed string representation of the template for debugging purposes.
func (t Template) Debug() string {
	childrenDebug := t.render(func(n *Node) string { return n.Debug() })
	return fmt.Sprintf("Template:\n\t%s\nVariables:%s\n", t.Text, childrenDebug)
}

// Dot generates a Graphviz dot graph representation of the template and its variables.
func (t Template) Dot() string {
	return fmt.Sprintf("digraph g {\n%s}\n", t.dot())
}

func (t Template) dot() string {
	childrenDot := t.render(func(n *Node) string {
		return fmt.Sprintf("\"tmpl root\" -> \"%s\" [label = \"[%d, %d]\"]\n%s",
			n.Text, n.Start, n.End, n.dot())
	})
	return fmt.Sprintf("\t\"tmpl root\" [label=\"%s\", shape=diamond]\n%s",
		t.Text, childrenDot)
}

// render helps in rendering the template or its components based on a provided renderer function.
func (t Template) render(nodeRenderer func(*Node) string) string {
	var sb strings.Builder
	for _, v := range t.Children {
		sb.WriteString(fmt.Sprint("\n\t", nodeRenderer(v)))
	}
	return sb.String()
}

// Clone creates a deep copy of the Template, including all child nodes.
func (t Template) Clone() *Template {
	clone := Template{Text: t.Text.Clone()}
	for _, child := range t.Children {
		clone.Children = append(clone.Children, child.Clone())
	}
	return &clone
}

// Verify checks the consistency of the template's text with its nodes.
func (t Template) Verify() {
	for i := 0; i < len(t.Children); i++ {
		c := t.Children[i]
		if c == nil {
			continue
		}
		slice := t.Text.Slice(c.Start, c.End)
		if !slice.Equals(c.Text) && !strings.HasPrefix(c.Text.String(), slice.String()) {
			panic(fmt.Errorf("verify template: %s :: %s => %s", t.Text, c.Debug(), slice))
		}
		c.Verify()
	}
}

// Execute clones the [Node] and executes the [Assignments] on the clone.
//
// Returns the result as well as the clone.
func (t Template) Execute(assignments Assignments) (result string, clone *Template) {
	clone = t.Clone()
	clone.ExecuteInPlace(assignments)
	result = clone.Text.String()
	return
}

// ExecuteInPlace modifies the [Template] using the data from the [Assignments].
//
// To select how "@_" gets evaluated, set:
//
//	assignments.Assign("", "_") // default value
//
// Returns when no more assignments are left to assign.
func (t *Template) ExecuteInPlace(assignments Assignments) (anyChange bool) {
	offset := 0
	for ci := 0; ci < len(t.Children); ci++ {
		child := t.Children[ci]
		ok := child.ExecuteInPlace(assignments)
		if !ok {
			continue
		}
		if !strings.Contains(child.Text.String(), ALPHA) {
			t.Children = slices.Delete(t.Children, ci, ci+1)
			ci--
		}
		child.Start += offset
		child.End += offset
		t.Text = t.Text.Splice(child.Start, child.End, child.Text)
		offset += child.Text.Len() - (child.End - child.Start)
		anyChange = true
	}

	defer func() {
		if anyChange {
			*t = *ParseTemplate(t.Text.String())
			if len(t.Children) > 0 {
				t.ExecuteInPlace(assignments)
			}
		}
		t.Verify()
	}()

	if anyChange {
		return t.ExecuteInPlace(assignments) || anyChange
	}
	return anyChange
}
