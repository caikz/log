package log

import (
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	nested "github.com/caikz/log/nested-logrus-formatter"

	"github.com/sirupsen/logrus"
)

var Colors = []int{
	208,                          //red,orange
	211, 220, 223, 225, 228, 230, //yellow
	86, 121, 153, //green
	117, 111, 123, 189, //skyblue
	195,                //light
	175, 217, 218, 219, //pink
	177, 183, 213, //purple
} //这啥呀(_^_)_这是
func br() string {
	return fmt.Sprintf("\033[48;5;%[1]dm{\033[0m%%v\033[48;5;%[1]dm}\033[0m", Colors[rand.IntN(len(Colors))])
}
func U(a ...any) string {
	var b strings.Builder
	n := len(a)
	b.Grow(n * 8)
	for range n {
		b.WriteString(br())
	}
	return fmt.Sprintf(b.String(), a...)
}

var D = logrus.Debug
var Df = logrus.Debugf

var P = logrus.Panic
var Pf = logrus.Panicf

var E = logrus.Error
var Ef = logrus.Errorf

var I = logrus.Info
var If = logrus.Infof

var W = logrus.Warn
var Wf = logrus.Warnf

var F = logrus.Fatal
var Ff = logrus.Fatalf

var T = logrus.Trace
var Tf = logrus.Tracef

func ST() {
	logrus.SetLevel(logrus.TraceLevel)
}
func SD() {
	logrus.SetLevel(logrus.DebugLevel)
}
func SI() {
	logrus.SetLevel(logrus.InfoLevel)
}
func SW() {
	logrus.SetLevel(logrus.WarnLevel)
}
func SE() {
	logrus.SetLevel(logrus.ErrorLevel)
}
func SF() {
	logrus.SetLevel(logrus.FatalLevel)
}
func SP() {
	logrus.SetLevel(logrus.PanicLevel)
}
func N(n int) {
	for range n {
		Wr.Write([]byte{'\n'})
	}
}

func Et[T any](v T, err error) {
	if err != nil {
		Ef("{%v}{%v}", v, err)
	} else {
		Df("{%v}", v)
	}
}

// func Dl(a ...any) {
// logrus.SetReportCaller(true)
// D(a...)
// logrus.SetReportCaller(false)
// }

var Wr = w{}

func init() {
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.DateTime,
	})
	logrus.SetReportCaller(true)

	logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetLevel(logrus.PanicLevel)

	Wr.a = os.Stdout
	f, _ := os.Create("board.log")
	Wr.b = f
	Wr = w{
		f, os.Stdout,
	}
	logrus.SetOutput(Wr)
}

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
