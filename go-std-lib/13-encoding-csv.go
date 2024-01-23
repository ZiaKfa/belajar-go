package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "Name,HP,Attack,Defense,Speed\n" +
		"Bulbasaur,45,49,49,45\n" +
		"Ivysaur,60,62,63,60\n" +
		"Venusaur,80,82,83,80"

	reader := csv.NewReader(strings.NewReader(csvString))

	for { //endless loop
		record, err := reader.Read()
		if err == io.EOF { //end of file
			break //break the loop
		}

		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"halo", "halo", "bandung"})
	writer.Write([]string{"ibu", "kota", "periangan"})
	writer.Flush()
}
