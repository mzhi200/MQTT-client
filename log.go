package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"sync"
)

type mqttTraceType uint8
const (
	MqttTraceTypeDebug mqttTraceType = 0x01
	MqttTraceTypeCritical mqttTraceType = 0x02
	MqttTraceTypeWarn mqttTraceType = 0x04
	MqttTraceTypeError mqttTraceType = 0x08
)

func mqttPackageTraceInit(out io.Writer, traceFlag mqttTraceType, flag int) {
	if traceFlag & MqttTraceTypeDebug != 0 {
		mqtt.DEBUG = log.New(out, "DEBUG", flag)
	}
	if traceFlag & MqttTraceTypeCritical != 0 {
		mqtt.CRITICAL = log.New(out, "CRITICAL", flag)
	}
	if traceFlag & MqttTraceTypeWarn != 0 {
		mqtt.WARN = log.New(out, "WARN", flag)
	}
	if traceFlag & MqttTraceTypeError != 0 {
		mqtt.ERROR = log.New(out, "ERROR", flag)
	}
}

type mqttLog interface {
	Init(out io.Writer, flag logLevelFlag)
	Info(msg string, args...interface{})
	Debug(msg string, args...interface{})
	Warn(msg string, args...interface{})
	Error(msg string, args...interface{})
}

type logLevelFlag uint8

const (
	LogLevelInfo  logLevelFlag = 0x01
	LogLevelDebug logLevelFlag = 0x02
	LogLevelWarn  logLevelFlag = 0x04
	LogLevelError logLevelFlag = 0x08
	LogLevelAll   logLevelFlag = 0x0f
)

type mqttClientLog struct {
	mutex    sync.RWMutex
	logger   *log.Logger
	logLevel logLevelFlag
}

func (l *mqttClientLog) Init(out io.Writer, flag logLevelFlag) {
	l.logger = log.New(out, "", log.LstdFlags|log.Lshortfile)
	l.logLevel = flag
}

func (l *mqttClientLog) Info(fmtString string, args...interface{}) {
	if l.logLevel & LogLevelInfo != 0 {
		if l.logger == nil {
			return
		}
		l.logger.SetPrefix("[INFO]")
		msg := fmt.Sprintf(fmtString, args...)
		l.logger.Printf(msg)
	}
}

func (l *mqttClientLog) Debug(fmtString string, args...interface{}) {
	if l.logLevel & LogLevelDebug != 0 {
		if l.logger == nil {
			return
		}
		l.logger.SetPrefix("[DEBUG]")
		msg := fmt.Sprintf(fmtString, args...)
		l.logger.Printf(msg)
	}
}

func (l *mqttClientLog) Warn(fmtString string, args...interface{}) {
	if l.logLevel & LogLevelWarn != 0 {
		if l.logger == nil {
			return
		}
		l.logger.SetPrefix("[WARN]")
		msg := fmt.Sprintf(fmtString, args...)
		l.logger.Printf(msg)
	}
}

func (l *mqttClientLog) Error(fmtString string, args...interface{}) {
	if l.logLevel & LogLevelError != 0 {
		if l.logger == nil {
			return
		}
		l.logger.SetPrefix("[ERROR]")
		msg := fmt.Sprintf(fmtString, args...)
		l.logger.Printf(msg)
	}
}

//log level: debug > warn > info > error
func GetLogLevelFromConfig() (ll logLevelFlag, err error) {
	ll = 0
	if config.LogLevel != "" {
		switch config.LogLevel {
		case "debug":
			ll |= LogLevelDebug
			fallthrough
		case "warn":
			ll |= LogLevelWarn
			fallthrough
		case "info":
			ll |= LogLevelInfo
			fallthrough
		case "error":
			ll |= LogLevelError
		case "all":
			ll = LogLevelAll
		default:
			err = errors.New(fmt.Sprintf("Unknown Log Level: %s", config.LogLevel))
		}
	}
	return
}

var newLog = func (out io.Writer) mqttLog {
	log := &mqttClientLog{}
	if out != nil {
		log.Init(out, ue.logLevel)
	} else {
		log.Init(os.Stdout, ue.logLevel)
	}
	return log
}