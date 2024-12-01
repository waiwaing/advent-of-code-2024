package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1() {
	file, err := os.Open("inputs/1.txt")
	check(err)
	defer file.Close()

	col1 := make([]int, 0)
	col2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elems := strings.Split(line, "   ")

		if len(elems) != 2 {
			continue
		}

		c1, err := strconv.Atoi(elems[0])
		check(err)
		col1 = append(col1, c1)

		c2, err := strconv.Atoi(elems[1])
		check(err)
		col2 = append(col2, c2)
	}

	sort.Ints(col1)
	sort.Ints(col2)

	col2dict := make(map[int]int)

	for _, e2 := range col2 {
		col2dict[e2] = col2dict[e2] + 1
	}

	result := 0
	result2 := 0

	for i, e1 := range col1 {
		e2 := col2[i]
		tmp := e1 - e2
		if tmp < 0 {
			tmp = -tmp
		}
		result += tmp
		result2 += e1 * col2dict[e1]
	}

	fmt.Println(result)
	fmt.Println(result2)
}
