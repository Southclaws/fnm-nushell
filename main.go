package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var re = regexp.MustCompile(`(?m)\$env:(.+) = "(.*)"`)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := json.NewEncoder(os.Stdout)

	table := map[string]interface{}{}

	for in.Scan() {
		line := in.Text()

		matches := re.FindAllStringSubmatch(line, -1)[0]
		key := matches[1]
		value := matches[2]

		if strings.ToLower(key) == "path" {
			pathkey := "Path"
			if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
				pathkey = "PATH"
			}
			table[pathkey] = strings.Split(matches[2], ";")
		} else {
			table[key] = value
		}
	}

	out.Encode(table)
}
