package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Complete the matchingStrings function below.
func matchingStrings(strings []string, queries []string) []int32 {
	set := make(map[string]int32)

	for _, q := range strings {
		set[q] += 1
	}

	result := []int32{}

	for _, q := range queries {
		result = append(result, set[q])
	}

	return result
}

func sparse_arrays() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	stringsCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var strings []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := readLine(reader)
		strings = append(strings, stringsItem)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(strings, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}
