# go-ntc-templates

This package aims to implement further golang functionality when working with network-to-code templates for network device automation by combining the awesome screen-scraping capabilities of TextFSM with the rigid structure of working with `golang structs`.

This package utilizes the `github.com/sirikothe/gotextfsm` package which parses TextFSM-templates to JSON.

Though JSON support for golang is extended, it is not as easy to work with as `structs`.

This package therefore includes programmatically parsed and rendered golang-structs from the latest ntc-templates release.

## Getting started

Get the library:

```bash
$ go get github.com/wessrow/go-ntc-templates
```

### Usage

This package is really simple, and for parsing purposes only includes one public function `ParseCommand`.

`ParseCommand` makes use of golang `Generics` to pass a `struct` for the parsed data to be returned as. On top of this it takes two arguments:

#### input

The data to be parsed as a string.

#### Example

```golang
// If the target device is of model cisco_ios, and command is ex. "show version"
command := "show version"
// Ssh-execution logic omitted for brevity, see below full example.
test, _ := gontc.ParseCommand[cisco_ios.ShowVersion](commandReturn)
```

### Models

The generated `golang structs` are sorted under the `models` package. All `ntc-templates` from the latest release are placed into sub-packages following this structure:

`/models/<manufacturer_os>/<template>`

The textfsm template for `cisco_ios_show_version` would therefore be usable via `cisco_ios.ShowVersion`

### Return

`ParseCommand` returns a list of the passed model. This is because of how TextFSM templates are parsed. For endpoints that return multiple outputs of the same type (interfaces, vlans, etc.) should be returned as a list of said output.

This package takes the same approach and returns a slice containing the parsed struct for each output.

## Full Usage Example

_This Example uses the popular Golang SSH-client `goph` - though any SSH-client that returns data as `[]byte` or `string` would work to be able to pass string to `ParseCommand`_

```golang
package main

import (
	"fmt"

	"github.com/melbahja/goph"
	"github.com/wessrow/go-ntc-templates"
	"github.com/wessrow/go-ntc-templates/models/cisco_ios"
)

func main() {

	command := "show version"

	// Initiate goph client.
	client, err := goph.New("<user>", "<host-ip>", goph.Password("<password>"))
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run(command)

	if err != nil {
		fmt.Println(err)
	}

	test, _ := gontc.ParseCommand[cisco_ios.ShowVersion](string(out))

	fmt.Println(test[0].Hardware)

}

```
