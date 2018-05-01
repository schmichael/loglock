package loglock

import (
	"log"
	"runtime"
	"sync"
)

type LogLock struct {
	mu   sync.Mutex
	name string
}

func NewLogLock(name string) *LogLock {
	return &LogLock{name: name}
}

func (l *LogLock) caller() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	f := runtime.FuncForPC(pc)
	if ok && f != nil {
		return f.Name()
	}
	return "unknown"
}

func (l *LogLock) Lock() {
	log.Printf("Locking %s by %s", l.name, l.caller())
	l.mu.Lock()
	log.Printf("Locked  %s by %s", l.name, l.caller())
}

func (l *LogLock) Unlock() {
	log.Printf("Unlocking %s by %s", l.name, l.caller())
	l.mu.Unlock()
	log.Printf("Unlocked  %s by %s", l.name, l.caller())
}
