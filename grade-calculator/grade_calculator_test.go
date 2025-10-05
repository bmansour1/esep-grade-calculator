package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)
	gradeCalculator.AddGrade("exam 1", 20, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 30, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}


func TestGradeTypeString_MapAccess(t *testing.T) {
    cases := []struct {
        gt   GradeType
        want string
    }{
        {Assignment, "assignment"},
        {Exam, "exam"},
        {Essay, "essay"},
    }
    for _, c := range cases {
        if got := c.gt.String(); got != c.want {
            t.Fatalf("GradeType(%d).String() = %q; want %q", int(c.gt), got, c.want)
        }
    }
}


func TestGetFinalGrade_ReturnsC(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("a1", 70, Assignment)
    gc.AddGrade("e1", 70, Exam)
    gc.AddGrade("s1", 70, Essay)
    if got := gc.GetFinalGrade(); got != "C" {
        t.Fatalf("got %s; want C", got)
    }
}


func TestGetFinalGrade_ReturnsD(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("a1", 60, Assignment)
    gc.AddGrade("e1", 60, Exam)
    gc.AddGrade("s1", 60, Essay)
    if got := gc.GetFinalGrade(); got != "D" {
        t.Fatalf("got %s; want D", got)
    }
}
