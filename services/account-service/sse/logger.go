package sse

import (
	"github.com/rs/zerolog/log"
)

const (
	protocolKey   = "protocol"
	methodKey     = "method"
	pathKey       = "path"
	durationKey   = "duration"
	statusCodeKey = "statusCode"
	statusTextKey = "statusText"
)

func LogInfo(msg string) {
	log.Info().
		Str(protocolKey, "HTTP SSE").
		Msg(msg)
}

func LogError(err error, msg string) {
	if msg == "" {
		log.Error().Err(err)
	} else {
		log.Error().Err(err).Msg(msg)
	}
}
