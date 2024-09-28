package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

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

func toCamelCase(str string) string {
	// Replace all - and . with spaces
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, ".", " ")
	str = strings.ReplaceAll(str, "_", " ")

	// Split the string by spaces
	parts := strings.Fields(str)

	// Convert each part to TitleCase (CamelCase)
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join them back together to form CamelCase
	return strings.Join(parts, "")
}

func parseFSM(name string, template string) {

	r, _ := regexp.Compile(`(?m)^\s*Value\s+((List,?|Required,?|Fillup,?|Filldown,?|Key,?)+)?\s?(\w+)`)
	entries := r.FindAllStringSubmatch(template, -1)

	pwd, _ := os.Getwd()
	f, err := os.Create(pwd + "/models/" + name + ".go")

	check(err)
	defer f.Close()

	name = toCamelCase(name)

	f.WriteString("package models\n\n")
	f.WriteString(fmt.Sprintf("type %v struct {\n", capitalizeFirstLetter(name)))
	for _, entry := range entries {
		if entry[2] == "List" {
			f.WriteString(fmt.Sprintf("\t%v\t[]string\t`json:\"%v\"`\n", capitalizeFirstLetterLowerRest(entry[3]), entry[3]))
			continue
		}
		f.WriteString(fmt.Sprintf("\t%v\tstring\t`json:\"%v\"`\n", capitalizeFirstLetterLowerRest(entry[3]), entry[3]))
	}
	f.WriteString("}")

}

func GenerateFSMStructs() {
	items, _ := ioutil.ReadDir("./templates")

	for _, item := range items {
		template, err := os.ReadFile("./templates/" + item.Name())

		if err != nil {
			logrus.Error(err)
		}
		name := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

		parseFSM(name, string(template))

	}
}
