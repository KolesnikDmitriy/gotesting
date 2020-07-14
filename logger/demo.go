package logger

import (
	"errors"
	"log"
	"os"
	"sync"
)

func DemoV1() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	err := doSomeThings()
	if err != nil {
		logger.Println("error in doSomeThings():", err)
	}
}

func DemoV2(logger *log.Logger) {
	err := doSomeThings()
	if err != nil {
		logger.Println("error in doSomeThings():", err)
	}
}

// logger := log.New(...)
// DemoV3(log.Println)
func DemoV3(logFn func(...interface{})) {
	err := doSomeThings()
	if err != nil {
		logFn("error in doSomeThings():", err)
	}
}

type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

// logger := log.New(...)
// DemoV4(logger)
func DemoV4(logger Logger) {
	err := doSomeThings()
	if err != nil {
		logger.Println("error in doSomeThings():", err)
		logger.Printf("error: %v\n", err)
	}
}

type Thing struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
}

func (t Thing) DemoV5() {
	err := doSomeThings()
	if err != nil {
		t.Logger.Println("error in doSomeThings():", err)
		t.Logger.Printf("error: %v\n", err)
	}
}

func doSomeThings() error {
	return errors.New("error at some things")
}

func DemoV6(logger Logger) {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	err := doSomeThings()
	if err != nil {
		logger.Println("error in doSomeThings():", err)
		logger.Printf("error: %v\n", err)
	}
}

type ThingV2 struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
	once sync.Once
}

func (t *ThingV2) logger() Logger {
	t.once.Do(func() {
		if t.Logger == nil {
			t.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
		}
	})
	return t.Logger
}

func (t *ThingV2) DemoV7() {
	err := doSomeThings()
	if err != nil {
		t.logger().Println("error in doSomeThings():", err)
		t.logger().Printf("error: %v\n", err)
	}
}
