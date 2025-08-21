package output

import (
	"fmt"
	"os"
	"regexp"

	"github.com/pterm/pterm"
)

// ----- Verbosity Level Enum ----- //

type VerbosityLevel int

const (
	undefined VerbosityLevel = iota
	Normal
	Verbose
	Detailed
)

// ----- Variables ----- //

var verbosityThreshold VerbosityLevel
var highlightedStyle *pterm.Style = pterm.NewStyle(pterm.BgYellow, pterm.FgBlack, pterm.Bold)

// ----- Public Functions ----- //

func SetThreshold(count int) {
	switch count {
	case 0:
		verbosityThreshold = Normal
	case 1:
		verbosityThreshold = Verbose
	default:
		verbosityThreshold = Detailed
		pterm.EnableDebugMessages()
	}
}

func Info(level VerbosityLevel, msg string, args ...any) {
	if level > verbosityThreshold {
		return
	}

	format, styledArgs := highlightArgs(msg, args...)

	pterm.Info.Println(fmt.Sprintf(format, styledArgs...))
}

func Success(level VerbosityLevel, msg string, args ...any) {
	if level > verbosityThreshold {
		return
	}

	format, styledArgs := highlightArgs(msg, args...)

	pterm.Success.Println(fmt.Sprintf(format, styledArgs...))
}

func Warn(level VerbosityLevel, msg string, args ...any) {
	if level > verbosityThreshold {
		return
	}

	format, styledArgs := highlightArgs(msg, args...)

	pterm.Warning.Println(fmt.Sprintf(format, styledArgs...))
}

func Error(msg string, args ...any) {
	pterm.Error.Println(fmt.Sprintf(msg, args...))
	os.Exit(1)
}

func Debug(msg string, args ...any) {
	format, styledArgs := highlightArgs(msg, args...)

	pterm.Debug.Println(fmt.Sprintf(format, styledArgs...))
}

// ----- Private Functions ----- //

func init() {
	pterm.Info.Prefix = pterm.Prefix{Text: " INFO  ", Style: pterm.NewStyle(pterm.BgBlue, pterm.FgBlack)}
	pterm.Info.MessageStyle = pterm.NewStyle(pterm.FgBlue)

	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		disableColor()
	}
}

func highlightArgs(msg string, args ...any) (string, []any) {
	// Apply the style to each argument
	styledArgs := make([]any, len(args))
	for i, arg := range args {
		styledArgs[i] = highlightedStyle.Sprint(arg)
	}

	// As all args are now strings replace format to match
	regEx := regexp.MustCompile(`%[a-zA-Z]`)
	format := regEx.ReplaceAllString(msg, "%s")

	return format, styledArgs
}

func disableColor() {
	pterm.DisableColor()
	pterm.Info.Prefix = pterm.Prefix{Text: "[*]"}
	pterm.Success.Prefix = pterm.Prefix{Text: "[*]"}
	pterm.Warning.Prefix = pterm.Prefix{Text: "[!]"}
	pterm.Error.Prefix = pterm.Prefix{Text: "[!]"}
	pterm.Debug.Prefix = pterm.Prefix{Text: "[DEBUG]"}

}
