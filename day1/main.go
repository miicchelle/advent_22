package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)

	}

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	f.Close()

	var sum int
	var totalSums []int
	for _, l := range text {
		if l == "" {
			if sum > 0 {
				totalSums = append(totalSums, sum)
			}
			sum = 0
			continue
		}
		lInt, err := strconv.Atoi(l)
		if err != nil {
			os.Exit(1)
		}
		sum = sum + lInt
	}

	sort.Ints(totalSums)
	fmt.Println(totalSums)

	fmt.Println("-----------")
	var topThree int
	topThree = totalSums[len(totalSums)-1] + totalSums[len(totalSums)-2] + totalSums[len(totalSums)-3]
	fmt.Println(topThree)
}
