package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// TODO : Add log rotation

type NoLoggerFoundError struct{}

func (err NoLoggerFoundError) Error() string {
	return "Logger not initialized"
}

// Log Levels
const (
	INFO     string = "INFO"
	DEBUG    string = "DEBUG"
	WARN     string = "WARN"
	ERROR    string = "ERROR"
	TRACE    string = "TRACE"
	CRITICAL string = "CRITICAL"
)

type Logger struct {
	isLogEnabled bool
	logPath      string
	logMode      string
	logger_      *log.Logger
}

var _logger *Logger = nil
var once sync.Once

func GetInstance() (*Logger, error) {
	if _logger == nil {
		return nil, &NoLoggerFoundError{}
	}
	return _logger, nil
}

func InitLogger(logEnabled bool, logPath string, logMode string) {
	once.Do(func() {
		_logger = &Logger{}
		_logger.init(logEnabled, logPath, logMode)
	})
}

func (logger *Logger) init(logEnabled bool, logPath string, logMode string) {
	logger.isLogEnabled = logEnabled
	logger.logMode = logMode
	logger.logPath = logPath
	if logger.isLogEnabled {

		filePath := filepath.Join(logPath, "imposter."+time.Now().Format("2006-01-02 15")+".log")
		logfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		multiWriter := io.MultiWriter(os.Stdout, logfile)
		// defer logfile.Close()
		logger.logger_ = log.New(multiWriter, logMode+"|", log.Ldate|log.Ltime|log.LstdFlags)
	}
}

func (logger *Logger) Fatal(v ...any) {
	if logger.isLogEnabled {
		logger.logger_.Fatal(v...)
	}
}

func (logger *Logger) Panic(v ...any) {
	if logger.isLogEnabled {
		logger.logger_.Panic(v...)
	}
}

func (logger *Logger) Print(v ...any) {
	if logger.isLogEnabled {
		logger.logger_.Print(v...)
	}
}

func (logger *Logger) Printf(format string, v ...any) {
	if logger.isLogEnabled {
		logger.logger_.Printf(format, v...)
	}
}

func (logger *Logger) Println(v ...any) {
	if logger.isLogEnabled {
		logger.logger_.Println(v...)
	}
}
