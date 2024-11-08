package log

import (
	"github.com/zeromicro/go-zero/core/logc"
	"golang.org/x/net/context"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
)

func logDemo() {
	var c logx.LogConf
	conf.MustLoad("config.yaml", &c)
	logx.MustSetup(c)

	logx.AddWriter(logx.NewWriter(os.Stdout))

	ctx := context.Background()

	logc.Info(ctx, "info message")
	logc.Errorf(ctx, "error message: %d", 123)
	logc.Debugw(ctx, "info filed", logc.Field("task", "transferTask"))
	logc.Slowv(ctx, "object")
	ctx = logx.ContextWithFields(context.Background(), logx.Field("task", "marketOrder"))
	logc.Infow(ctx, "hello world")
	logc.Error(ctx, "error log")

	for {
		select {
		case <-proc.Done():
			return
		default:
			time.Sleep(time.Second)
			logx.Info(time.Now())
		}
	}
}
