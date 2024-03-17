package lexer

import (
	"strings"

	"utilgo/pkg/sets"
	"utilgo/pkg/stringutil"
)

// FlagTokenizer holds the names of the boolean flags and uses them the tokenize the input args into args and actions.
// Actions are args that are not flags.
//
// After calling [FlagTokenizer.TokenizeFlags] the Tokenized and Action fields will be initialized.
type FlagTokenizer struct {
	BooleanFlags sets.TinySet[string]
	Tokenized    map[string]string
	Actions      []string
}

// NewFlagTokenizer creates a new [FlagTokenizer] from a list of boolean flag names
func NewFlagTokenizer(booleanFlags ...string) *FlagTokenizer {
	return &FlagTokenizer{
		BooleanFlags: sets.NewTinySet(booleanFlags...),
		Tokenized:    map[string]string{},
		Actions:      []string{},
	}
}

// TokenizeFlags sorts the args into flags and actions, without parsing the values.
func (ft *FlagTokenizer) TokenizeFlags(args []string) {
	ft.Actions = append(ft.Actions, args[0])
	for i := 1; i < len(args); i++ {
		arg := strings.TrimSpace(args[i])
		if len(arg) == 0 {
			continue
		}

		// action
		prefixLen := stringutil.CountPrefix(arg, func(r rune) bool { return r == '-' })
		if prefixLen == 0 {
			ft.Actions = append(ft.Actions, arg)
			continue
		}

		arg = arg[prefixLen:]
		if len(arg) == 0 {
			continue
		}

		// -flag=
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			arg = parts[0]
			if len(arg) == 0 {
				continue
			}

			val := parts[1]
			if len(val) == 0 {
				if ft.BooleanFlags.Contains(arg) {
					ft.Tokenized[arg] = "false"
					continue
				}

				ft.Tokenized[arg] = ""
				continue
			}

			ft.Tokenized[arg] = val
			continue
		}

		// -flag
		if ft.BooleanFlags.Contains(arg) {
			ft.Tokenized[arg] = "true"
			continue
		}

		if i+1 == len(args) {
			ft.Tokenized[arg] = ""
			continue
		}

		val := args[i+1]
		if strings.HasPrefix(val, "-") {
			ft.Tokenized[arg] = ""
			continue
		}

		ft.Tokenized[arg] = val
		i++
	}

	for _, v := range ft.BooleanFlags {
		if _, ok := ft.Tokenized[v]; !ok {
			ft.Tokenized[v] = "false"
		}
	}
}
