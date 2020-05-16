package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
    result := int32(math.MinInt32)

    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            total := int32(0)

            first := arr[i]
            second := arr[i+1]
            third := arr[i+2]

            first_slice := first[j : j+3]
            second_slice := third[j : j+3]
            second_n := second[j+1]
            fmt.Println(first_slice, second_n, second_slice)

            for k, n := range first_slice {
                total += n
                total += second_slice[k]
            }

            total += second_n
            fmt.Println(total)

            if total > result {
                result = total
            }
        }
    }

    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

    fmt.Fprintf(writer, "%d\n", result)

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
