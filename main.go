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

func testSsh(template string) models.Cisco_ios_show_version {

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

	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(string(out), fsm, true)
	if err != nil {
		logrus.Error(err)
	}

	str, err := json.Marshal(parser.Dict[0])
	if err != nil {
		fmt.Println("Error during marshaling:", err)
	}
	fmt.Println("Marshalled JSON:", string(str))

	// Unmarshal into the struct
	model := models.Cisco_ios_show_version{}
	err = json.Unmarshal(str, &model)
	if err != nil {
		fmt.Println("Error during unmarshaling:", err)
	}

	return model
}

func main() {

	template, err := os.ReadFile("./templates/cisco_ios_show_version.textfsm")
	if err != nil {
		logrus.Error(err)
	}
	test := testSsh(string(template))
	logrus.Info(test.MAC_ADDRESS)
	// parse.ParseFSM("cisco_ios_show_version", string(template))

}
