package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getMostCommon(bins []string, index int, isOx bool) string {
	bin := ""
	for _, s := range bins {
		bin += string(s[index])
	}
	numZero := 0
	numOne := 0
	for _, c := range bin {
		if string(c) == "0" {
			numZero += 1
		} else {
			numOne += 1
		}
	}
	if isOx {
		if numOne > numZero || numOne == numZero {
			return "1"
		}
		return "0"
	}
	if numOne > numZero || numOne == numZero {
		return "0"
	}
	return "1"
}

func filterSlice(val string, index int, slice []string) []string {
	filtered := []string{}
	for _, str := range slice {
		if string(str[index]) == val {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func DayThree() {
	fmt.Println("day 3")
	file, err := os.Open("./day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	s := []string{}
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	// part 1
	ep, gamma := "", ""
	numOne, numZero := 0, 0
	for i := 0; i < len(s[0]); i++ {
		for _, bin := range s {
			if string(bin[i]) == "0" {
				numZero += 1
			} else {
				numOne += 1
			}
		}
		if numOne >= numZero {
			gamma += "1"
			ep += "0"
		} else {
			gamma += "0"
			ep += "1"
		}
		numOne = 0
		numZero = 0
	}
	gI, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	eI, err := strconv.ParseInt(ep, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer to part 1 is %d\n", gI*eI)
	// part 2
	ox := s
	remainder := len(ox)
	ind := 0
	for remainder > 1 {
		val := getMostCommon(ox, ind, true)
		ox = filterSlice(val, ind, ox)
		ind += 1
		remainder = len(ox)
	}
	oxI, err := strconv.ParseInt(ox[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	di := s
	remainder = len(di)
	ind = 0
	for remainder > 1 {
		val := getMostCommon(di, ind, false)
		di = filterSlice(val, ind, di)
		ind += 1
		remainder = len(di)
	}
	diI, err := strconv.ParseInt(di[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer to part 2 is %d\n", oxI*diI)
}
