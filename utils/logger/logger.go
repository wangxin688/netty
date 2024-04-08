package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type ColorCode int

const (
	colorRed    ColorCode = 31
	colorYellow ColorCode = 33
	colorBlue   ColorCode = 36
	colorGray   ColorCode = 37
)

var logger = logrus.New()

type formatter struct {
	prefix string
}

func (f *formatter) Format(entry *logrus.Entry, withColor bool) ([]byte, error) {

	levelColor := getColorByLevel(entry.Level)
	level := strings.ToUpper(entry.Level.String())
	
	sb := &bytes.Buffer{}

	sb.WriteString(entry.Time.Format(time.RFC3339))
	if withColor {
		fmt.Fprintf(sb, "\x1b[%dm", levelColor)
	}
	sb.WriteString(" | ")
	sb.WriteString(level)
	sb.WriteString(" | ")
	sb.WriteString(f.prefix)
	sb.WriteString(entry.Message)
	return sb.Bytes(), nil
}

func setLogLevel(level logrus.Level) {
	logger.Level = level
}

func init() {
	logger.Level = logrus.InfoLevel
	logger.SetFormatter(&formatter{})
	logger.SetReportCaller(true)
}

func getColorByLevel(level logrus.Level) ColorCode {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return colorGray
	case logrus.InfoLevel:
		return colorYellow
	case logrus.WarnLevel:
		return colorRed
	case logrus.ErrorLevel:
		return colorRed
	default:
		return colorBlue
	}
}
