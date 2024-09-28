package main

import (
	"go-ntc-templates/models"

	"github.com/melbahja/goph"
	"github.com/sirupsen/logrus"
	"github.com/wessrow/go-ntc-templates/models"
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

func asdf() {

	command := "show version"

	client, err := goph.New("admin", "10.0.100.208", goph.Password("Netw0rking!"))
	if err != nil {
		logrus.Error(err)
	}

	commandReturn := testSsh(client, command)

	test, err := ParseCommand[models.CiscoIosShowVersion](commandReturn)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Info(test[0].Hardware)
	// for _, result := range test {
	// 	logrus.Info(result.Mac_address)
	// }

}
