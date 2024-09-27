package parse

import (
	"fmt"
	"os"
	"regexp"
)

type CidrResponse struct {
	ID       string `json:"id"`
	Progress string `json:"progress"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFSM(name string, template string) {
	// scanner := bufio.NewScanner(strings.NewReader(template))
	// for scanner.Scan() {
	// 	if strings.HasPrefix(scanner.Text(), "Value") {

	// 		match, _ := regexp.MatchString("^Value\\s", scanner.Text())
	// 		fmt.Println(match.)
	// 	}
	// }

	r, _ := regexp.Compile(`Value\s(List)?\s?(\w+)`)
	entries := r.FindAllStringSubmatch(template, -1)

	pwd, _ := os.Getwd()
	f, err := os.Create(pwd + "/models/test.go")

	check(err)
	defer f.Close()

	f.WriteString("package models\n\n")
	f.WriteString(fmt.Sprintf("type %v struct {\n", name))
	for _, entry := range entries {
		if entry[1] == "List" {
			f.WriteString(fmt.Sprintf("\t%v\t[]string\t`json:\"%v\"`\n", entry[2], entry[2]))
			continue
		}
		f.WriteString(fmt.Sprintf("\t%v\tstring\t`json:\"%v\"`\n", entry[2], entry[2]))
	}
	f.WriteString("}")

}
