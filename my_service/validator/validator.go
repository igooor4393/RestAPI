package validator

import (
	"RestAPI/pkg/logger"
	"errors"
	"regexp"
)

var l = logger.Get()

func Valid(s string) error {

	match, err := regexp.MatchString("^[a-zA-Z0-9]+$", s)
	if err != nil {

		l.Error().Msgf("Error in regular expression:", err)

		return err
	}
	if !match {

		l.Error().Msg("Input string contains characters other than letters or numbers")

		return errors.New("input string contains characters other than letters or numbers")
	} else {

		l.Info().Msg("Input string contains only letters or numbers")
	}
	return nil
}
