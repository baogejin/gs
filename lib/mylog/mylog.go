package mylog

import (
	"fmt"
	"gs/define"
	"gs/lib/config"
	"os"
	"strings"
	"sync"

	go_logger "gs/lib/mylog/go-logger"
)

var logger *go_logger.Logger
var once sync.Once

func getLogger() *go_logger.Logger {
	once.Do(func() {
		logger = go_logger.NewLogger()
		logger.Detach("console")

		format := "[%millisecond_format%][%level_string%][%file%:%line%]%body%"
		//格式字段参考:
		//{"timestamp":1669172510,"timestamp_format":"2022-11-23 11:01:50","millisecond":1669172510595,
		//"millisecond_format":"2022-11-23 11:01:50.595","level":3,"level_string":"Error","body":"this is a error format log!",
		//"file":"main.go","line":44,"function":"main.main"}

		levleStr := strings.ToLower(config.Get().LogLevel)
		logLevel := logger.LoggerLevel(levleStr)
		// 命令行输出配置
		consoleConfig := &go_logger.ConsoleConfig{
			Color:      true,   // 命令行输出字符串是否显示颜色
			JsonFormat: false,  // 命令行输出字符串是否格式化
			Format:     format, // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
		}
		// 添加 console 为 logger 的一个输出
		logger.Attach("console", logLevel, consoleConfig)

		name := ""
		for _, v := range os.Args {
			if s := strings.Split(v, "-node="); len(s) > 1 {
				name = s[1]
			}
		}
		if name == "" {
			name = "tmp"
		}
		path := os.Getenv(define.EnvName)
		if path == "" {
			path = "."
		}
		// 文件输出配置
		fileConfig := &go_logger.FileConfig{
			Filename: path + "/log/" + name + "_" + fmt.Sprintf("%d", os.Getpid()) + ".log", // 日志输出文件名，不自动存在
			// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
			// LevelFileName: map[int]string{
			// 	logger.LoggerLevel("error"): "./error.log", // Error 级别日志被写入 error .log 文件
			// 	logger.LoggerLevel("info"):  "./info.log",  // Info 级别日志被写入到 info.log 文件中
			// 	logger.LoggerLevel("debug"): "./debug.log", // Debug 级别日志被写入到 debug.log 文件中
			// },
			MaxSize:    0,      // 文件最大值（KB），默认值0不限
			MaxLine:    0,      // 文件最大行数，默认 0 不限制
			DateSlice:  "",     // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
			JsonFormat: false,  // 写入文件的数据是否 json 格式化
			Format:     format, // 如果写入文件的数据不 json 格式化，自定义日志格式
		}
		// 添加 file 为 logger 的一个输出
		logger.Attach("file", logLevel, fileConfig)
	})
	return logger
}

//log emergency level
func Emergency(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_EMERGENCY, msg)
}

//log emergency format
func Emergencyf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_EMERGENCY, msg)
}

//log alert level
func Alert(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_ALERT, msg)
}

//log alert format
func Alertf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_ALERT, msg)
}

//log critical level
func Critical(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_CRITICAL, msg)
}

//log critical format
func Criticalf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_CRITICAL, msg)
}

//log error level
func Error(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_ERROR, msg)
}

//log error format
func Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_ERROR, msg)
}

//log warning level
func Warning(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_WARNING, msg)
}

//log warning format
func Warningf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_WARNING, msg)
}

//log notice level
func Notice(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_NOTICE, msg)
}

//log notice format
func Noticef(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_NOTICE, msg)
}

//log info level
func Info(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_INFO, msg)
}

//log info format
func Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_INFO, msg)
}

//log debug level
func Debug(msg string) {
	getLogger().Writer(go_logger.LOGGER_LEVEL_DEBUG, msg)
}

//log debug format
func Debugf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	getLogger().Writer(go_logger.LOGGER_LEVEL_DEBUG, msg)
}
