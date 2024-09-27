package main

import (
	"encoding/json"
	"fmt"
	"go-ntc-templates/models"
	"os"

	"github.com/melbahja/goph"
	"github.com/sirikothe/gotextfsm"
	"github.com/sirupsen/logrus"
)

func testSsh(template string) error {

	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(string(template))
	if err != nil {
		logrus.Error(err)
	}

	client, err := goph.New("admin", "10.0.100.208", goph.Password("Netw0rking!"))
	if err != nil {
		logrus.Error(err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run("show version")

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info(string(out))

	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(string(out), fsm, true)
	if err != nil {
		logrus.Error(err)
	}

	str, err := json.Marshal(parser.Dict)
	if err != nil {
		fmt.Printf("Unable to convert dict to json \n", err)
	}

	logrus.Info(string(str))

	model := models.Cisco_ios_show_version{}
	return json.Unmarshal(str, &model)
}

func main() {

	template, err := os.ReadFile("./templates/cisco_ios_show_version.textfsm")
	if err != nil {
		logrus.Error(err)
	}

	test := testSsh(string(template))
	logrus.Info(test)
	// parse.ParseFSM("cisco_ios_show_version", string(template))

}
