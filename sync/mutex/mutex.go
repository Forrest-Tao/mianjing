package main

import (
	"sync"
	"time"
)

const cost = time.Microsecond

type RW interface {
	Write()
	Read()
}

type Lock struct {
	mu    sync.Mutex
	count int
}

type RWLock struct {
	mu    sync.Mutex
	count int
}

func (l *Lock) Write() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.count++
	time.Sleep(cost)
}

func (l *Lock) Read() {
	l.mu.Lock()
	defer l.mu.Unlock()
	_ = l.count
	time.Sleep(cost)
}

func (l *RWLock) Write() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.count++
	time.Sleep(cost)
}

func (l *RWLock) Read() {
	l.mu.Lock()
	defer l.mu.Unlock()
	_ = l.count
	time.Sleep(cost)
}
