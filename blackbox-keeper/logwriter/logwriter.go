package logwriter

import (
	"bytes"
	"io"
	"log"
	"sync"
)

type LogWriter interface {
	Save(name string) chan error
	Read(buf []byte) (int, error)
}

func NewLogWriter(stdout, stderr io.ReadCloser) *logWriter {
	return &logWriter{
		Stdout: stdout,
		Stderr: stderr,
	}
}

type logWriter struct {
	Stderr io.ReadCloser
	Stdout io.ReadCloser

	err bytes.Buffer
	out bytes.Buffer
}

// Save saves pipe output to buffers in goroutine and sends error via chan
// err == nil on success
func (l *logWriter) Save(name string) chan error {
	e := make(chan error)
	go func() {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			_, err := io.Copy(&l.err, l.Stderr)
			if err != nil {
				log.Printf("error on writing from stderr pipe of %s: %v", name, err)
			}
			wg.Done()
		}()
		go func() {
			_, err := io.Copy(&l.out, l.Stdout)
			if err != nil {
				log.Printf("error on writing from stdout pipe of %s: %v", name, err)
			}
			wg.Done()
		}()
		wg.Wait()
		e <- nil
	}()
	return e
}

// Read copies saved output from stderr and stdout in buf (relatively)
func (l *logWriter) Read(buf []byte) (int, error) {
	e, err1 := l.err.Read(buf)
	if err1 != nil {
		return 0, err1
	}
	o, err2 := l.out.Read(buf)
	if err2 != nil {
		return 0, err2
	}
	return e + o, nil
}
