package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func dynamicArray(n int32, queries [][]int32) []int32 {
    result := []int32{}
    seqList := make([][]int32, n)
    lastAnswer := int32(0)

    for _, seq := range queries {
        q := seq[0]
        x := seq[1]
        y := seq[2]
        seqN := (x ^ lastAnswer) % n
        sequence := seqList[seqN]

        if q == 1 {
            seqList[seqN] = append(seqList[seqN], y)
        } else {
            lastAnswer = sequence[y%int32(len(sequence))]
            result = append(result, lastAnswer)
        }
    }

    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    q := int32(qTemp)

    var queries [][]int32
    for i := 0; i < int(q); i++ {
        queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != 3 {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := dynamicArray(n, queries)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
