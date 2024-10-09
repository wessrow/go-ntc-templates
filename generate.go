//go:generate bash -c "./get_latest_templates.sh && go run ."

package main

import (
	"github.com/wessrow/go-ntc-templates/generate"
)

func main() {
	generate.GenerateFSMStructs("")
}
