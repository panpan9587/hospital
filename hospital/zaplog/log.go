package zaplog

import (
	"go.uber.org/zap"
	"log"
)

func InitZapLog(path string) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		path,
		"stderr",
		"stdout",
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatalln("初始化日志失败", err)
	}
	zap.ReplaceGlobals(logger)
}
