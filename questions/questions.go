package questions

import (
	"fmt"
	"github.com/crazcalm/term-quiz/answers"
	"math/rand"
	"time"
)

func init() {
	//Seeding the random number generator
	rand.Seed(time.Now().UnixNano())
}

//Question -- the interface for a question
type Question struct {
	Question     string
	Answers      answers.Answers
	Explaination string
}

//CorrectAnswer -- returns the currect answer
func (q Question) CorrectAnswer() (*answers.Answer, error) {
	result, err := q.Answers.CorrectAnswer()
	if err != nil {
		return result, fmt.Errorf("No correct answer was found")
	}
	return result, nil
}

//Questions -- The container that holds the questions
type Questions struct {
	Questions []*Question
	index     int
}

//Shuffle -- In place Shuffle of the questions
func (qs Questions) Shuffle() error {
	numOfQuestions := len(qs.Questions)

	if len(qs.Questions) < 2 {
		return fmt.Errorf("There are not enough question to be shuffled")
	}

	for i := range qs.Questions {
		swapIndex := rand.Intn(numOfQuestions)
		tempt := qs.Questions[i]
		qs.Questions[i] = qs.Questions[swapIndex]
		qs.Questions[swapIndex] = tempt
	}

	return nil
}

//Current -- returns the current question
func (qs Questions) Current() *Question {
	return qs.Questions[qs.index]
}

//NextExist -- Checks to see if there is a next question
func (qs Questions) NextExist() bool {
	return qs.index < len(qs.Questions)-1
}

//Next -- Moves index pointer to the next question
func (qs Questions) Next() (*Question, error) {
	var q *Question

	if !qs.NextExist() {
		return q, fmt.Errorf("There is no next question")
	}

	qs.index++
	return qs.Current(), nil
}

//PreviousExist -- Check to see if there is a previous question
func (qs Questions) PreviousExist() bool {
	return qs.index > 0
}

//Previous -- movies index pointer to the previous question
func (qs Questions) Previous() (*Question, error) {
	var q *Question

	if !qs.PreviousExist() {
		return q, fmt.Errorf("There is no previous question")
	}

	qs.index--
	return qs.Current(), nil
}
