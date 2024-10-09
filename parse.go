//go:generate bash -c "./get_latest_templates.sh && go run generate/generate.go"

package gontc

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/sirikothe/gotextfsm"
	"github.com/sirupsen/logrus"
	"github.com/wessrow/go-ntc-templates/models"
)

func parseOutput[returnModel any](input string, template string) ([]returnModel, error) {
	var model []returnModel

	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		logrus.Info(template)
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

	modelType := reflect.TypeOf((*returnModel)(nil)).Elem()

	// Find the corresponding template
	template, ok := models.TemplateMap[modelType]
	if !ok {
		return nil, errors.New("template not found for the given model")
	}

	parsedCommand, err := parseOutput[returnModel](input, template)
	if err != nil {
		return nil, err
	}

	if len(parsedCommand) == 0 {
		return nil, errors.New("could not parse string to chosen struct")
	}

	return parsedCommand, nil

}
