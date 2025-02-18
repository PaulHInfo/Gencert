package main

import (
	"flag"
	"fmt"
	"gencert/cert"
	"gencert/certhtml"
	"gencert/certpdf"
	"os"
)

func main() {
	file := flag.String("file", "", "CSV file inpute")
	outputType := flag.String("type", "pdf", "output type of the certificate")
	flag.Parse()

	if len(*file) <= 0 {
		fmt.Println("Error - CSv")
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = certhtml.New("Output") // not working ?
	case "pdf":
		saver, err = certpdf.New("output") // not working ?
	default:
		fmt.Println("Unknown output type")

	}
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*file)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err := saver.Save(*c)
		if err != nil {
			fmt.Println("Error1")
		}
	}
}
