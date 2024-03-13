package tmplz

import (
	"fmt"
	"slices"
	"strings"
)

// Node represents a single variable or literal within a template.
// It can have child nodes, allowing for nested variable resolution.
type Node struct {
	// Parent points to the node's parent in the variable tree, if any.
	Parent *Node
	// Start is the index of the first character of the node text in the parent template.
	Start int
	// End is the index of the last character of the node text in the parent template.
	End int
	// Text contains the literal text of the node or the resolved value of a variable.
	Text String
	// Children holds any child nodes, representing nested variables.
	Children []*Node
}

// AppendRunes adds runes to the node's text, updating the node and its ancestors accordingly.
func (n *Node) AppendRunes(s ...rune) {
	n.Text = append(n.Text, s...)
	n.End += len(s)
	if n.Parent != nil {
		n.Parent.AppendRunes(s...)
	}
}

// AppendAlpha adds the ALPHA_RUNE to the node's text, indicating the start of a variable.
func (n *Node) AppendAlpha() {
	n.Text = append(n.Text, ALPHA_RUNE)
	n.End += 1
	if n.Parent != nil {
		n.Parent.AppendAlpha()
	}
}

// AppendOmega adds the OMEGA_RUNE to the node's text, indicating the end of a variable.
func (n *Node) AppendOmega(real bool) {
	n.Text = append(n.Text, OMEGA_RUNE)
	if real {
		n.End += 1
	}
	if n.Parent != nil {
		if !real {
			n.End += 1
		}
		n.Parent.AppendOmega(real)
	}
}

// DeepStart calculates the absolute start index of the node's text within the root template.
func (n Node) DeepStart() int {
	if n.Parent != nil {
		return n.Start + n.Parent.DeepStart()
	}
	return n.Start
}

// Verify checks the consistency of the node's text with its children's texts.
func (n Node) Verify() {
	for i := 0; i < len(n.Children); i++ {
		c := n.Children[i]
		if c == nil {
			continue
		}
		slice := n.Text.Slice(c.Start, c.End)
		if !slice.Equals(c.Text) {
			panic(fmt.Errorf("verify node: %s :: %s => %s", n.Text, c.Debug(), slice))
		}
		c.Verify()
	}
}

// String returns a string representation of the node and its children.
func (n Node) String() string {
	var subVars string
	for _, sv := range n.Children {
		subVars += sv.String()
	}
	return fmt.Sprintf("{%s%s}", n.Text, subVars)
}

// Debug returns a detailed string representation of the node, including indices and child nodes.
func (n Node) Debug() string {
	var subVars string
	for _, sv := range n.Children {
		subVars += sv.Debug()
	}
	return fmt.Sprintf("{'%s'[%d,%d]%s}", n.Text, n.Start, n.End, subVars)
}

// Dot generates a Graphviz dot graph representation of the node and its children.
func (n Node) Dot() string {
	var sb strings.Builder
	sb.WriteString("digraph g {\n")
	sb.WriteString(fmt.Sprintf("\t\"%s\"\n", n.Text))
	sb.WriteString(n.dot())
	sb.WriteString("}\n")
	return sb.String()
}

func (n Node) dot() string {
	var sb strings.Builder
	for _, c := range n.Children {
		sb.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\" [label = \"[%d, %d]\"]\n", n.Text, c.Text, c.Start, c.End))
		sb.WriteString(c.dot())
	}
	return sb.String()
}

// Clone creates a deep copy of the Node and its children.
func (n Node) Clone() *Node {
	clone := Node{
		Text:   n.Text.Clone(),
		Start:  n.Start,
		End:    n.End,
		Parent: nil,
	}
	for _, child := range n.Children {
		childClone := child.Clone()
		childClone.Parent = &clone
		clone.Children = append(clone.Children, childClone)
	}
	return &clone
}

// Execute clones the [Node] and executes the [Assignments] on the clone.
//
// Returns the result as well as the clone.
func (n Node) Execute(assignments Assignments) (result string, clone *Node) {
	clone = n.Clone()
	clone.ExecuteInPlace(assignments)
	result = clone.Text.String()
	return
}

// ExecuteInPlace modifies the [Node] using the data from the [Assignments].
//
// To select how "@_" gets evaluated, set:
//
//	assignments.Assign("", "_") // default value
//
// Returns when no more assignments are left to assign.
func (n *Node) ExecuteInPlace(assignments Assignments) (anyChange bool) {
	if _, ok := assignments[""]; !ok {
		assignments.Assign("", "_") // default value
	}

	if len(n.Children) == 0 {
		varName := n.Text.String()
		if strings.HasPrefix(varName, ALPHA) {
			varName = varName[1:]
		}
		if strings.HasSuffix(varName, OMEGA) {
			varName = varName[:len(varName)-1]
		}
		value, ok := assignments.Get(varName)
		if ok {
			n.Text = value
		}
		return ok
	}

	offset := 0
	for ci := 0; ci < len(n.Children); ci++ {
		child := n.Children[ci]
		if ok := child.ExecuteInPlace(assignments); ok {
			n.Text = n.Text.Splice(child.Start+offset, child.End+offset, child.Text)
			offset += child.Start + child.Text.Len() - child.End
			child.End = child.Start + child.Text.Len()
			anyChange = true
			if !strings.Contains(child.Text.String(), "@") {
				n.Children = slices.Delete(n.Children, ci, ci+1)
				ci--
			}
		}
	}
	defer n.Verify()
	if anyChange {
		return n.ExecuteInPlace(assignments) || anyChange
	}
	return anyChange

}
