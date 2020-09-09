package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	duplicatedVowel bool = true
	removeVowel     bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := []byte(s.Text())
		if randBool() {
			vI := -1
			for i, char := range text {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						vI = i
					}
				}
			}

			if vI > -1 {
				switch randBool() {
				case duplicatedVowel:
					text = append(text[:vI+1], text[vI:]...)
				case removeVowel:
					if vI == len(text)-1 {
						text = text[:len(text)-1]
					} else {
						text = append(text[:vI], text[vI+1:]...)
					}
				}
			}
		}

		fmt.Println(string(text))
	}
}
