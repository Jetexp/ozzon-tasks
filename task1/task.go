package task1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Task() {
	hash := map[string]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if _, ok := hash[scanner.Text()]; !ok {
			hash[scanner.Text()] = 1
			continue
		}

		hash[scanner.Text()] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k, v := range hash {
		if v == 1 {
			fmt.Println(k)
		}
	}
}
