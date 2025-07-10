package output

import (
	"github.com/pterm/pterm"
)

func Info(msg string, args ...any) {
	pterm.Info.Prefix = pterm.Prefix{Text: "INFO", Style: pterm.NewStyle(pterm.BgBlue, pterm.FgBlack)}
	pterm.Info.MessageStyle = pterm.NewStyle(pterm.FgBlue)

	style := pterm.NewStyle(pterm.BgYellow, pterm.FgBlack, pterm.Bold)

	// Apply the style to each argument
	styledArgs := make([]any, len(args))
	for i, arg := range args {
		styledArgs[i] = style.Sprint(arg)
	}

	// Print the final formatted message with styled args
	pterm.Info.Printfln(msg, styledArgs...)
}
