package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`(?m)\$env:(.+) = "(.*)"`)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := json.NewEncoder(os.Stdout)

	table := map[string]string{}

	for in.Scan() {
		line := in.Text()

		matches := re.FindAllStringSubmatch(line, -1)[0]
		key := matches[1]
		value := matches[2]

		table[key] = value
	}

	out.Encode(table)
}
