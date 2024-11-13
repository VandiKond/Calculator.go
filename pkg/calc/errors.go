package calc

import "github.com/VandiKond/vanerrors"

var ShowCauseOptions vanerrors.Options = vanerrors.Options{
	ShowMessage: true,
	ShowCause:   true,
}

var DefaultOptions vanerrors.Options = vanerrors.Options{
	ShowMessage: true,
}

func GetDefaultLogSetting() vanerrors.LoggerOptions {
	options := vanerrors.DefaultLoggerOptions
	options.DoLog = false
	return options
}

// Error names
const (
	DBZ = "divide by zero not allowed"
	US  = "unknown symbol"
	NPE = "num parsing error"
	EDO = "error doing operation"
	COO = "error completing order operation"
	ECE = "error completing the expression"
	BSC = "bracket should be closed"
	GRB = "error in getting rid of brackets"
)
