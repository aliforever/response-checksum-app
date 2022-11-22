package stdlogger

import "fmt"

// StdLogger logs everything to stdout
type StdLogger struct{}

// Println calls fmt.Println to print passed arguments
func (StdLogger) Println(args ...any) {
	fmt.Println(args...)
}

// Errorf calls fmt.Errorf with format and args
// Error is ignored for the sake of saving time for this test
func (StdLogger) Errorf(format string, args ...any) {
	_ = fmt.Errorf(format, args...)
}
