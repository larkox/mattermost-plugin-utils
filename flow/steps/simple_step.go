package steps

import (
	"encoding/json"

	"github.com/mattermost/mattermost-server/v5/model"
)

type simpleStep struct {
	Title                string
	Message              string
	PropertyName         string
	TrueButtonMessage    string
	FalseButtonMessage   string
	TrueResponseMessage  string
	FalseResponseMessage string
	TrueSkip             int
	FalseSkip            int
}

func NewSimpleStep(title, message, propertyName, trueButtonMessage, falseButtonMessage, trueResponseMessage, falseResponseMessage string, trueSkip, falseSkip int) Step {
	return &simpleStep{
		Title:                title,
		Message:              message,
		PropertyName:         propertyName,
		TrueButtonMessage:    trueButtonMessage,
		FalseButtonMessage:   falseButtonMessage,
		TrueResponseMessage:  trueResponseMessage,
		FalseResponseMessage: falseResponseMessage,
		TrueSkip:             trueSkip,
		FalseSkip:            falseSkip,
	}
}

func (s *simpleStep) PostSlackAttachment(flowHandler string, i int) *model.SlackAttachment {
	trueValue, _ := json.Marshal(true)
	falseValue, _ := json.Marshal(false)
	stepValue, _ := json.Marshal(i)

	actionTrue := model.PostAction{
		Name: s.TrueButtonMessage,
		Integration: &model.PostActionIntegration{
			URL: flowHandler,
			Context: map[string]interface{}{
				ContextPropertyKey:    s.PropertyName,
				ContextButtonValueKey: trueValue,
				ContextStepKey:        stepValue,
			},
		},
	}

	actionFalse := model.PostAction{
		Name: s.FalseButtonMessage,
		Integration: &model.PostActionIntegration{
			URL: flowHandler,
			Context: map[string]interface{}{
				ContextPropertyKey:    s.PropertyName,
				ContextButtonValueKey: falseValue,
				ContextStepKey:        stepValue,
			},
		},
	}

	sa := model.SlackAttachment{
		Title:   s.Title,
		Text:    s.Message,
		Actions: []*model.PostAction{&actionTrue, &actionFalse},
	}

	return &sa
}

func (s *simpleStep) ResponseSlackAttachment(rawValue interface{}) *model.SlackAttachment {
	value, ok := rawValue.(bool)
	message := s.FalseResponseMessage
	if ok && value {
		message = s.TrueResponseMessage
	}

	sa := model.SlackAttachment{
		Title:   s.Title,
		Text:    message,
		Actions: []*model.PostAction{},
	}

	return &sa
}

func (s *simpleStep) GetPropertyName() string {
	return s.PropertyName
}

func (s *simpleStep) ShouldSkip(rawValue interface{}) int {
	value, ok := rawValue.(bool)

	if ok && value {
		return s.TrueSkip
	}

	return s.FalseSkip
}

func (s *simpleStep) IsEmpty() bool {
	return false
}

func (s *simpleStep) WaitForUserInput() bool {
	return false
}
