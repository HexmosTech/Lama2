// Package `outputmanager` provides facilities for controlling
// the logging library as well as capabilities to post-process
// API execution results (such as store results as a JSON file)
package outputmanager

import (
	"bytes"
	"fmt"
	"os"

	"github.com/HexmosTech/gabs/v2"
	"github.com/HexmosTech/httpie-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LogBuff is used to append various log statements into memory.
// If the user toggles the `Output (-o)` option, then the contents
// of LogBuff is pushed into a JSON file
var LogBuff bytes.Buffer

type ResponseTime struct {
	Type     string `json:"type"`
	TimeInMs int64  `json:"timeInMs"`
}
type StatusCode struct {
	Type   string `json:"type"`
	Status int    `json:"statusCode"`
}

type ContentSize struct {
	Type        string `json:"type"`
	SizeInBytes int    `json:"sizeInBytes"`
}

func init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
	consoleWriter2 := zerolog.ConsoleWriter{Out: &LogBuff}
	multi := zerolog.MultiLevelWriter(consoleWriter, consoleWriter2)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	log.Logger = logger
}

// ConfigureZeroLog provides global log level setting. By default,
// ZeroLog uses the DEBUG level; however, the function makes the
// desired level more explicit
func ConfigureZeroLog(level string) {
	logLevelMap := make(map[string]zerolog.Level)
	logLevelMap["ERROR"] = zerolog.ErrorLevel // less information
	logLevelMap["INFO"] = zerolog.InfoLevel
	logLevelMap["DEBUG"] = zerolog.DebugLevel
	logLevelMap["TRACE"] = zerolog.TraceLevel // maximum information
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(logLevelMap[level])
}

func wrapError(requestError string) *gabs.Container {
	temp := gabs.New()
	temp.Set(requestError, "errors")
	return temp
}

func ResponseToJSON(resp httpie.ExResponse, responseTimes []ResponseTime, statusCode []StatusCode, contentSizes []ContentSize) (*gabs.Container, error) {
	body := string(resp.Body)

	var headerMapStr string
	for k, v := range resp.Headers {
		headerMapStr += k + ": " + v + "\n"
	}
	var totalResponseTimeMs int64
	for _, rt := range responseTimes {
		totalResponseTimeMs += rt.TimeInMs
	}

	performance := map[string]interface{}{
		"responseTimes":         responseTimes,
		"totalResponseTimeInMs": totalResponseTimeMs,
	}

	temp := gabs.New()
	temp.Set(headerMapStr, "headers")
	temp.Set(body, "body")
	temp.Set(LogBuff.String(), "logs")
	temp.Set(performance, "performance")
	temp.Set(statusCode, "statusCodes")
	temp.Set(contentSizes, "contentSizes")

	return temp, nil
}

// WriteJSONOutput is primarily built for helping with
// Extension/Integration building with external tools.
// Extension writers may simply call `l2 -n -o /tmp/lama2.json ...`
// to invoke WriteJSONOutput; the generated json file contains
// three keys: `logs`, `headers`, `body`
func WriteJSONOutput(resp httpie.ExResponse, targetPath string, responseTime []ResponseTime, statusCodes []StatusCode, contentSizes []ContentSize) {
	temp, _ := ResponseToJSON(resp, responseTime, statusCodes, contentSizes)
	err := os.WriteFile(targetPath, []byte(temp.String()), 0o644)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Couldn't write JSON output to: %s", targetPath))
	}
}
