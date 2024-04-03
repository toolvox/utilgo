package tmplz

import (
	"fmt"
	"strings"

	"github.com/toolvox/utilgo/pkg/errs"
)

// ParseState represents the state of parsing a template string, tracking tokens and cursor position.
type ParseState struct {
	Tokens Tokens // Tokens represents the lexed tokens of the template string.
	Cursor int    // Cursor represents the current position in the token slice.
}

// NewState initializes a new ParseState for parsing a template string.
func NewState(s string) *ParseState {
	return &ParseState{
		Tokens: Lex(String(s)),
	}
}

// String returns a string representation of the current parse state, including tokens.
func (s ParseState) String() string {
	return s.Tokens.String()
}

// Debug returns a detailed string representation of the parse state, indicating the cursor position.
func (s ParseState) Debug() string {
	var sb strings.Builder
	for i := 0; i < len(s.Tokens); i++ {
		if i == s.Cursor {
			sb.WriteRune('→')
		}
		sb.WriteRune(s.Tokens[i].Rune)
	}
	return sb.String()
}

// NextToken returns the next token in the sequence, or EOF if at the end.
func (s ParseState) NextToken() Token {
	if s.Cursor >= len(s.Tokens) {
		if len(s.Tokens) == 0 {
			return Token{TokenEOF, 0, 0}
		}
		last := s.Tokens[len(s.Tokens)-1].Pos
		return Token{TokenEOF, 0, last}
	}
	return s.Tokens[s.Cursor]
}

// PeekNext searches for the next occurrence of a token of the specified kind and returns its position.
func (s ParseState) PeekNext(kind TokenKind) (int, bool) {
	if s.Cursor+1 >= len(s.Tokens) {
		return -1, false
	}
	for n, t := range s.Tokens[1+s.Cursor:] {
		if t.Kind&kind != 0 {
			return s.Cursor + n + 1, true
		}
	}
	return -1, false
}

// Parse processes tokens from the current cursor position to construct a Node tree.
func (s *ParseState) Parse(from int, root *Node) (*Node, int, error) {
	if from >= len(s.Tokens) || s.NextToken().Kind != TokenAlpha {
		return nil, from, errs.Newf("node must start with '%s'", ALPHA)
	}

	offset, valid := 0, 0
	if root != nil {
		offset = root.DeepStart()
	}

	root = &Node{
		Parent: root,
		Start:  from - offset,
		End:    from - offset,
	}

	index, ok := s.PeekNext(TokenAlpha | TokenOmega | TokenLetter)
	if !ok || index != from+1 {
		return root.Parent, from, errs.Newf("could not complete variable from: '%s'", s.Tokens)
	}
	root.AppendAlpha()

	s.Cursor, valid = index, 1
	for {
		switch t := s.NextToken(); t.Kind {
		case TokenEOF:
			root.AppendOmega(false)
			return root, t.Pos + 1, nil

		case TokenAlpha:
			node, next, err := s.Parse(s.Cursor, root)
			if err != nil {
				if valid < 2 {
					return root.Parent, from, err
				}
				root.AppendOmega(false)
				return root, from + valid, nil
			}
			root.Children = append(root.Children, node)
			s.Cursor = next

		case TokenOmega:
			root.AppendOmega(true)
			return root, t.Pos + 1, nil

		case TokenSpace, TokenUnknown:
			root.AppendOmega(false)
			return root, t.Pos, nil

		case TokenNumber:
			if valid < 2 {
				return root.Parent, from, errs.Newf("could not complete variable from: '%s'", s.Tokens)
			}

			fallthrough
		case TokenLetter:
			root.AppendRunes(t.Rune)
			valid += 1
			s.Cursor++

		default:
			panic(fmt.Errorf("unexpected token: %#v", t))
		}
	}
}
