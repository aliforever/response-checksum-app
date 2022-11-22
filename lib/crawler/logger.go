package crawler

type Logger interface {
	Println(a ...any)
	Errorf(format string, a ...any)
}
