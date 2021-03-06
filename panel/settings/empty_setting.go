package settings

import (
	"fmt"

	"github.com/larkox/mattermost-plugin-utils/freetext_fetcher"
	"github.com/mattermost/mattermost-server/v5/model"
)

type emptySetting struct {
	title       string
	description string
	id          string
}

func NewEmptySetting(id, title, description string) Setting {
	return &emptySetting{
		id:          id,
		title:       title,
		description: description,
	}
}

func (s *emptySetting) Set(userID string, value interface{}) error {
	return nil
}
func (s *emptySetting) Get(userID string) (interface{}, error) {
	return "", nil
}
func (s *emptySetting) GetID() string {
	return s.id
}
func (s *emptySetting) GetDependency() string {
	return ""
}
func (s *emptySetting) IsDisabled(foreignValue interface{}) bool {
	return false
}
func (s *emptySetting) GetTitle() string {
	return s.title
}
func (s *emptySetting) GetDescription() string {
	return s.description
}
func (s *emptySetting) GetSlackAttachments(userID, settingHandler string, disabled bool) (*model.SlackAttachment, error) {
	title := fmt.Sprintf("Setting: %s", s.title)
	sa := model.SlackAttachment{
		Title: title,
		Text:  s.description,
	}

	return &sa, nil
}

func (s *emptySetting) GetFreetextFetcher() freetext_fetcher.FreetextFetcher {
	return nil
}
