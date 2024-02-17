package tests

import (
	"log/slog"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	slog.Info("Start integration test..")
	StartDatabase()
	code := m.Run()
	slog.Info("End integration test..")
	os.Exit(code)
}
