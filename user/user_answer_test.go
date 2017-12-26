package user

import (
	"github.com/crazcalm/term-quiz/answers"
	"github.com/crazcalm/term-quiz/questions"
	"testing"
)

func TestCorrect(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}
	aTrue := &answers.Answer{"true", true}
	aFalse := &answers.Answer{"false", false}
	aStatement := &answers.Answer{"Batman is Awesome", true}
	aUserInput := &answers.Answer{"batman Is awesome", false}

	q1 := &questions.Question{"q1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &questions.Question{"q2", answers.Answers{[]*answers.Answer{aTrue, aFalse}}, "none"}
	q3 := &questions.Question{"q3", answers.Answers{[]*answers.Answer{aStatement}}, "none"}
	q4 := &questions.Question{"q4", answers.Answers{[]*answers.Answer{}}, "none"}
	q5 := &questions.Question{"q5", answers.Answers{[]*answers.Answer{a1, a2, a3}}, "none"}

	tests := []struct {
		User        Answer
		Expected    bool
		ExpectError bool
	}{
		{Answer{q1, a4}, true, false},
		{Answer{q1, a3}, false, false},
		{Answer{q2, aTrue}, true, false},
		{Answer{q2, aFalse}, false, false},
		{Answer{q3, aStatement}, true, false},
		{Answer{q3, aUserInput}, true, false},
		{Answer{q4, a1}, false, true},
		{Answer{q5, a2}, false, true},
	}

	for _, test := range tests {
		result, err := test.User.Correct()

		if test.ExpectError && err == nil {
			t.Errorf("Correct was expecting an error, but did not receive an err")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("Correct was not expecting an error, but it got an err: %s", err.Error())
		}

		//I was expecting an error
		//I got an error
		//Test case passed
		if test.ExpectError && err != nil {
			return
		}

		if test.Expected != result {
			t.Errorf("Expected %t, but got %t", test.Expected, result)
		}
	}
}
