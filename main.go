package main

import "fmt"
import "strings"

const Version = "0.1.1"

type runeinfo map[string]int

func main() {
	fmt.Printf("runecounter %s\n", Version)
	rawtext := readText()
	lines := strings.Split(rawtext, "\n")
	for idx, line := range lines {
		runemap := mapRunes(line)
		fmt.Printf("%2d) %s\n", idx, line)
		formattedMap := formatMap(runemap)
		fmt.Printf("%s\n\n", formattedMap)
	}
}

func mapRunes(line string) runeinfo {
	runes := make(map[string]int)
	for _, r := range line {
		s := string(r)
		oldCount, ok := runes[s]
		if !ok {
			oldCount = 0
		}
		newCount := oldCount + 1
		runes[s] = newCount
	}
	return runes
}

func formatMap(info runeinfo) string {
	var allItems []string
	for k, v := range info {
		item := fmt.Sprintf("'%s'=%d", k, v)
		allItems = append(allItems, item)
	}
	textRepr := strings.Join(allItems, ", ")
	return textRepr
}

func readText() string {
	text := `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`
	return text
}
