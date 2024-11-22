package calc

import "github.com/VandiKond/vanerrors"

var DefaultOptions vanerrors.Options = vanerrors.Options{
	ShowMessage: true,
	ShowCause:   true,
}

var DefaultLoggerOptions vanerrors.LoggerOptions = vanerrors.LoggerOptions{
	ShowMessage: true,
	ShowCause:   true,
}

// Error names
const (
	ErrorDivideByZero             = "divide by zero not allowed"
	ErrorUnknownOperator          = "unknown operator"
	ErrorParsingNumber            = "number parsing error"
	ErrorDoingOperation           = "error doing operation"
	ErrorCompletingOrderOperation = "error completing order operation"
	ErrorExpressionCompleting     = "error completing the expression"
	ErrorBracketShouldBeClosed    = "bracket should be closed"
	ErrorBracketOf                = "error getting rid of brackets"
	ErrorBracketShouldBeOpened    = "bracket should be opened"
)

func DefaultCalcVanError(Name string, Message string, Cause error) vanerrors.VanError {
	err := errorW.NewWrap(Name, Cause, nil)
	err.Message = Message
	return err
}

var errorW = vanerrors.NewW(DefaultOptions, DefaultLoggerOptions)
