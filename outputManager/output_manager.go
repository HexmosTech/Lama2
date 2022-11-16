package outputmanager

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/HexmosTech/gabs/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var LogBuff bytes.Buffer

func init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	consoleWriter2 := zerolog.ConsoleWriter{Out: &LogBuff}
	multi := zerolog.MultiLevelWriter(consoleWriter, consoleWriter2)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	log.Logger = logger
}

func ConfigureZeroLog(level string) {
	log_level_map := make(map[string]zerolog.Level)
	log_level_map["ERROR"] = zerolog.ErrorLevel // less information
	log_level_map["INFO"] = zerolog.InfoLevel
	log_level_map["DEBUG"] = zerolog.DebugLevel
	log_level_map["TRACE"] = zerolog.TraceLevel // maximum information
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(log_level_map[level])
}

func WriteJSONOutput(requestLog string, targetPath string) {
	fmt.Println("## Log output\n", LogBuff.String())
	fmt.Println("## HTTPie output\n", requestLog)
	var re = regexp.MustCompile(`(?m)^\s*[{\[<]`)

	idx := re.FindStringIndex(requestLog)
	headers := string(requestLog[:idx[0]])
	body := string(requestLog[idx[0]:])

	temp := gabs.New()
	temp.Set(headers, "headers")
	temp.Set(body, "body")
	temp.Set(LogBuff.String(), "logs")

	fmt.Println(temp)

}
