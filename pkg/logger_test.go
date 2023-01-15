package pkg

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{InfoLogger: log.New(&buf, "", 0)}
	logger.Info("test message")

	assert.Equal(t, "test message\n", buf.String())
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{ErrorLogger: log.New(&buf, "", 0)}
	logger.Error("test message")

	assert.Equal(t, "test message\n", buf.String())
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{DebugLogger: log.New(&buf, "", 0), IsDevMode: true}
	logger.Debug("test message")

	assert.Equal(t, "test message\n", buf.String())
}

func TestDebugDevModeOff(t *testing.T) {
	var buf bytes.Buffer
	logger := Logger{DebugLogger: log.New(&buf, "", 0), IsDevMode: false}
	logger.Debug("test message")

	assert.Equal(t, "", buf.String())
}
