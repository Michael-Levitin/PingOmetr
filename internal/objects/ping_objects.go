package objects

type PingUser struct {
	Msec  int32
	Site  string
	Error string
}

type PingAdmin struct {
	Slowest  int64
	Fastest  int64
	Specific int64
}
