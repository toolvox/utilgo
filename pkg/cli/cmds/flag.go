// Package cmds defines interfaces and types for flags and commands.
//
// cmds is somewhat compatible with go's [pkg/flag] package.
package cmds

// FlagValue is an interface to a value that can be set via flags.
// Compatible with go's [pkg/flag.Value] and [pkg/flag.Getter] interfaces.
//
// The result of calling [String] before [Set] is used to represent the default value.
//
// Boolean flags (which don't require an explicit value) must be marked by implementing:
//
//	func (*) IsBoolFlag() bool { return true }
//
// Definition:
type FlagValue interface {
	Set(string) error
	Get() any
	String() string
}

// isBoolFlag is the implicit interface that implements IsBoolFlag() bool.
type isBoolFlag interface {
	IsBoolFlag() bool
}

// Flag is an interface to flag with a name, a value, and optional usage text.
//
// Go's [pkg/flag.Flag] struct can be used by wrapping with [FromFlag].
type Flag interface {
	FlagValue

	FlagName() string
	FlagUsage() string
}

// CmdFlag implements [Flag] wrapping [FlagValue] with a name and usage text.
type CmdFlag struct {
	FlagValue

	Name  string
	Usage string
}

// FlagName returns the flag's name.
func (vf CmdFlag) FlagName() string { return vf.Name }

// FlagUsage returns the flag's usage text.
func (vf CmdFlag) FlagUsage() string { return vf.Usage }

// NewFlag creates a new [Flag] from a [FlagValue], a name, and usage text.
func NewFlag(value FlagValue, name string, usage string) Flag {
	return &CmdFlag{
		FlagValue: value,
		Name:      name,
		Usage:     usage,
	}
}
