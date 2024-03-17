package cmds

// Cmdlet is an interface that has one function Run which takes a slice of action args and returns an [error].
//
// The actions are the program args with all -flags parsed and removed.
type Cmdlet interface {
	Run(actions []string) error
}

// CmdFunc wraps a function with the correct type to a [Cmdlet].
type CmdFunc func(args []string) error

// Run invokes the [CmdFunc].
func (c CmdFunc) Run(actions []string) error { return c(actions) }
