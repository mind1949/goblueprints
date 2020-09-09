package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
	otherWord + "hq",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		text := s.Text()
		if isSpace(text) {
			continue
		}
		fmt.Println(strings.Replace(t, otherWord, text, -1))
	}
}

func isSpace(s string) bool {
	is := true
	for _, r := range s {
		if !unicode.IsSpace(r) {
			is = false
			break
		}
	}
	return is
}
