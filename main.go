package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`([0-9]{4})-([0-9]{2})-([0-9]{2}) ([0-9]{2}):([0-9]{2}):([0-9]{2}) \+0000 UTC`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		replaced := replace(line)
		fmt.Println(replaced)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
}

func replace(line string) string {
	matches := re.FindAllStringSubmatch(line, -1)

	replaced := line

	for _, match := range matches {
		time := makeTime(match)
		replaced = strings.ReplaceAll(replaced, match[0], time)
	}

	return replaced
}

func makeTime(match []string) string {
	y := mustParse(match[1])
	m := mustParse(match[2])
	d := mustParse(match[3])
	h := mustParse(match[4])
	i := mustParse(match[5])
	s := mustParse(match[6])

	if y == 1 && m == 1 && d == 1 && h == 0 && i == 0 && s == 0 {
		return "time.Time{}"
	}

	return fmt.Sprintf("time.Date(%d, %d, %d, %d, %d, %d, 0, time.UTC)", y, m, d, h, i, s)
}

func mustParse(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse integer: %v", err)
	}

	return int(i)
}
