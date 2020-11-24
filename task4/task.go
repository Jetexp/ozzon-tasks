package task1

import (
	"bufio"
	"math/big"
	"os"
	"strings"
)

func Task() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	inputs := strings.Split(text, " ")

	numbers := []*big.Int{}
	for _, v := range inputs {
		i := new(big.Int)
		n, _ := i.SetString(v, 10)
		numbers = append(numbers, n)
	}

	sum := big.NewInt(0)
	sum.Add(sum, numbers[0])
	sum.Add(sum, numbers[1])
	return sum.String()
}
