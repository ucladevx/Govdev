package logger

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	assert := assert.New(t)

	// Test with debug as true
	{
		var debug bool = true
		l := NewLogger(debug, os.Stdout).(*logger)

		assert.Equal(l.l.Level, logrus.DebugLevel, "When set to debug, level should be DebugLevel")
		assert.Equal(l.l.Out, os.Stdout, "Output should be set to Stdout")
		assert.Equal(l.l.Formatter, &logrus.TextFormatter{}, "When set to debug, formatter should be TextFormatter")
	}

	// Test with debug as false, mimic production mode
	{
		var debug bool = false
		l := NewLogger(debug, os.Stdout).(*logger)

		assert.Equal(l.l.Level, logrus.InfoLevel, "When not set to debug, level should be InfoLevel")
		assert.Equal(l.l.Out, os.Stdout, "Output should be set to Stdout")
		assert.Equal(l.l.Formatter, &logrus.JSONFormatter{}, "When not set to debug, formatter should be JSONFormatter")
	}

}

// TODO: Regex debug outpu
func TestDebugLogging(t *testing.T) {

}

func TestInfoLogging(t *testing.T) {
	assert := assert.New(t)

	var debug bool = false
	var buffer bytes.Buffer
	l := NewLogger(debug, &buffer)

	l.Debug("test debug", map[string]string{
		"value1": "debug",
	})
	l.Info("test info", map[string]string{
		"value2": "info",
	})
	l.Warn("test warn", map[string]string{
		"value3": "warn",
	})
	l.Error("test error", map[string]string{
		"value4": "error",
	})

	outputs := strings.Split(strings.TrimSpace(buffer.String()), "\n")
	assert.Equal(len(outputs), 3)
	{
		jsonBuf := make(map[string]interface{})
		err := json.Unmarshal([]byte(outputs[0]), &jsonBuf)
		assert.Nil(err, "error should be nil while parsing JSON")
		assert.Equal(jsonBuf["msg"], "test info", "message was not logged")
		assert.Equal(jsonBuf["value2"], "info", "value2=info should be logged")
	}

	{
		jsonBuf := make(map[string]interface{})
		err := json.Unmarshal([]byte(outputs[1]), &jsonBuf)
		assert.Nil(err, "error should be nil while parsing JSON")
		assert.Equal(jsonBuf["msg"], "test warn", "message was not logged")
		assert.Equal(jsonBuf["value3"], "warn", "value3=warn should be logged")
	}

	{
		jsonBuf := make(map[string]interface{})
		err := json.Unmarshal([]byte(outputs[2]), &jsonBuf)
		assert.Nil(err, "error should be nil while parsing JSON")
		assert.Equal(jsonBuf["msg"], "test error", "message was not logged")
		assert.Equal(jsonBuf["value4"], "error", "value4=warn should be logged")
	}
}
