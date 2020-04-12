package common

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	today    int64
	filepath string
	file     *os.File
	lock     sync.Mutex
}

func (l *Logger) Write(p []byte) (n int, err error) {
	if err := l.rotate(); err != nil {
		return 0, err
	}
	return l.file.Write(p)
}

func (l *Logger) Close() error {
	return nil
}

func (l *Logger) close() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil
	return err
}

func (l *Logger) rotate() error {
	today := l.get_current_day()
	l.lock.Lock()
	defer l.lock.Unlock()
	if today != l.today {
		if err := l.close(); err != nil {
			return err
		}
		if err := l.openFile(); err != nil {
			return err
		}
		l.today = today
	}
	return nil
}

func (l *Logger) openFile() error {
	name := l.get_filename()
	_, err := os.Stat(name)
	var f *os.File = nil
	if os.IsNotExist(err) {
		mode := os.FileMode(0600)
		file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
		if err != nil {
			return fmt.Errorf("can't open new logfile: %s", err)
		}
		f = file
	} else {
		file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("open logfile: %s", err)
		}
		f = file
	}
	l.file = f
	return nil
}

func (l *Logger) get_filename() string {
	current := time.Now()
	day_str := current.Format("2006_01_02")
	name := l.filepath + "_" + day_str + ".log"
	return name
}

func (l *Logger) get_current_day() int64 {
	return int64(time.Now().YearDay())
}

func (l *Logger) Init(path string) error {
	l.filepath = path
	l.today = l.get_current_day()
	err := l.openFile()
	return err
}

