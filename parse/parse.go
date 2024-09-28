package parse

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sirikothe/gotextfsm"
	"github.com/sirupsen/logrus"
)

func parseOutput[returnModel any](input string, template string) ([]returnModel, error) {
	var model []returnModel

	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		logrus.Error(err)
		return model, err
	}

	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(input, fsm, true)
	if err != nil {
		logrus.Error(err)
		return model, err
	}

	str, err := json.Marshal(parser.Dict)
	if err != nil {
		logrus.Error("Error during marshaling:", err)
		return model, err
	}

	err = json.Unmarshal(str, &model)
	if err != nil {
		logrus.Error("Error during unmarshaling:", err)
		return model, err
	}

	return model, nil
}

func ParseCommand[returnModel any](command string, input string, platform string) ([]returnModel, error) {

	command = strings.ReplaceAll(command, " ", "_")

	pwd, _ := os.Getwd()
	template, err := os.ReadFile(pwd + "/templates/" + platform + "_" + command + ".textfsm")
	if err != nil {
		logrus.Fatal(err)
	}

	parsedCommand, err := parseOutput[returnModel](input, string(template))
	if err != nil {
		logrus.Error(err)
	}

	return parsedCommand, nil

}
