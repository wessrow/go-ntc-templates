package main

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/sirikothe/gotextfsm"
)

func parsePlatformAndCommand(modelName string) (string, string, error) {

	re := regexp.MustCompile(`(\w+?)(Show.+)`)
	matches := re.FindStringSubmatch(modelName)

	if len(matches) < 3 {
		return "", "", errors.New("could not parse model name to platform")
	}

	platform := matches[1]
	command := matches[2]

	// Use regex to convert CamelCase to snake_case
	reCamel := regexp.MustCompile("([a-z0-9])([A-Z])")
	platformSnakeCase := reCamel.ReplaceAllString(platform, "${1}_${2}")
	commandSnakeCase := reCamel.ReplaceAllString(command, "${1}_${2}")

	// Convert to lowercase
	return strings.ToLower(platformSnakeCase), strings.ToLower(commandSnakeCase), nil
}

func parseOutput[returnModel any](input string, template string) ([]returnModel, error) {
	var model []returnModel

	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		return model, err
	}

	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(input, fsm, true)
	if err != nil {
		return model, err
	}

	str, err := json.Marshal(parser.Dict)
	if err != nil {
		return model, err
	}

	err = json.Unmarshal(str, &model)
	if err != nil {
		return model, err
	}

	return model, nil
}

func ParseCommand[returnModel any](input string) ([]returnModel, error) {

	var modelName returnModel

	platform, command, err := parsePlatformAndCommand(reflect.TypeOf(modelName).Name())
	if err != nil {
		return nil, err
	}

	template, err := os.ReadFile("./templates/" + platform + "_" + command + ".textfsm")
	if err != nil {
		return nil, err
	}

	parsedCommand, err := parseOutput[returnModel](input, string(template))
	if err != nil {
		return nil, err
	}

	if len(parsedCommand) == 0 {
		return nil, errors.New("could not parse string to chosen struct")
	}

	return parsedCommand, nil

}
