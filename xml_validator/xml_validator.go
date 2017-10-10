package xml_validator

import (
	"errors"
	"regexp"

	"web-application/xml_parser"
)

func Validate(config xml_parser.Configuration) error {

	if config.Http.Listen < 1024 || config.Http.Listen > 65535 {
		return errors.New("Wrong port")
	}

	if config.Http.Timeout < 100 {
		return errors.New("Too short timeout")
	}

	match, _ := regexp.MatchString("^\\w+$", config.Database.User)
	if !match {
		return errors.New("Wrong user name")
	}

	return nil
}
