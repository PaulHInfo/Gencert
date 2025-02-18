package cert

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSV(fileName string) ([]*Cert, error) {
	certs := make([]*Cert, 0)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Print("ok 1")
		return certs, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Print("ok 2")
		return certs, err
	}
	for _, rec := range records {
		course := rec[0]
		name := rec[1]
		date := rec[2]
		fmt.Println(course)
		fmt.Println(name)
		fmt.Println(date)

		c, err := New(course, name, date)

		if err != nil {
			fmt.Print("ok 3")
			return certs, err
		}
		certs = append(certs, c)
	}
	return certs, nil
}
