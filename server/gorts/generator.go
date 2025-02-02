package gorts

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func goTypeToTSType(goType reflect.Type) string {
	switch goType.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return "number"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.Slice, reflect.Array:
		return fmt.Sprintf("%s[]", goTypeToTSType(goType.Elem()))
	case reflect.Map:
		return fmt.Sprintf("{ [key: %s]: %s }", goTypeToTSType(goType.Key()), goTypeToTSType(goType.Elem()))
	case reflect.Struct:
		fields := []string{}
		for i := 0; i < goType.NumField(); i++ {
			field := goType.Field(i)
			fields = append(fields, fmt.Sprintf("  %s: %s;", field.Name, goTypeToTSType(field.Type)))
		}
		return fmt.Sprintf("{\n%s\n}", strings.Join(fields, "\n"))
	default:
		return "any"
	}
}

func GenerateTSTypes(service RPCClass) error {
	fileName := fmt.Sprintf("../client/src/gorts/%s.ts", strings.ToLower(service.Name))
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create TypeScript file: %v", err)
	}
	defer file.Close()

	var tsDefinitions []string

	for _, method := range service.Methods {
		argsType := goTypeToTSType(method.ArgsType)
		replyType := goTypeToTSType(method.ReplyType)

		tsDefinitions = append(tsDefinitions, fmt.Sprintf("export type %sArgs = %s;", method.Name, argsType))
		tsDefinitions = append(tsDefinitions, fmt.Sprintf("export type %sReply = %s;", method.Name, replyType))
	}

	serviceInterface := fmt.Sprintf("export interface %s {\n", service.Name)
	for _, method := range service.Methods {
		serviceInterface += fmt.Sprintf("  %s(args: %sArgs): Promise<%sReply>;\n", method.Name, method.Name, method.Name)
	}
	serviceInterface += "}"

	tsDefinitions = append(tsDefinitions, serviceInterface)

	_, err = file.WriteString(strings.Join(tsDefinitions, "\n\n"))
	if err != nil {
		return fmt.Errorf("failed to write TypeScript definitions: %v", err)
	}

	fmt.Printf("Generated TypeScript definitions for %s in %s\n", service.Name, fileName)
	return nil
}
