package answers

import (
	"strings"
	"testing"
)

//TestShuffle --
func TestShuffle(t *testing.T) {
	answerTrue := &Answer{"true", true}
	answerFalse := &Answer{"false", false}
	answerA := &Answer{"a", false}
	answerB := &Answer{"b", false}
	answerC := &Answer{"c", false}
	answerD := &Answer{"d", true}

	answers1 := Answers{[]*Answer{}}
	answers2 := Answers{[]*Answer{answerA}}
	answers3 := Answers{[]*Answer{answerTrue, answerFalse}}
	answers4 := Answers{[]*Answer{answerA, answerB, answerC, answerD}}

	tests := []struct {
		Answers     Answers
		ExpectError bool
	}{
		{answers1, true},
		{answers2, true},
		{answers3, false},
		{answers4, false},
	}

	for _, test := range tests {
		var firstAnswer string

		if len(test.Answers.Answers) != 0 {
			firstAnswer = test.Answers.Answers[0].Answer
		}
		shuffled := false

		for i := 0; i < 10; i++ {
			err := test.Answers.Shuffle()

			if test.ExpectError && err == nil {
				t.Errorf("Shuffle expected an error, but received no error")
				return
			}

			if !test.ExpectError && err != nil {
				t.Errorf("Shuffle received an expected error: %s", err.Error())
				return
			}

			//These cases that do not shuffle have already passed.
			//They will stop here.
			if len(test.Answers.Answers) < 2 {
				return
			}

			if !strings.EqualFold(firstAnswer, test.Answers.Answers[0].Answer) {
				shuffled = true
			}
		}

		if !shuffled {
			t.Errorf("The answers were not shuffled")
		}

	}
}

//TestCorrectAnswer --
func TestCorrectAnswer(t *testing.T) {
	answer1 := &Answer{"1", false}
	answer2 := &Answer{"2", false}
	answer3 := &Answer{"3", false}
	answer4 := &Answer{"4", true}
	answerTrue := &Answer{"真的", true}
	answerFalse := &Answer{"假的", false}

	answers1 := Answers{[]*Answer{answer1, answer2, answer3, answer4}}
	answers2 := Answers{[]*Answer{answer1, answer2, answer3}}
	answers3 := Answers{[]*Answer{answerFalse, answerTrue}}

	tests := []struct {
		Answers     Answers
		Expected    *Answer
		ExpectError bool
	}{
		{answers1, answer4, false},
		{answers2, answer4, true},
		{answers3, answerTrue, false},
	}

	for _, test := range tests {
		answer, err := test.Answers.CorrectAnswer()

		if test.ExpectError && err == nil {
			t.Errorf("CorrectAnswer expected an error but received no error")
		}

		if !test.ExpectError && err != nil {
			t.Errorf("Correctanswer received an error, but was not expecting an error")
		}

		//Test cases that expect an error stop here.
		//They have already passed
		if test.ExpectError {
			return
		}

		if !strings.EqualFold(test.Expected.Answer, answer.Answer) {
			t.Errorf("Expected the answer to be %s, but got %s", test.Expected.Answer, answer.Answer)
		}
	}
}
