package parse

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s // return the original string if it's empty
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func capitalizeFirstLetterLowerRest(s string) string {
	if len(s) == 0 {
		return s // return the original string if it's empty
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func ParseFSM(name string, template string) {
	// scanner := bufio.NewScanner(strings.NewReader(template))
	// for scanner.Scan() {
	// 	if strings.HasPrefix(scanner.Text(), "Value") {

	// 		match, _ := regexp.MatchString("^Value\\s", scanner.Text())
	// 		fmt.Println(match.)
	// 	}
	// }

	r, _ := regexp.Compile(`Value\s(List|Required)?\s?(\w+)`)
	entries := r.FindAllStringSubmatch(template, -1)

	pwd, _ := os.Getwd()
	f, err := os.Create(pwd + "/models/" + name + ".go")

	check(err)
	defer f.Close()

	f.WriteString("package models\n\n")
	f.WriteString(fmt.Sprintf("type %v struct {\n", capitalizeFirstLetter(name)))
	for _, entry := range entries {
		if entry[1] == "List" {
			f.WriteString(fmt.Sprintf("\t%v\t[]string\t`json:\"%v\"`\n", capitalizeFirstLetterLowerRest(entry[2]), entry[2]))
			continue
		}
		f.WriteString(fmt.Sprintf("\t%v\tstring\t`json:\"%v\"`\n", capitalizeFirstLetterLowerRest(entry[2]), entry[2]))
	}
	f.WriteString("}")

}
