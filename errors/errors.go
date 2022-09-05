package errors

import (
	"fmt"
)

type colorCode struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	White  string
}

var colors = colorCode{
	Reset:  "\033[0m",
	Red:    "\033[31m",
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Blue:   "\033[34m",
	Purple: "\033[35m",
	Cyan:   "\033[36m",
	White:  "\033[37m",
}

/*
 *
 *
 */

/* Custom Error ------------------------------------------------------------------------------------------------- */

type Error struct {
	message string
	info    string
}

func (err *Error) Error() string {
	return err.message
}

/* Custom Error ------------------------------------------------------------------------------------------------- */

/*
 *
 *
 */

/* Error Handlers ------------------------------------------------------------------------------------------------- */

func ErrorHandler(err error) {
	errorHandler(err, "", nil)
}

func ErrorHandlerWithInfo(err error, info string) {
	errorHandler(err, info, nil)
}

func CustomErrorHandler(err error, cb func(err *string)) {
	errorHandler(err, "", cb)
}

func CapturePanic(printError bool, cb func(a *interface{})) {
	if err := recover(); err != nil {
		if printError {
			checkAndPrint(err)
		}

		if cb != nil {
			cb(&err)
		}
	}
}

/* Error Handlers ------------------------------------------------------------------------------------------------- */

/*
 *
 *
 */

/* Local functions ------------------------------------------------------------------------------------------------- */

func errorHandler(err error, info string, cb func(err *string)) {
	if err == nil {
		return
	}

	errMessage := err.Error()

	if cb == nil {
		panic(&Error{message: errMessage, info: info})
	}

	cb(&errMessage)
}

func checkAndPrint(err interface{}) {
	switch e := err.(type) {
	case *Error:
		fmt.Printf("%s[PANIC]%s (%v) %s %v\n", colors.Red, colors.Blue, e.info, colors.Reset, e.message)
	default:
		fmt.Printf("%s[PANIC]%s %+v \n", colors.Red, colors.Reset, err)
	}
}

/* Local functions ------------------------------------------------------------------------------------------------- */
