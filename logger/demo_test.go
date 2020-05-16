package logger

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

type fakeLogger struct {
	sb strings.Builder
}

func (fl *fakeLogger) Println(v ...interface{}) {
	fmt.Fprintln(&fl.sb, v...)
}

func (fl *fakeLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&fl.sb, format, v...)
}

func TestDemoV3(t *testing.T) {
	var fl fakeLogger
	DemoV3(fl.Println)
	want := "error at some things"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Log = %q; want substing %q\n", got, want)
	}
}

func TestDemoV4(t *testing.T) {
	var fl fakeLogger
	DemoV4(&fl)
	want := "error at some things"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Log = %q; want substing %q\n", got, want)
	}
}

func TestDemoV5(t *testing.T) {
	var sb strings.Builder
	logger := log.New(&sb, "", 0)
	thing := Thing{
		Logger: logger,
	}
	thing.DemoV5()
	want := "error at some things"
	got := sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Log = %q; want substing %q\n", got, want)
	}
}
