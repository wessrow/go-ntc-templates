package main

import (
	"encoding/json"
	"fmt"
	"go-ntc-templates/models"
	"go-ntc-templates/parse"
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

func parseOutput[T any](input string, template string) (T, error) {
	var model T // Declare a variable of type T

	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		logrus.Error(err)
		return model, err // Return zero value and error
	}

	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(input, fsm, true)
	if err != nil {
		logrus.Error(err)
		return model, err // Return zero value and error
	}

	str, err := json.Marshal(parser.Dict)
	if err != nil {
		fmt.Println("Error during marshaling:", err)
		return model, err // Return zero value and error
	}
	fmt.Println("Marshalled JSON:", string(str))

	// Unmarshal into the struct
	err = json.Unmarshal(str, &model)
	if err != nil {
		fmt.Println("Error during unmarshaling:", err)
		return model, err // Return zero value and error
	}

	return model, nil // Return the populated model and nil error
}

func main() {

	// items, _ := ioutil.ReadDir("./templates")

	// for _, item := range items {
	// 	template, err := os.ReadFile("./templates/" + item.Name())

	// 	if err != nil {
	// 		logrus.Error(err)
	// 	}
	// 	name := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

	// 	parse.ParseFSM(name, string(template))

	// }

	command := "show interfaces"
	template, err := os.ReadFile("./templates/" + "cisco_ios_show_interfaces" + ".textfsm")
	if err != nil {
		logrus.Fatal(err)
	}

	client, err := goph.New("admin", "10.0.100.208", goph.Password("Netw0rking!"))
	if err != nil {
		logrus.Error(err)
	}

	commandReturn := testSsh(client, command)

	parsedCommand, err := parseOutput[[]models.Cisco_ios_show_interfaces](commandReturn, string(template))
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(parsedCommand)

	parse.ParseFSM("cisco_ios_show_interfaces", string(template))

}
