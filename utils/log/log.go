package log

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("---> \033[34m%-6s\033[0m\n", strings.ToUpper(i.(string)))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log.Logger = zerolog.New(output).With().Timestamp().Logger()
}

func Logging(state string, function string, payload interface{}) {
	_payload, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		msg := fmt.Sprintf("\033[31m  at : \033[0m%s\n\033[31mError : \033[0m%s", function, payload)
		log.Fatal().Msgf(msg)
		return
	}
	switch state {
	case utils.REQUEST_LOG:
		msg := fmt.Sprintf("\033[36m    at : \033[0m%s\n\033[36mrequest : \033[0m%s", function, string(_payload))
		log.Info().Msgf(msg)
	case utils.RESPONSE_LOG:
		msg := fmt.Sprintf("\033[36m          at : \033[0m%s\n\033[36m    resposne : \033[0m%s\n\033[33m", function, string(_payload))
		log.Info().Msgf(msg)
	case utils.EXCEPTION_LOG:
		msg := fmt.Sprintf("\033[31m    at : \033[0m%s\n\033[31mpayload : \033[0m%s", function, string(_payload))
		log.Error().Msgf(msg)
	case utils.ERR_LOG:
		msg := fmt.Sprintf("\033[31m    at : \033[0m%s\n\033[31mpayload : \033[0m%s", function, string(_payload))
		log.Fatal().Msgf(msg)
	case utils.INFO_LOG:
		msg := fmt.Sprintf("\033[36m at : \033[0m%s\n\033[36mInfo : \033[0m%s", function, string(_payload))
		log.Info().Msgf(msg)
	}
}
