package plugs

import (
	"fmt"
	"io"
	"log"
	"strings"

	d "github.com/echenim/journal/declarations"
)

// Logger defines our wrapper around the system logger
type Logger struct {
	Precedency int
	Before     string
	logger     *log.Logger
}

// New creates a new Logger.
func New(out io.Writer, before string, flag int, precedency int) *Logger {
	return &Logger{
		Precedency: precedency,
		Before:     before,
		logger:     log.New(out, before, flag),
	}
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(before string) {
	l.Before = before
	l.logger.SetPrefix(before)
}

// Prefix returns the current logger prefix
func (l *Logger) Prefix() string {
	return l.Before
}

func (l *Logger) setFullPrefix(precedency int) {
	if l.logger.Flags()&d.Lprecedency != 0 {
		l.logger.SetPrefix(fmt.Sprintf("%s ", d.PrecencyName[precedency]) + l.Before)
	}
}

// Calls Output to print to the logger.
func (l *Logger) print(precedency int, v ...interface{}) {
	if precedency <= l.Precedency {
		l.setFullPrefix(precedency)
		l.logger.Print(v...)
	}
}

// Calls Output to printf to the logger.
func (l *Logger) printf(precedency int, format string, v ...interface{}) {
	if precedency <= l.Precedency {
		l.setFullPrefix(precedency)
		l.logger.Printf(format, v...)
	}
}

// Calls Output to println to the logger.
func (l *Logger) println(precedency int, v ...interface{}) {
	if precedency <= l.Precedency {
		l.setFullPrefix(precedency)
		l.logger.Println(v...)
	}
}

// Priority returns the output priority for the logger.
func (l *Logger) Priority() int {
	return l.Precedency
}

// SetPriority sets the output priority for the logger.
func (l *Logger) SetPriority(precedency int) {
	l.Precedency = precedency
}

// SetPriorityString sets the output priority by the name of a debug level
func (l *Logger) SetPriorityString(s string) error {
	s = strings.ToUpper(s)
	for i, name := range d.PrecencyName {
		if name == s {
			l.SetPriority(i)
			return nil
		}
	}
	return fmt.Errorf("Unable to find priority %s", s)
}

// Flags returns the output layouts for the logger.
func (l *Logger) Flags() int {
	return l.logger.Flags()
}

// SetFlags sets the output layouts for the logger.
func (l *Logger) SetFlags(layouts int) {
	l.logger.SetFlags(layouts)
}

// Fatal prints the message it's given and quits the program
func (l *Logger) Fatal(v ...interface{}) {
	l.setFullPrefix(d.Fatal)
	l.logger.Fatal(v...)
}

// Fatalf prints the message it's given and quits the program
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.setFullPrefix(d.Fatal)
	l.logger.Fatalf(format, v...)
}

// Fatalln prints the message it's given and quits the program
func (l *Logger) Fatalln(v ...interface{}) {
	l.setFullPrefix(d.Fatal)
	l.logger.Fatalln(v...)
}

// Panic prints the message it's given and panic()s the program
func (l *Logger) Panic(v ...interface{}) {
	l.setFullPrefix(d.Fatal)
	l.logger.Panic(v...)
}

// Panicf prints the message it's given and panic()s the program
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.setFullPrefix(d.Fatal)
	l.logger.Panicf(format, v...)
}

// Panicln prints the message it's given and panic()s the program
func (l *Logger) Panicln(v ...interface{}) {
	l.setFullPrefix(d.Fatal)
	l.logger.Panicln(v...)
}

// Error prints to the standard logger with the Error level.
func (l *Logger) Error(v ...interface{}) {
	l.print(d.Error, v...)
}

// Errorf prints to the standard logger with the Error level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(d.Error, format, v...)
}

// Errorln prints to the standard logger with the Error level.
func (l *Logger) Errorln(v ...interface{}) {
	l.println(d.Error, v...)
}

// Warn prints to the standard logger with the Warn level.
func (l *Logger) Warn(v ...interface{}) {
	l.print(d.Warn, v...)
}

// Warnf prints to the standard logger with the Warn level.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.printf(d.Warn, format, v...)
}

// Warnln prints to the standard logger with the Warn level.
func (l *Logger) Warnln(v ...interface{}) {
	l.println(d.Warn, v...)
}

// Info prints to the standard logger with the Info level.
func (l *Logger) Info(v ...interface{}) {
	l.print(d.Info, v...)
}

// Infof prints to the standard logger with the Info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(d.Info, format, v...)
}

// Infoln prints to the standard logger with the Info level.
func (l *Logger) Infoln(v ...interface{}) {
	l.println(d.Info, v...)
}

// Debug prints to the standard logger with the Debug level.
func (l *Logger) Debug(v ...interface{}) {
	l.print(d.Debug, v...)
}

// Debugf prints to the standard logger with the Debug level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(d.Debug, format, v...)
}

// Debugln prints to the standard logger with the Debug level.
func (l *Logger) Debugln(v ...interface{}) {
	l.println(d.Debug, v...)
}

// Trace prints to the standard logger with the Trace level.
func (l *Logger) Trace(v ...interface{}) {
	l.print(d.Trace, v...)
}

// Tracef prints to the standard logger with the Trace level.
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.printf(d.Trace, format, v...)
}

// Traceln prints to the standard logger with the Trace level.
func (l *Logger) Traceln(v ...interface{}) {
	l.println(d.Trace, v...)
}
