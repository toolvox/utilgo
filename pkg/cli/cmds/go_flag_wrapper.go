package cmds

import (
	"flag"

	"github.com/toolvox/utilgo/pkg/errs"
)

// GoFlag wraps a go [pkg/flag.Flag] making it compatible with this package.
type GoFlag struct {
	*flag.Flag
}

// Set implements [Flag] by calling Set on the [pkg/flag.Value] field of [pkg/flag.Flag].
func (gf GoFlag) Set(value string) error {
	if gf.Value == nil {
		return errs.Newf("flag '%s' value was nil", gf.Name)
	}
	return gf.Value.Set(value)
}

// Get implements [Flag] by calling Get on the [pkg/flag.Value] field of [pkg/flag.Flag].
func (gf GoFlag) Get() any {
	if gf.Value == nil {
		return errs.Newf("flag '%s' value was nil", gf.Name)
	}

	if gfGetter, ok := gf.Value.(flag.Getter); ok {
		return gfGetter.Get()
	}

	return errs.Newf("flag '%s' value has no Get() any function", gf.Name)
}

// String implements [Flag] by returning the [pkg/flag.Flag]'s DefValue.
func (gf GoFlag) String() string {
	return gf.DefValue
}

// IsBoolFlag returns true of the wrapped flag is a bool flag.
func (gf GoFlag) IsBoolFlag() bool {
	bf, ok := gf.Value.(isBoolFlag)
	return ok && bf.IsBoolFlag()
}

// FlagName returns the name of the [Flag].
func (gf GoFlag) FlagName() string { return gf.Name }

// FlagUsage returns the usage text of the [Flag].
func (gf GoFlag) FlagUsage() string { return gf.Usage }

// FromFlag creates a new [Flag] by wrapping a [pkg/flag.Flag].
func FromFlag(goFlag *flag.Flag) Flag {

	return GoFlag{goFlag}
}
