package mariuscommon

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(name string) (*zap.Logger, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout", fmt.Sprintf("%s.log", name)},
		ErrorOutputPaths: []string{"stderr", fmt.Sprintf("%s.err", name)},
	}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
