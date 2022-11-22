package crawler

type result struct {
	Address  string
	Checksum string
}

type failure struct {
	Address string
	Error   error
}
