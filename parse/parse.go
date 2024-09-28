package parse

import (
	"encoding/json"
	"fmt"
	"go-ntc-templates/models"
	"os"
	"strings"

	"github.com/sirikothe/gotextfsm"
	"github.com/sirupsen/logrus"
)

func getModel(platform, command string) (interface{}, error) {
	// Create a map linking platform and command strings to struct constructors
	modelMap := map[string]map[string]func() interface{}{
		"cisco_ios": {
			"show_ip_route":          func() interface{} { return &models.CiscoIosShowIpRoute{} },
			"show_interfaces_status": func() interface{} { return &models.CiscoIosShowInterfacesStatus{} },
		},
		// Add more platforms and commands as needed
	}

	// Check if the platform exists
	if platformMap, ok := modelMap[platform]; ok {
		// Check if the command exists for the platform
		if constructor, ok := platformMap[command]; ok {
			// Return the struct instance
			return constructor(), nil
		}
		return nil, fmt.Errorf("command %s not found for platform %s", command, platform)
	}
	return nil, fmt.Errorf("platform %s not found", platform)
}

func ParseOutput(input string, template string, model interface{}) (interface{}, error) {

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

func ParseCommand(command string, input string, platform string) interface{} {

	command = strings.ReplaceAll(command, " ", "_")

	pwd, _ := os.Getwd()
	template, err := os.ReadFile(pwd + "/templates/" + platform + "_" + command + ".textfsm")
	if err != nil {
		logrus.Fatal(err)
	}

	testmodel, err := getModel(platform, command)
	if err != nil {

	}

	parsedCommand, err := ParseOutput(input, string(template), testmodel)
	if err != nil {
		logrus.Error(err)
	}

	return parsedCommand

	// for _, intf := range parsedCommand {
	// 	logrus.Info(intf.Network)
	// }
}
