package log

import (
	"io"
	"os"
	"time"

	nested "github.com/caikz/log/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type w struct{ a, b io.Writer }

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

	f, _ := os.Create(".log")
	logrus.SetOutput(w{f, os.Stdout})
}

var stdloger = Logger{logrus.StandardLogger()}
var T = stdloger.T
var D = stdloger.D
var I = stdloger.I
var W = stdloger.W
var E = stdloger.E
var F = stdloger.F
var P = stdloger.P

var Tf = stdloger.Tf
var Df = stdloger.Df
var If = stdloger.If
var Wf = stdloger.Wf
var Ef = stdloger.Ef
var Ff = stdloger.Ff
var Pf = stdloger.Pf

var Write = stdloger.Write
var Writef = stdloger.Writef
var WriteString = stdloger.WriteString
var Ln = stdloger.Ln

var AttrF256 = stdloger.AttrF256
var AttrB256 = stdloger.AttrB256
var AttrF24 = stdloger.AttrF24
var AttrB24 = stdloger.AttrB24
var AttrReset = stdloger.AttrReset

func SetT() { stdloger.SetLevel(logrus.TraceLevel) }
func SetD() { stdloger.SetLevel(logrus.DebugLevel) }
func SetI() { stdloger.SetLevel(logrus.InfoLevel) }
func SetW() { stdloger.SetLevel(logrus.WarnLevel) }
func SetE() { stdloger.SetLevel(logrus.ErrorLevel) }
func SetF() { stdloger.SetLevel(logrus.FatalLevel) }
func SetP() { stdloger.SetLevel(logrus.PanicLevel) }
