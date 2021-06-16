package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// Downloads a PDF and returns it as a io reader.
func downloadPdf(url string) io.Reader {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(errors.New(response.Status))
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(data.Bytes())
}
