package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		if countBytes {
			scanner.Split(bufio.ScanBytes)
		} else {
			scanner.Split(bufio.ScanWords)
		}
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytesF := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytesF))
}
