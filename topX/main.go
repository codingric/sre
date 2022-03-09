package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type Totals struct {
	Indexes map[string]int
	Numbers Numbers
}

type Numbers []Number

type Number struct {
	Value string
	Count int
}

func main() {

	if len(os.Args) != 3 {
		// Incorrect args
		fmt.Printf("Usage: %s TOP FILENAME\n\n  TOP\t\tHow many top totals to display\n  FILENAME\tFile to read numbers from\n", path.Base(os.Args[0]))
		os.Exit(1)
	}
	top := os.Args[1]
	filePath := os.Args[2]
	totals := &Totals{}

	if err := ReadFile(filePath, totals); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(2)
	}

	stop, _ := strconv.Atoi(top)
	sort.Sort(&totals.Numbers)

	for i, n := range totals.Numbers {
		fmt.Printf("#%d: %s (%d)\n", i+1, n.Value, n.Count)
		if i == stop {
			break
		}
	}

}

func ReadFile(filePath string, totals *Totals) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		totals.Add(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// Struct to keep track of numbers in memory
func (t *Totals) Add(number string) {
	// Use lookup map
	index, exist := t.Indexes[number]
	if !exist {
		// new entry
		index = len(t.Indexes)
		if index == 0 {
			//need to inital lookup map
			t.Indexes = make(map[string]int)
		}
		t.Indexes[number] = index //save
		t.Numbers = append(t.Numbers, Number{number, 0})
	}

	t.Numbers[index].Count++
}

// Logic for quicksort
func (a Numbers) Len() int           { return len(a) }
func (a Numbers) Less(i, j int) bool { return a[i].Count > a[j].Count }
func (a Numbers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
