package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

const TimeFormat = "2006-01-02 15:04:05"

func New() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out: os.Stdout,
		FormatTimestamp: func(i interface{}) string {
			parse, _ := time.Parse(time.RFC3339, i.(string))
			return parse.Format(TimeFormat)
		},
	}
	return zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(2).Logger()
}
