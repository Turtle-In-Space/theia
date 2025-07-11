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

	// As all args are now strings replace format to match
	regEx := regexp.MustCompile(`%[a-zA-Z]`)
	format := regEx.ReplaceAllString(msg, "%s")

	// Print the final formatted message with styled args
	pterm.Info.Printfln(format, styledArgs...)
}
