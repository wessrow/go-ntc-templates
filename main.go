package main

import (
	"encoding/json"
	"go-ntc-templates/models"
	"os"

	"github.com/melbahja/goph"
	"github.com/sirikothe/gotextfsm"
	"github.com/sirupsen/logrus"
)

func testSsh(client *goph.Client, command string) string {

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run(command)

	if err != nil {
		logrus.Fatal(err)
	}

	return string(out)
}

func parseOutput[T any](input string, template string) ([]T, error) {
	var model []T

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

func main() {

	command := "show ip route"
	template, err := os.ReadFile("./templates/" + "cisco_ios_show_ip_route" + ".textfsm")
	if err != nil {
		logrus.Fatal(err)
	}

	client, err := goph.New("admin", "10.0.100.208", goph.Password("Netw0rking!"))
	if err != nil {
		logrus.Error(err)
	}

	commandReturn := testSsh(client, command)

	parsedCommand, err := parseOutput[models.CiscoIosShowIpRoute](commandReturn, string(template))
	if err != nil {
		logrus.Error(err)
	}

	for _, intf := range parsedCommand {
		logrus.Info(intf.Network)
	}

}
