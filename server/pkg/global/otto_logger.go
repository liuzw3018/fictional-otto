package global

import "fmt"

/**
 * @Author: liu zw
 * @Date: 2022/10/25 11:42
 * @File: otto_logger.go
 * @Description: //TODO $
 * @Version:
 */

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Infoln(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Fatalln(args ...interface{})
	Warnln(args ...interface{})
}

func NewLogger(l Logger) *FactoryLogger {
	return &FactoryLogger{
		l: l,
	}
}

type FactoryLogger struct {
	l Logger
}

func (f *FactoryLogger) Info(args ...interface{}) {
	f.l.Info(args)
}

func (f *FactoryLogger) Infoln(args ...interface{}) {
	f.l.Infoln(args)
}

func (f *FactoryLogger) Errorf(format string, args ...interface{}) {
	f.l.Errorf(format, args)
}

func (f *FactoryLogger) Fatalf(format string, args ...interface{}) {
	f.l.Fatalf(format, args)
}

func (f *FactoryLogger) Fatal(args ...interface{}) {
	f.l.Fatal(args)
}

func (f *FactoryLogger) Infof(format string, args ...interface{}) {
	f.l.Infof(format, args)
}

func (f *FactoryLogger) Warnf(format string, args ...interface{}) {
	f.l.Warnf(format, args)
}

func (f *FactoryLogger) Debugf(format string, args ...interface{}) {
	f.l.Debugf(format, args)
}

func (f *FactoryLogger) Debug(args ...interface{}) {
	f.l.Debug(args)
}

func (f *FactoryLogger) Println(args ...interface{}) {
	fmt.Println(args)
}

func (f *FactoryLogger) Error(args ...interface{}) {
	f.l.Error(args)
}

func (f *FactoryLogger) Fatalln(args ...interface{}) {
	f.l.Fatalln(args)
}

func (f *FactoryLogger) Warnln(args ...interface{}) {
	f.l.Warnln(args)
}
