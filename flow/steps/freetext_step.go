package steps

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"
)

type freetextStep struct {
	Title          string
	Message        string
	PropertyName   string
	ResponseFormat string
}

func NewFreeTextStep(title, message, propertyName, responseFormat string) Step {
	return &freetextStep{
		Title:          title,
		Message:        message,
		PropertyName:   propertyName,
		ResponseFormat: responseFormat,
	}
}

func (s *freetextStep) PostSlackAttachment(flowHandler string, i int) *model.SlackAttachment {
	sa := model.SlackAttachment{
		Title: s.Title,
		Text:  s.Message,
	}

	return &sa
}

func (s *freetextStep) ResponseSlackAttachment(value interface{}) *model.SlackAttachment {
	message := fmt.Sprintf(s.ResponseFormat, value)

	sa := model.SlackAttachment{
		Title:   s.Title,
		Text:    message,
		Actions: []*model.PostAction{},
	}

	return &sa
}

func (s *freetextStep) GetPropertyName() string {
	return s.PropertyName
}

func (s *freetextStep) ShouldSkip(value interface{}) int {
	if value.(string) == "" {
		return -1
	}

	return 0
}

func (s *freetextStep) IsEmpty() bool {
	return false
}

func (s *freetextStep) WaitForUserInput() bool {
	return false
}
