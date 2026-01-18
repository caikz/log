package log

import (
	"github.com/sirupsen/logrus"
)

func init() {

}

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

var endln = []byte{'\n'}

func (l *Logger) Ln() { l.Logger.Out.Write(endln) }

// func (l *Logger) Write(p []byte) (n int, err error) { return l.Logger.Out.Write(p) }
// func (l *Logger) Writef(f string, args ...any) (n int, err error) {
// 	return fmt.Fprintf(l.Logger.Out, f, args...)
// }
