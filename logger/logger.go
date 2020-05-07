package logger

import (
	"os"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates zap sugared logger
func NewLogger() *zap.SugaredLogger {
	e := struct {
		IsDev bool `env:"DEV_MODE" envDefault:"false"`
	}{}
	if err := env.Parse(&e); err != nil {
		return nil
	}

	loggingLevel := zap.NewAtomicLevel()
	if e.IsDev {
		loggingLevel.SetLevel(zap.DebugLevel)
	}
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		loggingLevel,
	)).Sugar()

	return logger
}
