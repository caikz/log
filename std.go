package log

import (
	"io"
	"os"
	"time"

	nested "github.com/caikz/log/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type w struct {
	a, b io.Writer
}

func (w w) Write(p []byte) (n int, err error) {
	if w.a != nil {
		n, err = w.a.Write(p)
	}
	if w.b != nil {
		n, err = w.b.Write(p)
	}
	return
}
func init() {
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.DateTime,
	})
	logrus.SetReportCaller(true)
	SetD()

	f, _ := os.Create("board.log")
	// Stdout := w{}
	// Stdout.a = os.Stdout
	// Stdout.b = f
	Stdout := w{
		f, os.Stdout,
	}
	logrus.SetOutput(Stdout)
}

var stdloger = Logger{logrus.StandardLogger()}
var T = stdloger.Trace
var D = stdloger.Debug
var I = stdloger.Info
var W = stdloger.Warn
var E = stdloger.Error
var F = stdloger.Fatal
var P = stdloger.Panic

var Tf = stdloger.Tracef
var Df = stdloger.Debugf
var If = stdloger.Infof
var Wf = stdloger.Warnf
var Ef = stdloger.Errorf
var Ff = stdloger.Fatalf
var Pf = stdloger.Panicf

var Ln = stdloger.Ln

func SetT() { stdloger.SetLevel(logrus.TraceLevel) }
func SetD() { stdloger.SetLevel(logrus.DebugLevel) }
func SetI() { stdloger.SetLevel(logrus.InfoLevel) }
func SetW() { stdloger.SetLevel(logrus.WarnLevel) }
func SetE() { stdloger.SetLevel(logrus.ErrorLevel) }
func SetF() { stdloger.SetLevel(logrus.FatalLevel) }
func SetP() { stdloger.SetLevel(logrus.PanicLevel) }
