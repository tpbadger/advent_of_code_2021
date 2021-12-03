package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func DayTwo() {
	fmt.Println("day 2")
	file, err := os.Open("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// part 1
	scanner := bufio.NewScanner(file)
	xLoc, yLoc := 0, 0
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		val, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		switch split[0] {
		case "forward":
			xLoc += val
		case "up":
			yLoc -= val
		case "down":
			yLoc += val
		default:
			continue
		}
	}
	fmt.Printf("answer to part 1 is %d\n", xLoc*yLoc)
	// part 2
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
	xLoc, yLoc, aim := 0, 0, 0
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		val, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		switch split[0] {
		case "forward":
			xLoc += val
			yLoc += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		default:
			continue
		}
	}
	fmt.Printf("answer to part 1 is %d\n", xLoc*yLoc)
}
