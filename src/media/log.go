package media

import (
	"io"
	"log"
	"sync"
)

const (
	LogDebug = iota
	LogInfo
	LogWarn
	LogError
)

type Log struct {
	level  uint8
	logger *log.Logger
	mu     sync.Mutex
}

func NewLog(level uint8, out io.Writer) *Log {
	l := new(Log)
	l.level = level
	l.logger = log.New(out, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return l
}

func (l *Log) SetLevel(newlevel uint8) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = newlevel
}

func (l *Log) Debug(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogDebug {
		return
	}

	l.logger.SetPrefix("[Debug]")
	l.logger.Println(v)
}

func (l *Log) Info(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogInfo {
		return
	}

	l.logger.SetPrefix("[Info] ")
	l.logger.Println(v)
}

func (l *Log) Warn(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogWarn {
		return
	}

	l.logger.SetPrefix("[Warn] ")
	l.logger.Println(v)
}

func (l *Log) Error(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogError {
		return
	}

	l.logger.SetPrefix("[Error]")
	l.logger.Println(v)
}

func (l *Log) Fatal(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.logger.SetPrefix("[Fatal]")
	l.logger.Fatalln(v)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogDebug {
		return
	}

	l.logger.SetPrefix("[Debug]")
	l.logger.Printf(format, v)
}

func (l *Log) Infof(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogInfo {
		return
	}

	l.logger.SetPrefix("[Info] ")
	l.logger.Printf(format, v)
}

func (l *Log) Warnf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogWarn {
		return
	}

	l.logger.SetPrefix("[Warn] ")
	l.logger.Printf(format, v)
}

func (l *Log) Errorf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > LogError {
		return
	}

	l.logger.SetPrefix("[Error]")
	l.logger.Printf(format, v)
}

func (l *Log) Fatalf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.logger.SetPrefix("[Fatal]")
	l.logger.Fatalf(format, v)
}
