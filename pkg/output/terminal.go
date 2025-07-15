package output

import (
	"fmt"
	"os"
	"regexp"

	"github.com/pterm/pterm"
)

var highlightedStyle *pterm.Style

func init() {
	highlightedStyle = pterm.NewStyle(pterm.BgYellow, pterm.FgBlack, pterm.Bold)

	pterm.Info.Prefix = pterm.Prefix{Text: " INFO  ", Style: pterm.NewStyle(pterm.BgBlue, pterm.FgBlack)}
	pterm.Info.MessageStyle = pterm.NewStyle(pterm.FgBlue)

	pterm.EnableDebugMessages()
}

func Info(msg string, args ...any) {
	format, styledArgs := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Info.Println(fmt.Sprintf(format, styledArgs...))
}

func Success(msg string, args ...any) {
	format, styledArgs := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Success.Println(fmt.Sprintf(format, styledArgs...))
}

func Warn(msg string, args ...any) {
	format, styledArgs := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Warning.Println(fmt.Sprintf(format, styledArgs...))
}

func Error(msg string) {
	pterm.Error.Println(msg)
	os.Exit(1)
}

func Debug(msg string, args ...any) {
	format, styledArgs := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Debug.Println(fmt.Sprintf(format, styledArgs...))
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
