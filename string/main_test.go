package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// +
func plusConnect(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return str
}

// SprintConnect fmt.Sprintf()
func SprintConnect(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return str
}

// BytesConnect bytes
func BytesConnect(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		// 将元素打散插入其中
		buf = append(buf, str...)
	}
	return string(buf)
}

// buffer
func BufferConnect(n int, str string) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

// Strings.Builder
func StringBuilderConnect(n int, str string) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(str)
	}
	return sb.String()
}

// preloadBytes
func PreLoadBytesConnect(n int, str string) string {
	b := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		b = append(b, str...)
	}
	return string(b)
}

func benchmarkString(b *testing.B, f func(n int, str string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

// buffer
func BenchmarkBufferConnect(b *testing.B) {
	benchmarkString(b, BufferConnect)
}

// plus
func BenchmarkPlusConnect(b *testing.B) {
	benchmarkString(b, plusConnect)
}

// bytes
func BenchmarkBytesConnect(b *testing.B) {
	benchmarkString(b, BytesConnect)
}

// PreLoadBytesConnect
func BenchmarkPreLoadBytesConnect(b *testing.B) {
	benchmarkString(b, PreLoadBytesConnect)
}

// SprintConnect
func BenchmarkSprintConnect(b *testing.B) {
	benchmarkString(b, SprintConnect)
}

// StringBuilderConnect
func BenchmarkStringBuilderConnect(b *testing.B) {
	benchmarkString(b, StringBuilderConnect)
}
