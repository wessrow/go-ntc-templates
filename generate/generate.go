package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

func containsAny(str string, substrings []string) bool {
	for _, substring := range substrings {
		if strings.Contains(str, substring) {
			return true
		}
	}
	return false
}

func getPlatform(s string) string {

	exceptions := []string{"custom", "edgecore", "eltex", "fortinet", "linux", "yamaha"}

	regexPattern := `^([a-zA-Z0-9]+_[a-zA-Z0-9]+)`
	if containsAny(s, exceptions) {
		regexPattern = `^([a-zA-Z0-9]+)`
	}

	regex := regexp.MustCompile(regexPattern)

	return regex.FindString(s)
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

func extractEntries(template string) ([][]string, error) {
	r, err := regexp.Compile(`(?m)^\s*Value\s+((List,?|Required,?|Fillup,?|Filldown,?|Key,?)+)?\s?(\w+)`)
	if err != nil {
		return nil, err
	}
	return r.FindAllStringSubmatch(template, -1), nil
}

func createPlatformDirectory(platform string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dirPath := pwd + "/models/" + platform
	return os.MkdirAll(dirPath, os.ModePerm)
}

func writeStructToFile(f *os.File, platform, structName string, entries [][]string, template string) error {
	if _, err := fmt.Fprintf(f, "package %v \n\n", platform); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(f, "type %v struct {\n", capitalizeFirstLetter(structName)); err != nil {
		return err
	}

	for _, entry := range entries {
		fieldName := capitalizeFirstLetterLowerRest(entry[3])
		fieldType := "string"
		if entry[2] == "List" {
			fieldType = "[]string"
		}
		if _, err := fmt.Fprintf(f, "\t%v\t%v\t`json:\"%v\"`\n", fieldName, fieldType, entry[3]); err != nil {
			return err
		}
	}
	if _, err := f.WriteString("}\n\n"); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(f, "var %v = `", capitalizeFirstLetter(structName)+"_Template"); err != nil {
		return err
	}
	if _, err := f.WriteString(template + "`"); err != nil {
		return err
	}

	return nil
}

func parseFSM(name string, template string) error {
	entries, err := extractEntries(template)
	if err != nil {
		return fmt.Errorf("failed to extract entries: %w", err)
	}

	platform := getPlatform(name)
	name = strings.TrimPrefix(name, platform+"_")
	if platform == "" {
		return fmt.Errorf("platform not found")
	}

	if err := createPlatformDirectory(platform); err != nil {
		return err
	}

	filePath := fmt.Sprintf("models/%s/%s.go", platform, name)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	structName := toCamelCase(name)
	if err := writeStructToFile(f, platform, structName, entries, template); err != nil {
		return fmt.Errorf("failed to write struct to file: %w", err)
	}

	return nil
}

func GenerateFSMStructs(path string) {

	if path == "" {
		path = "./tmp/ntc_templates/templates"
	}

	items, _ := ioutil.ReadDir(path)

	f, err := os.Create("models/map.go")
	if err != nil {
		logrus.Error(err)
	}

	f.WriteString(`package models

import (
	"reflect"
)` + "\n\n")

	f.WriteString("var TemplateMap = map[reflect.Type]string{\n")

	for _, item := range items {
		if item.Name() == "index" {
			continue
		}

		template, err := os.ReadFile(path + "/" + item.Name())

		if err != nil {
			logrus.Error(err)
		}
		name := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

		err = parseFSM(name, string(template))
		if err != nil {
			logrus.Error(err)
			continue
		}

		platform := getPlatform(name)
		structName := toCamelCase(strings.TrimPrefix(name, platform+"_"))
		_, err = fmt.Fprintf(f, "reflect.TypeOf(%[1]v.%[2]v{}): %[1]v.%[2]v_Template,\n", platform, structName)
		if err != nil {
			logrus.Error(err)
		}

	}

	// dirs, _ := os.ReadDir("models/")
	// for _, dir := range dirs {
	// 	logrus.Info(dir.Name())
	// }
	f.WriteString("}\n")
}

func main() {
	GenerateFSMStructs("")
}
