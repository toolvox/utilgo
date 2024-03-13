package errors

import (
	"fmt"
)

func WrapCommaError[T any](val T, err error) func(fmtString string, args ...any) (T, error) {
	if err != nil {
		return func(fmtString string, args ...any) (T, error) {
			args = append(args, err)
			return val, fmt.Errorf(fmtString+": %w", args...)
		}
	}
	return func(_ string, _ ...any) (T, error) { return val, nil }
}

func MustCommaError[T any](val T, err error) func(fmtString string, args ...any) T {
	if err != nil {
		return func(fmtString string, args ...any) T {
			args = append(args, err)
			panic(fmt.Errorf(fmtString+": %w", args...))
		}
	}
	return func(_ string, _ ...any) T { return val }
}

func Wrap[T any](val T, err error, wrapMsgAndArgs ...any) (T, error) {
	if err != nil {
		if len(wrapMsgAndArgs) > 0 {
			wrapMsgAndArgs = append(wrapMsgAndArgs, err)
			err = fmt.Errorf(wrapMsgAndArgs[0].(string)+": %w", wrapMsgAndArgs[1:]...)
		}
		return val, err
	}
	return val, nil
}

func Must[T any](val T, err error, wrapMsgAndArgs ...any) T {
	if err != nil {
		if len(wrapMsgAndArgs) > 0 {
			wrapMsgAndArgs = append(wrapMsgAndArgs, err)
			err = fmt.Errorf(wrapMsgAndArgs[0].(string)+": %w", wrapMsgAndArgs[1:]...)
		}
		panic(err)
	}
	return val
}
