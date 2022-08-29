package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Log() zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(config)
	defaultConsoleLogLevel := zap.InfoLevel

	// uncomment for logging to file
	//fileEncoder := zapcore.NewConsoleEncoder(config)
	//defaultFileLogLevel := zap.DebugLevel
	//logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//writer := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultConsoleLogLevel),

		// uncomment for logging to file
		//zapcore.NewCore(fileEncoder, writer, defaultFileLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return *logger.Sugar()
}
