package main

import (
	"io"
	"log/slog"
	"os"
)

type logRead struct {
	io.Reader
	logger *slog.Logger
}

func NewLogRead(r io.Reader, logger *slog.Logger) logRead {
	return logRead{
		Reader: r,
		logger: logger,
	}
}

func (lr logRead) Read(p []byte) (int, error) {
	n, err := lr.Reader.Read(p)
	if err != nil {
		lr.logger.Error("read", slog.Any("err", err))
	} else {
		lr.logger.Info("read", slog.Int("n", n))
	}
	return n, err
}

func main() {
	lr := NewLogRead(os.Stdin, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))
	io.Copy(os.Stdout, lr)
}
