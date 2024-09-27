package main

import (
	"encoding/json"
	"fmt"
	"go-ntc-templates/models"
	"os"
	"strings"

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

func parseOutput(input string, template string) any {

	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		logrus.Error(err)
	}
	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(string(input), fsm, true)
	if err != nil {
		logrus.Error(err)
	}

	str, err := json.Marshal(parser.Dict)
	if err != nil {
		fmt.Println("Error during marshaling:", err)
	}
	fmt.Println("Marshalled JSON:", string(str))

	// Unmarshal into the struct
	model := []models.Cisco_ios_show_version{}
	err = json.Unmarshal(str, &model)
	if err != nil {
		fmt.Println("Error during unmarshaling:", err)
	}

	return model
}

func main() {

	command := "show version"

	template, err := os.ReadFile("./templates/cisco_ios_" + strings.ReplaceAll(command, " ", "_") + ".textfsm")
	if err != nil {
		logrus.Error(err)
	}

	client, err := goph.New("admin", "10.0.100.208", goph.Password("Netw0rking!"))
	if err != nil {
		logrus.Error(err)
	}
	commandReturn := testSsh(client, command)

	parsedCommand := parseOutput(commandReturn, string(template))

	logrus.Info(parsedCommand)

	//parse.ParseFSM("cisco_ios_show_interfaces", string(template))

}
