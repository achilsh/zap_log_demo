package log_wrapper

import (
	"context"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLog struct {
	logger *zap.Logger
}

var zapLogHandle *ZapLog

func NewZapLogHandler() *ZapLog {
	//
	zapLogHandle = &ZapLog{}
	zapLogHandle.logger, _ = zap.NewProduction()
	return zapLogHandle
}

func ZError(ctx context.Context, msg string) {
	zapLogHandle.logger.Error(msg)
}
func ZInfo(ctx context.Context, msg string) {
	zapLogHandle.logger.Info(msg)
}

// ////////
type SugaredLog struct {
	logger *zap.SugaredLogger
}

//DebugLevel Level = iota - 1
//// InfoLevel is the default logging priority.
//InfoLevel
//// WarnLevel logs are more important than Info, but don't need individual
//// human review.
//WarnLevel
//// ErrorLevel logs are high-priority. If an application is running smoothly,
//// it shouldn't generate any error-level logs.
//ErrorLevel
//// DPanicLevel logs are particularly important errors. In development the
//// logger panics after writing the message.
//DPanicLevel
//// PanicLevel logs a message, then panics.
//PanicLevel
//// FatalLevel logs a message, then calls os.Exit(1).
//FatalLevel

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpainc":
		return zapcore.DPanicLevel
	case "painc":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

var sugarZapLogHandle *SugaredLog

func NewSugaredZapLogHandler(cfg *LogConfig) *SugaredLog {
	sugarZapLogHandle = &SugaredLog{}

	getEncoder := func() zapcore.Encoder {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

		encoderConfig.FunctionKey = "func"
		//encoderConfig.NameKey = "1"
		//encoderConfig.CallerKey = "2"

		// 在日志文件中使用大写字母记录日志级别
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		// NewConsoleEncoder 打印更符合人们观察的方式
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	getLogWriter := func() zapcore.WriteSyncer {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   cfg.LogPath,
			MaxSize:    int(cfg.FileMaxSizeMB),    // MB,在进行切割之前，日志文件的最大大小
			MaxAge:     int(cfg.OldFileRemainDay), //保留旧文件的最大天数
			MaxBackups: int(cfg.OldFileNums),      // 保留旧文件的最大个数
			Compress:   cfg.OldCompress,           //是否压缩/归档旧文件
		}
		if cfg.LogStdout == true {
			return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
		}
		return zapcore.AddSync(lumberJackLogger)
	}

	core := zapcore.NewCore(getEncoder(), getLogWriter(), getLogLevel(cfg.LogLevel))

	development := zap.Development()
	//field := zap.Fields(zap.String("serviceName", "serviceName"))
	//
	//logger := zap.New(core, field, zap.AddCaller(), development)
	//logger := zap.New(core, zap.WithCaller(true), development)
	logger := zap.New(core, zap.AddCallerSkip(1), zap.WithCaller(true), development)
	sugarZapLogHandle.logger = logger.Sugar()

	return sugarZapLogHandle
}

func Infof(ctx context.Context, msg ...string) {
	sugarZapLogHandle.logger.Infof("%v", msg)
}
func Info(ctx context.Context, msg string) {
	sugarZapLogHandle.logger.Info(msg)
}
