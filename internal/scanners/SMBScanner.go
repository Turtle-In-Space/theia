package scanners

import (
	out "github.com/Turtle-In-Space/theia/pkg/output"
)

type SMBScanner struct{}

func (s SMBScanner) Run() {
	out.Info("SMBScanner")
}

func init() {
	Register("smb", SMBScanner{})
}
