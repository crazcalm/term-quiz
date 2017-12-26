package questions

import (
	"github.com/crazcalm/term-quiz/answers"
	"strings"
	"testing"
)

func TestCorrectAnswer(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	as := answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}

	q1 := Question{"question", as, "none"}

	tests := []struct {
		Question    Question
		Expected    *answers.Answer
		ExpectError bool
	}{
		{q1, a4, false},
	}

	for _, test := range tests {
		result, err := test.Question.CorrectAnswer()

		if test.ExpectError && err == nil {
			t.Errorf("CorrectAnswer was expecting and error, but no error was received")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("CorrectAnswer did not expect and error, but there was an error")
		}

		//An error was expected.
		//An error was recieved.
		//Test case passed
		if test.ExpectError && err != nil {
			return
		}

		if !strings.EqualFold(test.Expected.Answer, result.Answer) {
			t.Errorf("CorrectAnswer expected %s, but got %s", test.Expected.Answer, result.Answer)
		}
	}
}

func TestShuffle(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	q1 := &Question{"question1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &Question{"question2", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q3 := &Question{"question3", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q4 := &Question{"question4", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	tests := []struct {
		Questions   Questions
		ExpectError bool
	}{
		{Questions{[]*Question{q1, q2, q3, q4}, 0}, false},
	}

	for _, test := range tests {
		var firstQuestion string

		if len(test.Questions.Questions) > 1 {
			firstQuestion = test.Questions.Questions[0].Question
		}

		err := test.Questions.Shuffle()

		if test.ExpectError && err == nil {
			t.Errorf("Shuffle was expecting an error, but no error was received")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("Unexpectedly recieved this error: %s", err.Error())
		}

		//I expected and error
		//I got an error
		//Test case passed
		if !test.ExpectError && err != nil {
			return
		}

		shuffled := false
		for i := 0; i < 10; i++ {
			test.Questions.Shuffle()
			if !strings.EqualFold(firstQuestion, test.Questions.Questions[0].Question) {
				shuffled = true
			}
		}

		if !shuffled {
			t.Errorf("Questions were not shuffled")
		}
	}
}

func TestNextExist(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	q1 := &Question{"question1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &Question{"question2", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q3 := &Question{"question3", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q4 := &Question{"question4", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	tests := []struct {
		Questions Questions
		Expected  bool
	}{
		{Questions{[]*Question{q1, q2, q3, q4}, 0}, true},
		{Questions{[]*Question{q1, q2, q3, q4}, 1}, true},
		{Questions{[]*Question{q1, q2, q3, q4}, 2}, true},
		{Questions{[]*Question{q1, q2, q3, q4}, 3}, false},
	}

	for _, test := range tests {
		result := test.Questions.NextExist()
		if result != test.Expected {
			t.Errorf("Expected %t, but got %t", test.Expected, result)
		}
	}
}

func TestNext(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	q1 := &Question{"question1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &Question{"question2", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q3 := &Question{"question3", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q4 := &Question{"question4", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	tests := []struct {
		Questions   Questions
		Expected    *Question
		ExpectError bool
	}{
		{Questions{[]*Question{q1, q2, q3, q4}, 0}, q2, false},
		{Questions{[]*Question{q1, q2, q3, q4}, 1}, q3, false},
		{Questions{[]*Question{q1, q2, q3, q4}, 2}, q4, false},
		{Questions{[]*Question{q1, q2, q3, q4}, 3}, q1, true},
	}

	for _, test := range tests {
		result, err := test.Questions.Next()

		if test.ExpectError && err == nil {
			t.Errorf("Questions.Next expected and error, but did not receive an error")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("Questions.Next received an expected error: %s", err.Error())
		}

		//I expected and error
		//I got an error
		//Test case passed
		if test.ExpectError && err != nil {
			return
		}

		if !strings.EqualFold(test.Expected.Question, result.Question) {
			t.Errorf("Questions.Next expected %s, but got %s", test.Expected.Question, result.Question)
		}
	}
}

func TestPreviousExist(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	q1 := &Question{"question1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &Question{"question2", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q3 := &Question{"question3", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q4 := &Question{"question4", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	tests := []struct {
		Questions Questions
		Expected  bool
	}{
		{Questions{[]*Question{q1, q2, q3, q4}, 0}, false},
		{Questions{[]*Question{q1, q2, q3, q4}, 1}, true},
		{Questions{[]*Question{q1, q2, q3, q4}, 2}, true},
		{Questions{[]*Question{q1, q2, q3, q4}, 3}, true},
	}

	for _, test := range tests {
		result := test.Questions.PreviousExist()
		if result != test.Expected {
			t.Errorf("Expected %t, but got %t", test.Expected, result)
		}
	}
}

func TestPrevious(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	q1 := &Question{"question1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &Question{"question2", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q3 := &Question{"question3", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q4 := &Question{"question4", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	tests := []struct {
		Questions   Questions
		Expected    *Question
		ExpectError bool
	}{
		{Questions{[]*Question{q1, q2, q3, q4}, 0}, q2, true},
		{Questions{[]*Question{q1, q2, q3, q4}, 1}, q1, false},
		{Questions{[]*Question{q1, q2, q3, q4}, 2}, q2, false},
		{Questions{[]*Question{q1, q2, q3, q4}, 3}, q3, false},
	}

	for _, test := range tests {
		result, err := test.Questions.Previous()

		if test.ExpectError && err == nil {
			t.Errorf("Questions.Previous expected and error, but did not receive an error")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("Questions.Previous received an expected error: %s", err.Error())
		}

		//I expected and error
		//I got an error
		//Test case passed
		if test.ExpectError && err != nil {
			return
		}

		if !strings.EqualFold(test.Expected.Question, result.Question) {
			t.Errorf("Questions.Previous expected %s, but got %s", test.Expected.Question, result.Question)
		}
	}
}

func TestCurrent(t *testing.T) {
	a1 := &answers.Answer{"1", false}
	a2 := &answers.Answer{"2", false}
	a3 := &answers.Answer{"3", false}
	a4 := &answers.Answer{"4", true}

	q1 := &Question{"question1", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q2 := &Question{"question2", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q3 := &Question{"question3", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}
	q4 := &Question{"question4", answers.Answers{[]*answers.Answer{a1, a2, a3, a4}}, "none"}

	tests := []struct {
		Questions Questions
		Expected  *Question
	}{
		{Questions{[]*Question{q1, q2, q3, q4}, 0}, q1},
		{Questions{[]*Question{q1, q2, q3, q4}, 1}, q2},
		{Questions{[]*Question{q1, q2, q3, q4}, 2}, q3},
		{Questions{[]*Question{q1, q2, q3, q4}, 3}, q4},
	}

	for _, test := range tests {
		result := test.Questions.Current()
		if !strings.EqualFold(test.Expected.Question, result.Question) {
			t.Errorf("I expected %s, but got %s", test.Expected.Question, result.Question)
		}
	}
}
