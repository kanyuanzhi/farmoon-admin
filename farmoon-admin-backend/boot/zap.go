package boot

import (
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func Zap() *zap.Logger {
	zapConfig := global.FXConfig.Zap
	if err := utils.CreatePath(zapConfig.Path); err != nil {
		panic(err)
	}
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   zapConfig.Path + "/" + zapConfig.Filename + ".log",
		MaxSize:    zapConfig.MaxSize,
		MaxAge:     zapConfig.MaxAge,
		MaxBackups: zapConfig.MaxBackups,
	})
	multiWriterSyncer := zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	encoderConfig.TimeKey = "time"
	var encoder zapcore.Encoder
	if global.FXConfig.Zap.JsonEncoder {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(zapConfig.Level)); err != nil {
		panic(err)
	}
	var core zapcore.Core
	if global.FXConfig.Zap.OsStdout {
		core = zapcore.NewCore(encoder, multiWriterSyncer, l)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	logger := zap.New(core, zap.AddCaller())
	logger.Sugar()
	return logger
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.FXConfig.Zap.Prefix + " 2006-01-02T15:04:05.000Z07:00"))
}
