# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

**Bugs in code:**
- `computeAverage` sums **indexes** instead of the grade values (ranges give `index, elem`). It then divides by `len(grades)` which would also panic on empty slices. 
- `calculateNumericalGrade` mistakenly uses **exams** for `essay_average` instead of essays. 

**Bug in tests (must fix):**
- The “F” test expects `"F"` for 100/95/91. With the given weights/scale it should be **A** (≈96.9). The test’s expected value is wrong; fix to `"A"`.