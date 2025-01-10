package structPractice

/*

Guys please take it as your assignment and please complete it by tomorrow and will review it..

Create a Student struct with fields Name (string) and Marks (slice of integers).
Implement the following receiver methods:

AddMark(mark int) - Adds a new mark to the Marks slice.
CalculateAverage() float64 - Returns the average of the marks.

*/

type Student struct {
	Name  string
	Marks []int
}

func (s *Student) AddMark(mark int) {
	s.Marks = append(s.Marks, mark)
}

func (s Student) CalculateAverage() float64 {
	var v float64
	i := 0
	for _, val := range s.Marks {
		i += val
	}
	v = float64(float64(i) / float64(len(s.Marks)))
	return v
}
