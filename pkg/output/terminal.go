package output

import (
	"fmt"
	"regexp"

	"github.com/pterm/pterm"
)

func init() {
	pterm.Info.Prefix = pterm.Prefix{Text: "INFO", Style: pterm.NewStyle(pterm.BgBlue, pterm.FgBlack)}
	pterm.Info.MessageStyle = pterm.NewStyle(pterm.FgBlue)
}

func Info(msg string, args ...any) {
	highlightedMsg := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Info.Println(highlightedMsg)
}

func Success(msg string, args ...any) {
	highlightedMsg := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Info.Println(highlightedMsg)
}

func Warn(msg string, args ...any) {
	highlightedMsg := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Warning.Println(highlightedMsg)
}

func Error(msg string, args ...any) {
	highlightedMsg := highlightArgs(msg, args...)

	// Print the final formatted message with styled args
	pterm.Error.Println(highlightedMsg)
}

func highlightArgs(msg string, args ...any) (highlightedMsg string) {
	style := pterm.NewStyle(pterm.BgYellow, pterm.FgBlack, pterm.Bold)

	// Apply the style to each argument
	styledArgs := make([]any, len(args))
	for i, arg := range args {
		styledArgs[i] = style.Sprint(arg)
	}

	// As all args are now strings replace format to match
	regEx := regexp.MustCompile(`%[a-zA-Z]`)
	format := regEx.ReplaceAllString(msg, "%s")

	highlightedMsg = fmt.Sprintf(format, args...)

	return
}
