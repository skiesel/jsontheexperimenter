package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func getValue(obj interface{}, fieldName string) interface{} {
	mapObj := obj.(map[string]interface{})
	pieces := strings.Split(fieldName, ".")
	lastPiece := pieces[len(pieces)-1]
	pieces = pieces[:len(pieces)-1]

	for _, piece := range(pieces) {
		mapObj = mapObj[piece].(map[string]interface{})
	}

	return mapObj[lastPiece]
}

func getString(obj interface{}, fieldName string) string {
	return getValue(obj, fieldName).(string)
}

func getFloat(obj interface{}, fieldName string) float64 {
	return getValue(obj, fieldName).(float64)
}

func getInt(obj interface{}, fieldName string) int64 {
	return int64(getFloat(obj, fieldName))
}

func primitiveSlice(inter interface{}, formatString string) []string {
	returnSlice := []string{}
	slice := inter.([]interface{})
	for _, val := range(slice) {
		str := fmt.Sprintf(formatString, val)
		returnSlice = append(returnSlice, str)
	}
	return returnSlice
}

func primitive(inter interface{}, formatString string) string {
	return fmt.Sprintf(formatString, inter)
}

func rangeStruct(r interface{}, formatString string) []string {
	returnSlice := []string{}
	min := r.(map[string]interface{})["Min"].(float64)
	max := r.(map[string]interface{})["Max"].(float64)
	increment := r.(map[string]interface{})["Increment"].(float64)

	for val := min; val <= max; val += increment {
		str := fmt.Sprintf(formatString, val)
		returnSlice = append(returnSlice, str)
	}
	return returnSlice
}

func nestedArgs(inter interface{}, formatString string) []string {
	returnSlice := []string{}

	slice := inter.([]interface{})

	for _, arg := range(slice) {
		value := arg.(map[string]interface{})["Value"]

		slice := []string{ fmt.Sprintf(formatString, value) }

		nestedArgs, ok := arg.(map[string]interface{})["Args"]
		if !ok  {
			returnSlice = append(returnSlice, slice...)
			continue
		}

		nestedArgSlice := parseArgs(nestedArgs)

		crossArgs := []string{}

		for _, arg := range(slice) {
			for _, arg2 := range(nestedArgSlice) {
				crossArgs = append(crossArgs, arg + " " + arg2)
			}
		}

		returnSlice = append(returnSlice, crossArgs...)
	}

	return returnSlice
}

func getArgValueSlice(arg interface{}) []string {
	returnSlice := []string{}

	formatString := arg.(map[string]interface{})["Format"].(string)
	value := arg.(map[string]interface{})["Value"]

	switch typed := value.(type) {
		case string:
			returnSlice = append(returnSlice, primitive(typed, formatString))
		case float64:
			returnSlice = append(returnSlice, primitive(typed, formatString))
		case []interface{}:

			switch typed[0].(type) {
				case float64:
					returnSlice = append(returnSlice, primitiveSlice(typed, formatString)...)
				case string:
					returnSlice = append(returnSlice, primitiveSlice(typed, formatString)...)
				case interface{}:
					returnSlice = append(returnSlice, nestedArgs(typed, formatString)...)
			}
			
		case interface{}:
			returnSlice = append(returnSlice, rangeStruct(typed, formatString)...)
	}

	return returnSlice
}

func parseArgs(argsObject interface{}) []string {
	argsMap := argsObject.(map[string]interface{})
	args := []string{}
	for key, _ := range(argsMap) {
		argSlice := getArgValueSlice(argsMap[key])

		if len(args) == 0 {
			args = argSlice
		} else {
			crossArgs := []string{}
			for _, arg := range(args) {
				for _, arg2 := range(argSlice) {
					crossArgs = append(crossArgs, arg + " " + arg2)
				}
			}
			args = crossArgs
		}
	}

	return args
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no enough arguments")
		return
	}

	reader, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(reader)

	var experiment interface{}
	err = decoder.Decode(&experiment)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := parseArgs(experiment.(map[string]interface{})["Args"])

	outputFileType := getString(experiment, "OutputFile")
	commandFormat := getString(experiment, "CommandFormat.Format")
	commandArgs := experiment.(map[string]interface{})["CommandFormat"].(map[string]interface{})["Args"].([]interface{})

	fmt.Println(getString(experiment, "CommentPrefix") + getString(experiment, "Name"))

	for index, arg := range(args) {
		commandSlice := make([]interface{}, 0)
		for _, commandArg := range(commandArgs) {
			commandArg := commandArg.(string)
			if commandArg == "Args" {
				commandSlice = append(commandSlice, arg)
			} else if commandArg == "OutputFile" {
				if outputFileType == "NUMBERED" {
					commandSlice = append(commandSlice, strconv.FormatInt(int64(index), 10))
				} else {
					panic("Unrecognized OutputFileType : " + outputFileType)
				}
			} else {
				commandSlice = append(commandSlice, getString(experiment, commandArg))
			}
		}

		fmt.Printf(commandFormat, commandSlice...)
		fmt.Println()
	}	
}





