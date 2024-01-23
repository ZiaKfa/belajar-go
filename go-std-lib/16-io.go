package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("a very very very very very very long string that is also very very very very very very long \n and also a new line")
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("halo halo bandung\n")
	writer.WriteString("kota kenang kenangan\n")
	writer.Flush()
}
