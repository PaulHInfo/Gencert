package main

import (
	"fmt"
	"gencert/cert"
	"gencert/certpdf"
	"os"
)

func main() {
	c, err := cert.New("le plus bo", "Simon Frey", "2025-02-12")
	if err != nil {
		fmt.Printf("Error0")
		os.Exit(1)
	}
	//var saver cert.Saver
	saver, err := certpdf.New("output")
	if err != nil {
		fmt.Println("Error1")
		panic(err)

		os.Exit(1)
	}
	saver.Save(*c)
}
