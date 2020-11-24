package task1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Task() bool {
	hash := map[int]int{}
	lines := []string{}
	target := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(lines[0])

	if err != nil {
		log.Fatal(err)
	}

	target = n

	numbers := strings.Split(lines[1], " ")

	for _, v := range numbers {
		n, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		if _, ok := hash[n]; !ok {
			hash[n] = 1
			continue
		}

		hash[n] += 1
	}

	for k, _ := range hash {
		need := target - k
		if v, ok := hash[need]; ok {
			if need == k {
				// target = 10
				// numbers = [1 7 4 5 4 10]
				// 10 - 5 = 5, 5 is exist. ONE 5, needs two five
				if v <= 1 {
					return false
				}
			}

			return true
		}
	}

	return false
}
