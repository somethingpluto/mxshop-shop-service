package initialize

import (
	"Shop_service/user_service/global"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var dest io.Writer

// InitLogger
// @Description: 初始化日志
//

var lumberjackLogger *lumberjack.Logger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	zap.S().Infow("日志初始化成功")
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberjackLogger = &lumberjack.Logger{
		Filename:   createLogFileName(),
		MaxSize:    1,
		MaxAge:     5,
		MaxBackups: 30,
	}
	dest = io.MultiWriter(lumberjackLogger, os.Stdout)
	return zapcore.AddSync(dest)
}

func createLogFileName() string {
	today := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s/%s.log", global.FilePath.LogFile, today)
}
