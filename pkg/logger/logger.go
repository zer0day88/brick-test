package logger

import (
	"github.com/zer0day88/brick-test/pkg/config"
	"github.com/zer0day88/brick-test/pkg/environment"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func shortCaller(pc uintptr, file string, line int) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			//break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}

func New() zerolog.Logger {

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var output io.Writer

	output = zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}

	logMinimumLevel := config.Key.LogLevel

	if config.Key.Environment == environment.Production {
		//force log minimum level to info on production
		logMinimumLevel = zerolog.WarnLevel
		output = os.Stderr

	}

	log := zerolog.New(output).
		With().
		Timestamp().Caller().Logger()

	log.Level(logMinimumLevel)

	return log

}
