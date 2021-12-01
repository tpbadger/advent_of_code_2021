package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func sum(slice []int) int {
	t := 0
	for _, i := range slice {
		t += i
	}
	return t
}

func main() {
	file, err := os.Open("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// part 1
	scanner := bufio.NewScanner(file)
	prev:= 0
	numIncreased := 0
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if next > prev {
			numIncreased += 1
		}
		prev = next
	}
	fmt.Printf("Answer to part 1 is %d\n", numIncreased - 1)

	// part 2
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, i)
	}
	var sums []int
	winStart := 0
	winEnd := 2
	for winEnd < len(nums) {
		sums = append(sums, sum(nums[winStart:winEnd + 1]))
		winStart += 1
		winEnd += 1
	}
	numIncreased = 0
	for i, sum := range sums {
		if i == 0 {
			continue // skip first element
		}
		if sum > sums[i - 1] {
			numIncreased ++
		}
	}
	fmt.Printf("Answer to part 2 is %d\n", numIncreased)
}



