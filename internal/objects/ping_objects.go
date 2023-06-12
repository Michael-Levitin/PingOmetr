package objects

type PingUser struct {
	Msec  int32
	Site  string
	Error error
}

type PingAdmin struct {
	Min      int32
	Max      int32
	Specific int32
}
