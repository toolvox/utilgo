package tmplz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/toolvox/utilgo/pkg/stringutil"
)

type String = stringutil.String

const (
	ALPHA_RUNE = '@'
	ALPHA      = "@"

	OMEGA_RUNE = '_'
	OMEGA      = "_"

	OMICRON_RUNE = '‗'
	OMICRON      = "‗"
)

// TokenKind represents the type of lexical token identified in template strings.
type TokenKind uint

const (
	TokenEOF   TokenKind = 0
	TokenAlpha TokenKind = 1 << (iota - 1)
	TokenOmega
	TokenNumber
	TokenLetter
	TokenSpace
	TokenUnknown
)

var tokenNames = map[TokenKind]string{
	TokenEOF:     "EOF",
	TokenAlpha:   "Alpha",
	TokenOmega:   "Omega",
	TokenNumber:  "Number",
	TokenLetter:  "Letter",
	TokenSpace:   "Space",
	TokenUnknown: "Unknown",
}

// Token defines a lexical token with its kind, rune, and position in the input string.
type Token struct {
	Kind TokenKind
	Rune rune
	Pos  int
}

// String returns a string representation of the Token, including its rune and position.
func (t Token) String() string {
	return fmt.Sprintf("'%s'[%d]", string(t.Rune), t.Pos)
}

// Tokens is a slice of Token instances, representing a sequence of tokens lexed from a template string.
type Tokens []Token

// String returns a concatenated string representation of all Tokens in the slice.
func (ts Tokens) String() string {
	var sb strings.Builder
	for _, t := range ts {
		sb.WriteRune(t.Rune)
	}
	return sb.String()
}

// Lex performs lexical analysis on the input string, identifying and categorizing each rune as a token based on its type (e.g., letter, number, special rune).
// Returns a slice of Tokens representing the lexed input.
func Lex(s String) (result Tokens) {
	for i, r := range s {
		newToken := Token{Rune: r, Pos: i}

		switch {
		case unicode.IsSpace(r):
			newToken.Kind = TokenSpace
		case unicode.IsLetter(r):
			newToken.Kind = TokenLetter
		case unicode.IsNumber(r):
			newToken.Kind = TokenNumber
		case r == ALPHA_RUNE:
			newToken.Kind = TokenAlpha
		case r == OMEGA_RUNE:
			newToken.Kind = TokenOmega
		default:
			newToken.Kind = TokenUnknown
		}

		result = append(result, newToken)
	}

	return result
}
