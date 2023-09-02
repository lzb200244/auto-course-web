package initialize

import (
	"go-template/global"
	"go-template/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

// InitLogger 初始化logger实例
func InitLogger() {
	//控制台输出
	if global.Config.Project.Mode == "dev" {
		global.Logger, _ = zap.NewDevelopment()
	} else {
		// 创建根目录
		createRootDir()
		// 设置日志等级
		setLogLevel()
		if global.Config.Log.ShowLine {
			options = append(options, zap.AddCaller())
		}
		global.Logger = zap.New(getZapCore(), options...)
	}

	global.Logger.Debug("日志初始化成功!")
}

func createRootDir() {
	if ok, _ := utils.PathExists(global.Config.Log.RootDir); !ok {
		_ = os.Mkdir(global.Config.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {

	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}
	//encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
	//	encoder.AppendString(global.Config.Project.Mode + "." + l.String())
	//}

	// 设置编码器
	if global.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.Config.Log.RootDir + "/" + global.Config.Log.Filename,
		MaxSize:    global.Config.Log.MaxSize,
		MaxBackups: global.Config.Log.MaxBackups,
		MaxAge:     global.Config.Log.MaxAge,
		Compress:   global.Config.Log.Compress,
	}

	return zapcore.AddSync(file)
}
