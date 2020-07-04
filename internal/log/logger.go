package log

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Logger = log.Sugar()
}
