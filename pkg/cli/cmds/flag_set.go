package cmds

import (
	"fmt"
	"os"

	"utilgo/pkg/cli/lexer"
	"utilgo/pkg/errs"
	"utilgo/pkg/sliceutil"
)

// FlagSet is a collection of [Flag]s.
type FlagSet []Flag

// Flag adds configured [Flag] to the set.
func (fs *FlagSet) Flag(flag Flag) FlagValue {
	*fs = append(*fs, flag)
	return flag
}

// Var creates a new [Flag] and adds it to the set.
func (fs *FlagSet) Var(value FlagValue, name string, usage string) FlagValue {
	flag := NewFlag(value, name, usage)
	*fs = append(*fs, flag)
	return flag
}

// BoolFlags returns the name of all the [Flag]s which are boolean flags.
func (fs *FlagSet) BoolFlags() []string {
	return sliceutil.SelectNonZeroFunc(*fs,
		func(flag Flag) string {
			if boolFlag, ok := flag.(isBoolFlag); ok && boolFlag.IsBoolFlag() {
				return flag.FlagName()
			}
			return ""
		},
	)
}

// Parse parses the [Flag]s and actions in the [FlagSet], leaving the actions in [pkg/os.Args].
func (fs FlagSet) Parse(args []string) error {
	lex := lexer.NewFlagTokenizer(fs.BoolFlags()...)
	lex.TokenizeFlags(args)

	var errors errs.Errors
	for _, flag := range fs {
		if value, ok := lex.Tokenized[flag.FlagName()]; ok {
			if err := flag.Set(value); err != nil {
				errors.WithErrorf("flag '%s' set, error: %w", flag.FlagName(), err)
			}
		}
	}

	if err := errors.OrNil(); err != nil {
		return fmt.Errorf("flags for command: %w", err)
	}

	os.Args = lex.Actions
	return nil
}
