package utils

import (
	"github.com/rs/zerolog/log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Error().Err(err).Str("message", msg)
	}
}
