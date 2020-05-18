package steps

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/larkox/mattermost-plugin-utils/bot/poster"
	"github.com/larkox/mattermost-plugin-utils/freetext_fetcher"
	"github.com/mattermost/mattermost-server/v5/model"
)

type freetextStep struct {
	Title           string
	Message         string
	PropertyName    string
	ResponseFormat  string
	FreetextFetcher freetext_fetcher.FreetextFetcher
}

func NewFreeTextStep(title, message, propertyName, responseFormat string, baseURL string, store freetext_fetcher.FreetextStore, validate func(string) string, r *mux.Router, posterBot poster.Poster) Step {
	return &freetextStep{
		Title:           title,
		Message:         message,
		PropertyName:    propertyName,
		ResponseFormat:  responseFormat,
		FreetextFetcher: freetext_fetcher.NewFreeTextFetcher(baseURL, store, validate, nil, nil, r, posterBot),
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
	if value.(string) == "" {
		message = "Text input cancelled."
	}

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

func (s *freetextStep) GetFreeTextFetcher() freetext_fetcher.FreetextFetcher {
	return s.FreetextFetcher
}
