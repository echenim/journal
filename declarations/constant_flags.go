package declaration

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	Lprecedency
	LstdFlags = Ldate | Ltime
)
