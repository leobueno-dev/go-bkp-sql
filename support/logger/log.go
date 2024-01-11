package logger

import (
	"os"
	"strings"

	"github.com/leobueno-dev/go-bkp-sql/support"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	//log.Logger = buildBaseLogger(log.Logger, "").With().Fields(nil).Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldInteger = true
}

func SetupLogger(level string) {

	zerolog.SetGlobalLevel(getLogLevel(level))

	log.Logger = buildBaseLogger(log.Logger, "").
		With().
		Logger()
}

func Logger(process string, tags map[string]interface{}) zerolog.Logger {
	builder := log.Logger.
		With().
		Str("process", process)

	return builder.Fields(tags).Logger()
}

func getLogLevel(val string) zerolog.Level {
	if val == "" {
		val = support.GetEnv("LOG_LEVEL", "info")
	}

	level := strings.ToLower(val)

	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func buildBaseLogger(l zerolog.Logger, format string) zerolog.Logger {
	if format == "" {
		format = strings.ToLower(support.GetEnv("LOG_FORMAT", "text"))
	}

	switch format {
	case "json":
		return l
	default:
		return l.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
