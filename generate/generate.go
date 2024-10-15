package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
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

func makeEntriesMap(entries [][]string) map[string]string {
	mapReturn := make(map[string]string)

	for _, entry := range entries {
		mapKey := entry[3]
		fieldType := "string"
		if entry[2] == "List" {
			fieldType = "[]string"
		}
		mapReturn[mapKey] = fieldType
	}

	return mapReturn
}

func generateStruct(structName string, template string, fields map[string]string) error {
	fset := token.NewFileSet()

	platform := getPlatform(structName)
	structName = strings.TrimPrefix(structName, platform+"_")

	file := &ast.File{
		Name: ast.NewIdent(platform),
	}

	structType := &ast.StructType{
		Fields: &ast.FieldList{},
	}

	for fieldName, fieldType := range fields {
		fieldName = capitalizeFirstLetterLowerRest(fieldName)
		jsonTag := "`json:\"" + fieldName + "\"`"
		field := &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(fieldName)},
			Type:  ast.NewIdent(fieldType),
			Tag:   &ast.BasicLit{Kind: token.STRING, Value: jsonTag},
		}
		structType.Fields.List = append(structType.Fields.List, field)
	}

	typeDecl := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(toCamelCase(structName)),
				Type: structType,
			},
		},
	}

	file.Decls = append(file.Decls, typeDecl)

	varDecl := &ast.GenDecl{
		Tok: token.VAR,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{ast.NewIdent(toCamelCase(structName) + "_Template")},
				Type:  ast.NewIdent("string"),
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: "`" + template + "`", // Use backticks for multiline string
					},
				},
			},
		},
	}

	file.Decls = append(file.Decls, varDecl)

	outFile, err := os.Create("models/" + platform + "/" + structName + ".go")
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := format.Node(outFile, fset, file); err != nil {
		return err
	}

	return nil
}

func parseFSM(name string, template string) error {
	entries, err := extractEntries(template)
	if err != nil {
		return fmt.Errorf("failed to extract entries: %w", err)
	}

	entriesMap := makeEntriesMap(entries)

	err = generateStruct(name, template, entriesMap)
	if err != nil {
		logrus.Error(err)
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

	f.WriteString("}\n")
}

func main() {
	GenerateFSMStructs("")
}
