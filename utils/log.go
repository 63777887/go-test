package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jwk/test/pkg/config"
	"os"
	"strings"
)

var (
	Logger *zap.Logger
)

func init() {
	InitLogger(config.Config)
}

func InitLogger(serverConfig config.ServerConfig) {

	logLevel := getLogLevel(serverConfig.Log.Level)

	encoder := getEncoder()
	// 写入到文件
	writeSyncer := getLogWriter(serverConfig)
	// 写入到终端
	//stdin，stdout和stderr，这3个可以称为终端（Terminal）的标准输入（standard input），标准输出（ standard out）和标准错误输出（standard error）。
	stdout := zapcore.AddSync(os.Stdout)
	//zapcore.AddSync(os.Stdin)
	//zapcore.AddSync(os.Stderr)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, stdout), logLevel)
	Logger = zap.New(core, zap.AddCaller())
}

func getLogLevel(level string) zapcore.Level {
	var logLevel zapcore.Level
	// debug，info，warn，error
	switch strings.ToUpper(level) {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
		break
	case "INFO":
		logLevel = zapcore.InfoLevel
		break
	case "WARN":
		logLevel = zapcore.WarnLevel
		break
	case "ERROR":
		logLevel = zapcore.ErrorLevel
		break
	default:
		logLevel = zapcore.DebugLevel
	}
	return logLevel
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 在zap中加入Lumberjack支持
func getLogWriter(serverConfig config.ServerConfig) zapcore.WriteSyncer {
	logConfig := serverConfig.Log

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logConfig.Filename,
		MaxSize:    logConfig.MaxSize,    // 以 MB 为单位
		MaxBackups: logConfig.MaxBackups, // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:     logConfig.MaxAge,     // 保留旧文件的最大天数
		Compress:   logConfig.Compress,   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
