package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/cortezaproject/corteza-server/pkg/options"
)

var (
	DefaultLevel  = zap.NewAtomicLevel()
	defaultLogger = zap.NewNop()
)

func MakeDebugLogger() *zap.Logger {
	conf := zap.NewDevelopmentConfig()
	conf.Level = DefaultLevel
	conf.Level.SetLevel(zap.DebugLevel)

	// Print log level in colors
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Shorten timestamp, we do not care about the date
	conf.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("15:04:05.000"))
	}

	logger, err := conf.Build()
	if err != nil {
		panic(err)
	}

	logger.Debug("full debug mode enabled (LOG_DEBUG=true), enjoy all the colors and verbosity")

	return logger
}

func Init() {
	var (
		err error

		// Set WARN as default log level -- running serve-api
		// will override this if LOG_LEVEL is not set
		logLevel = zapcore.WarnLevel

		// Do we want to enable debug logger
		// with a bit more dev-friendly output
		debuggingLogger = options.EnvBool("LOG_DEBUG", false)
	)

	if ll, has := os.LookupEnv("LOG_LEVEL"); has {
		_ = logLevel.Set(ll)
	}

	DefaultLevel.SetLevel(logLevel)

	if debuggingLogger {
		defaultLogger = MakeDebugLogger()
		return
	}

	conf := zap.NewProductionConfig()
	conf.Level = DefaultLevel

	// We do not want sampling
	conf.Sampling = nil

	defaultLogger, err = conf.Build()
	if err != nil {
		panic(err)
	}
}

func Default() *zap.Logger {
	return defaultLogger
}

func SetDefault(logger *zap.Logger) {
	if logger == nil {
		logger = zap.NewNop()
	}

	defaultLogger = logger
}
