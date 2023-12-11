package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type calibration struct {
	firstDigit  int32
	secondDigit int32
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./p1/input_1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int32(0)
	for scanner.Scan() {
		line := scanner.Text()
		sum += ValueOf(Calibrate(line))
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Calibrate(line string) *calibration {
	c := calibration{}
	// Second part
	line = strings.Replace(line, "one", "o1e", -1)
	line = strings.Replace(line, "two", "t2o", -1)
	line = strings.Replace(line, "three", "t3e", -1)
	line = strings.Replace(line, "four", "f4r", -1)
	line = strings.Replace(line, "five", "f5e", -1)
	line = strings.Replace(line, "six", "s6x", -1)
	line = strings.Replace(line, "seven", "s7n", -1)
	line = strings.Replace(line, "eight", "e8t", -1)
	line = strings.Replace(line, "nine", "n9e", -1)

	firstDigitParsed := false
	secondDigitParsed := false
	for _, elem := range line {
		if elem >= 48 && elem <= 57 {
			digit := elem - 48
			if firstDigitParsed {
				c.secondDigit = digit
				secondDigitParsed = true
			} else if secondDigitParsed {
				c.secondDigit = digit
			} else {
				c.firstDigit = digit
				firstDigitParsed = true
			}
		}
	}

	if !secondDigitParsed {
		c.secondDigit = c.firstDigit
	}
	fmt.Println("Calibration: ", line, c.firstDigit, c.secondDigit)
	return &c
}

func ValueOf(calibration *calibration) int32 {
	return calibration.firstDigit*10 + calibration.secondDigit
}
