package answers

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//Seeding the random number generator
	rand.Seed(time.Now().UnixNano())
}

//Answer -- interface for a valid answer
type Answer struct {
	Answer  string
	Correct bool
}

//Answers -- interface for a slice of answers
type Answers struct {
	Answers []*Answer
}

//CorrectAnswer -- Returns the correct answer. Note: there should only be one correct answer
func (as Answers) CorrectAnswer() (*Answer, error) {
	for _, a := range as.Answers {
		if a.Correct {
			return a, nil
		}
	}
	return &Answer{}, fmt.Errorf("No Answer was found")
}

//Shuffle -- In place shuffle of the answers
func (as Answers) Shuffle() error {
	numOfAnswers := len(as.Answers)

	//If there are 1 or less answers, then you cannot shuffle them. Throw and Error
	if numOfAnswers < 2 {
		return fmt.Errorf("Not enough Answers to shuffle")
	}

	for i := range as.Answers {
		swapIndex := rand.Intn(numOfAnswers - 1)
		tempt := as.Answers[i]
		as.Answers[i] = as.Answers[swapIndex]
		as.Answers[swapIndex] = tempt
	}

	return nil
}
