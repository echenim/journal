package declaration

const (
	Ldate         = 1 << iota     // the date: 2012/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	Lprecedency                     // the priority: Debug
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
