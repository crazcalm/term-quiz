package user

import (
	"github.com/crazcalm/term-quiz/answers"
	"github.com/crazcalm/term-quiz/questions"
	"testing"
)

func TestTotal(t *testing.T) {
	a1 := &answers.Answer{"a1", false}
	a2 := &answers.Answer{"a2", false}
	a3 := &answers.Answer{"a3", false}
	a4 := &answers.Answer{"a4", true}

	q1 := &questions.Question{"q1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	userAnswer := &Answer{q1, a1}

	tests := []struct {
		Answers  Answers
		Expected int
	}{
		{Answers{}, 0},
		{Answers{"1": userAnswer}, 1},
		{Answers{
			"1": userAnswer,
			"2": userAnswer,
			"3": userAnswer,
			"4": userAnswer}, 4},
	}

	for _, test := range tests {
		result := test.Answers.Total()
		if result != test.Expected {
			t.Errorf("Expected %d, but got %d", test.Expected, result)
		}
	}
}

func TestTotalCorrect(t *testing.T) {
	a1 := &answers.Answer{"a1", false}
	a2 := &answers.Answer{"a2", false}
	a3 := &answers.Answer{"a3", false}
	a4 := &answers.Answer{"a4", true}

	q1 := &questions.Question{"q1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	qError := &questions.Question{"qError", answers.Answers{[]*answers.Answer{a1, a2, a3}}, "none"}

	uAnswer1 := &Answer{q1, a1}
	uAnswer2 := &Answer{q1, a2}
	uAnswer3 := &Answer{q1, a3}
	uAnswer4 := &Answer{q1, a4}
	uAnswerError := &Answer{qError, a1}

	tests := []struct {
		Answers     Answers
		Expected    int
		ExpectError bool
	}{
		{Answers{}, 0, false},
		{Answers{
			"1": uAnswer1,
			"2": uAnswer2,
			"3": uAnswer3,
			"4": uAnswer4}, 1, false},
		{Answers{
			"1": uAnswer4,
			"2": uAnswer4,
			"3": uAnswer4,
			"4": uAnswer4}, 4, false},
		{Answers{
			"1": uAnswer4,
			"2": uAnswer4,
			"3": uAnswerError,
			"4": uAnswer4}, 2, true},
	}

	for _, test := range tests {
		result, err := test.Answers.TotalCorrect()

		if test.ExpectError && err == nil {
			t.Errorf("TotalCorrect was expecting an err, but no error was receieved")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("TotalCorrect received an unexpected error: %s", err.Error())
		}

		//I expected an error
		//I receieved an error
		//Test Case passed
		if test.ExpectError && err != nil {
			return
		}

		if result != test.Expected {
			t.Errorf("Expected %d, but got %d", test.Expected, result)
		}
	}
}
