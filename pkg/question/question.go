package question

import (
	"time"
)

type AnswerVisibility int

const (
  Unknown AnswerVisibility = iota
	Server
	Client
)

type Answer struct {
	// TODO
}

type QuestionFormat struct {
	// TODO
}

type QuestionVariable struct {
	// TODO
}

type QuestionInputs struct {
	// TODO
}

type Question struct {
	answer Answer                      // answer to each question part
  answerVisibility AnswerVisibility  // how client can access the answer
  created time.Time                       // timestamp when question was created
  format []QuestionFormat            // multiple choice, fill in the blank, etc
  id string                          // unique identifier
  inputs QuestionInputs                         // TODO - better name? How to use tables, sounds, multiple choices, etc?
  variables []QuestionVariable       // Inputs needed for dynamic questions for Multiple Choices
  lang string                        // spoken/written language question is in
  owner string                       // owner id who created question
  revisions []string                 // list of question IDs that are revisions of this question
  sanitized bool                  // true when detected vulgarness
  tags []string                      // tags to search question with
  text string                        // markup text for the question
  version int                     // question format verison
}