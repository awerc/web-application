//Package for validation xml file parsed into custom structure
package xml_validator

import (
	"errors"
	"regexp"

	"web-application/xml_parser"
)

//Function that checks Listen, Timeout and User fileds
func Validate(config xml_parser.Configuration) error {
	//The range of valid ports is [1024;65535]
	if config.Connect.Port < 1024 || config.Connect.Port > 65535 {
		return errors.New("Wrong port")
	}
	//The minimum timeout is 100
	if config.Connect.Timeout < 100 {
		return errors.New("Too short timeout")
	}
	//Username must contain only letters and numbers
	match, _ := regexp.MatchString("^\\w+$", config.User.Login)
	if !match {
		return errors.New("Wrong user name")
	}

	return nil
}
