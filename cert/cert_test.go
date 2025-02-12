package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-02-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err =%v", err)
	}
	if c == nil {
		t.Errorf("Cert data should be a valid")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is wrong")
	}
}

func TestCourseEmptyValur(t *testing.T) {
	_, err := New("", "Bob", "2018-02-31")
	if err == nil {
		t.Error("Empty course")
	}
}

func TestCourseTooLang(t *testing.T) {
	name := "rehjdbnfj439875ztieshfsajnfèosdnfsouehfosaeiènfasonfsaidfiwpejoèefwe4f68sd4fsdf,msiodhfgdf"
	_, err := New(name, "Bob", "2018-02-31")
	if err == nil {
		t.Errorf("Course is too long")
	}
}
func TestNameEmptyValur(t *testing.T) {
	_, err := New("Golang", "", "2018-02-31")
	if err == nil {
		t.Error("Empty course")
	}
}
func TestNameTooLang(t *testing.T) {
	name := "rehjdbnfj439875ztieshfsajnfèosdnfsouehfosaeiènfasonfsaidfiwpejoèefwe4f68sd4fsdf,msiodhfgdf"
	_, err := New("Golang", name, "2018-02-31")
	if err == nil {
		t.Errorf("Course is too long")
	}
}
