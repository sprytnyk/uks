package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `Usage of uks:

    $ uks "Acme Ltd."
    Checks if the provided company in the register of licensed sponsors.

    $ uks -o json "Acme Ltd."
    Print output in format: json/text (default "text")

`

func main() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
	}
	outputType := flag.String("o", "text", "")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("A company name should be provided.")
		flag.Usage()
		os.Exit(1)
	}
	if *outputType != "json" && *outputType != "text" {
		fmt.Println("Format is either json or text.")
		flag.Usage()
		os.Exit(1)
	}

	name := args[0]
	url := getPdfLink(baseURL.String())
	data := downloadPdf(url)
	r := parsePdf(&data, name)
	c := Check{Found: r, Company: name}

	if *outputType == "json" {
		fmt.Println(c.JSON())
	} else {
		fmt.Println(c.prettyString())
	}
}
