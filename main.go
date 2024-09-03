package main

import (
	"boilerplate/cmd"

	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		log.Fatal().Msgf("failed run app: %s", err.Error())
	}
}
