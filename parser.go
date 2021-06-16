package main

import (
	"io"
	"strings"

	"code.sajari.com/docconv"
)

// Parses a PDF and evaluates if a company name is inside it.
func parsePdf(data *io.Reader, name string) bool {
	output, _, err := docconv.ConvertPDF(*data)
	if err != nil {
		panic(err)
	}

	s := strings.Split(output, "\n")
	for _, v := range s {
		if strings.EqualFold(v, name) {
			return true
		}
	}

	return false
}
