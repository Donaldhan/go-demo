package middleware

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var sugarLogger *zap.SugaredLogger

func sugarLoggerDemo() {
	InitSugaredLogger()
	defer sugarLogger.Sync()
	simpleHttpGetSugaredLogger("www.topgoer.com")
	simpleHttpGetSugaredLogger("http://www.topgoer.com")
}

func InitSugaredLogger() {
	fileSyncer := getLogWriter()
	// 创建输出到控制台的 `WriteSyncer`
	consoleSyncer := zapcore.AddSync(os.Stdout)
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleSyncer, zapcore.DebugLevel), // 输出到控制台
		zapcore.NewCore(encoder, fileSyncer, zapcore.DebugLevel),    // 输出到文件
	)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "../log/app.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	// 创建两个 `Core`，一个输出到控制台，另一个输出到文件
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGetSugaredLogger(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
