package output

import "fmt"

func Info(msg string, args ...any) {
	fmt.Printf("[*] "+msg+"\n", args...)
}
