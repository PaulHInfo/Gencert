package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 30

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

/*
	type Saver interface {
		Save(c Cert)
	}
*/
type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := ValidateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := ParseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certification - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is Presented to ",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date : %v", d.Format("02/01/2006")),
	}
	return cert, nil
}
func ValidateStr(str string, max int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 || len(c) > max {
		return c, fmt.Errorf("Invalide String")
	}
	return c, nil
}

func ValidateCourse(course string) (string, error) {
	c, err := ValidateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, "course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateName(name string) (string, error) {
	_, err := ValidateStr(strings.ToTitle(name), MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(name), nil
}

func ParseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
