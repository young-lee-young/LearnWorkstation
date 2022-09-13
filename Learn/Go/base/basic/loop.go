package main

import (
	"strconv"
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

/**
	循环
 */
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	printContent(file)
}

func printContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("forever print")
	}
}

func main() {
	result := convertToBin(5)
	fmt.Println(result)

	filename := "/Users/leeyoung/File/Temp/vim_learn.py"
	printFile(filename)

	printContent(strings.NewReader(filename))

	forever()
}
