package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// content := getFileContent("sample.txt")
	content := getFileContent("input.txt")
	calibrationValues := processContent(content)
	result := sumCalibrationValues(calibrationValues)

	fmt.Println(result)
}

func sumCalibrationValues(calibrationValues []int) int {
	sumOfCalibrationValues := 0
	for _, i := range calibrationValues {
		sumOfCalibrationValues = sumOfCalibrationValues + i
	}
	return sumOfCalibrationValues
}

func processContent(content []byte) []int {
	calibrationValues := []int{}

	for _, encodedValue := range strings.Split(string(content[:]), "\n") {
		if encodedValue == "" {
			break
		}
		allValues := extractCalibrationValues(encodedValue)
		calibrationValue := allValues[0] + allValues[len(allValues)-1]
		calibrationInteger, err := strconv.Atoi(calibrationValue)
		if err != nil {
			fmt.Println(err)
		}
		calibrationValues = append(calibrationValues, calibrationInteger)
	}
	return calibrationValues
}

func extractCalibrationValues(s string) []string {
	digits := []string{}
	for _, char := range s {
		if unicode.IsNumber(char) {
			digits = append(digits, string(char))
		}
	}
	return digits
}

func getFileContent(fileName string) []byte {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("error occurred reading file")
	}
	return content
}
