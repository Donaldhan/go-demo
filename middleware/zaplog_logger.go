package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

// 结构化数据
var logger *zap.Logger

func loggerDemo() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.5lmh.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
