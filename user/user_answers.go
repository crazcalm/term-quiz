package user

//Answers -- Container to hold user answers
type Answers map[string]*Answer

//Total -- Returns the total number of answers
func (as Answers) Total() int {
	return len(as)
}

//TotalCorrect -- Returns the total number of correct answers
func (as Answers) TotalCorrect() (int, error) {
	var total int
	var err error

	for _, a := range as {
		result, err := a.Correct()
		if err != nil {
			return total, err
		}
		if result {
			total++
		}
	}

	return total, err
}
