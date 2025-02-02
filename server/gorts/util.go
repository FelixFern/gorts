package gorts

import (
	"io"
	"net/http"
)

type readWriteCloser struct {
	r []byte
	w http.ResponseWriter
}

func (rwc *readWriteCloser) Read(p []byte) (n int, err error) {
	n = copy(p, rwc.r)
	if n == 0 {
		return 0, io.EOF
	}
	rwc.r = rwc.r[n:]
	return n, nil
}

func (rwc *readWriteCloser) Write(p []byte) (n int, err error) {
	return rwc.w.Write(p)
}

func (rwc *readWriteCloser) Close() error {
	return nil
}
