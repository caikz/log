package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func (l *Logger) T(args ...any)                 { l.Trace(args...) }
func (l *Logger) D(args ...any)                 { l.Debug(args...) }
func (l *Logger) I(args ...any)                 { l.Info(args...) }
func (l *Logger) W(args ...any)                 { l.Warn(args...) }
func (l *Logger) E(args ...any)                 { l.Error(args...) }
func (l *Logger) F(args ...any)                 { l.Fatal(args...) }
func (l *Logger) P(args ...any)                 { l.Panic(args...) }
func (l *Logger) Tf(format string, args ...any) { l.Tracef(format, args...) }
func (l *Logger) Df(format string, args ...any) { l.Debugf(format, args...) }
func (l *Logger) If(format string, args ...any) { l.Infof(format, args...) }
func (l *Logger) Wf(format string, args ...any) { l.Warnf(format, args...) }
func (l *Logger) Ef(format string, args ...any) { l.Errorf(format, args...) }
func (l *Logger) Ff(format string, args ...any) { l.Fatalf(format, args...) }
func (l *Logger) Pf(format string, args ...any) { l.Panicf(format, args...) }

func (l *Logger) Write(p []byte) (n int, err error) { return l.Out.Write(p) }
func (l *Logger) Writef(f string, args ...any) (n int, err error) {
	return fmt.Fprintf(l.Out, f, args...)
}
func (l *Logger) WriteString(s string) (n int, err error) {
	return l.Out.Write([]byte(s))
}
func (l *Logger) Ln() { l.Out.Write([]byte("\n")) }

func (l *Logger) AttrF256(c uint8) { l.Writef("\033[38;5;%dm", c) }
func (l *Logger) AttrB256(c uint8) { l.Writef("\033[48;5;%dm", c) }
func (l *Logger) AttrF24(c uint32) {
	l.Writef("\033[38;2;%d;%d;%dm", (c>>16)&0xFF, (c>>8)&0xFF, c&0xFF)
}
func (l *Logger) AttrB24(c uint32) {
	l.Writef("\033[48;2;%d;%d;%dm", (c>>16)&0xFF, (c>>8)&0xFF, c&0xFF)
}
func (l *Logger) AttrReset() { l.Out.Write([]byte("\033[0m")) }
