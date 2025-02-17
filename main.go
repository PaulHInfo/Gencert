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
	outputType := flag.String("type", "pdf", "output type of the certificate")
	flag.Parse()

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
		fmt.Println("")
		os.Exit(1)
	}

	c, err := cert.New("le plus bo", "Simon Frey", "2025-02-12")
	if err != nil {
		fmt.Printf("Error0")
		os.Exit(1)
	}
	saver.Save(*c)
}
