package objects

type PingUser struct {
	Msec  int32
	Site  string
	Error error
}

type PingAdmin struct {
	Slowest  int32
	Fastest  int32
	Specific int32
}
