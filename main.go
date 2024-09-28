package main

import (
	"go-ntc-templates/parse"

	"github.com/melbahja/goph"
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

func main() {

	command := "show ip route"

	client, err := goph.New("admin", "10.0.100.208", goph.Password("Netw0rking!"))
	if err != nil {
		logrus.Error(err)
	}

	commandReturn := testSsh(client, command)

	test := parse.ParseCommand(command, commandReturn, "cisco_ios")

	logrus.Info(test)

	// MERGE ALL BELOW LOGIC!
	// template, err := os.ReadFile("./templates/" + "cisco_ios_show_ip_route" + ".textfsm")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// parsedCommand, err := parse.ParseOutput[models.CiscoIosShowIpRoute](commandReturn, string(template))
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// for _, intf := range parsedCommand {
	// 	logrus.Info(intf.Network)
	// }

	// generate.GenerateFSMStructs()

}
