package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var words map[string]bool

func readWordList() {
	words = make(map[string]bool)
	file, err := os.Open("sowpods.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	readWordList()

	cols := make([][]rune, 100)
	cols[0] = append(cols[0], 'h', 'i', 'i', 'n', 'o', 'e', 'i', 'l', 'n', 't')
	cols[1] = append(cols[1], 'e', 'i', 'o', 't', 'a', 'h', 'i', 'k', 's')
	cols[2] = append(cols[2], 'a', 'e', 'h', 'i', 'w', 'o', 's', 't')
	cols[3] = append(cols[3], 'a', 'c', 's', 'b', 'e', 'i', 'n', 't')
	cols[4] = append(cols[4], 'a', 'l', 't', 'e', 'o', 'u')
	cols[5] = append(cols[5], 'n', 'p', 't', 'w', 'y', 'a', 'i', 's', 't')
	cols[6] = append(cols[6], 'a', 'o', 's', 's', 'e', 'h', 'i', 'o', 't')
	cols[7] = append(cols[7], 'b', 'i', 'n', 't', 't', 'd', 'e', 'e', 'u', 'u')
	cols[8] = append(cols[8], 'c', 'l', 'o', 's', 'a', 'd', 'h', 's')
	cols[9] = append(cols[9], 'a', 'e', 'w', 'y', 'e', 'e', 'h', 'p', 't')
	cols[10] = append(cols[10], 'a', 'n', 't', 'c', 'l', 'o', 'r')
	cols[11] = append(cols[11], 'a', 'h', 's', 't', 'e', 'n', 't', 'u')
	cols[12] = append(cols[12], 'a', 'd', 'l', 'l', 'o', 'a', 'e', 'l', 't')
	cols[13] = append(cols[13], 'i', 'l', 'n', 'n', 'v', 'a', 'c', 'l', 'o', 'r')
	cols[14] = append(cols[14], 'a', 'e', 'i', 's', 'c', 'e', 'e', 'f')

	solve(cols[0:4], 4)
	//solve(cols[5:9], 4)
}

func solve(cols [][]rune, wordLen int) {

	fmt.Printf("len is %d\n", wordLen)
	fmt.Printf("cols is %+v\n", cols)
	seen := make(map[string]bool)

	seen = solveWorker(cols, seen, "", wordLen, 0)

	fmt.Printf("Found %d possible words\n", len(seen))

}

func solveWorker(cols [][]rune, seen map[string]bool, partial string, wordLen int, position int) map[string]bool {
	for i := 0; i < len(cols[position]); i++ {
		tmp := partial + strings.ToUpper(string(cols[position][i]))
		if position == wordLen-1 {
			// innermost loop
			if _, ok := seen[tmp]; ok {
				continue
			}
			if _, ok := words[tmp]; ok {
				seen[tmp] = true
				fmt.Printf("word: %s\n", tmp)
			}
		} else {
			seen = solveWorker(cols, seen, tmp, wordLen, position+1)
		}
	}
	return seen
}
