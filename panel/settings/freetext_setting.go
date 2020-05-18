package settings

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/larkox/mattermost-plugin-utils/bot/poster"
	"github.com/larkox/mattermost-plugin-utils/freetext_fetcher"
	"github.com/mattermost/mattermost-server/v5/model"
)

type freetextSetting struct {
	title         string
	description   string
	modifyMessage string
	id            string
	dependsOn     string
	pluginURL     string
	store         SettingStore
	ftf           freetext_fetcher.FreetextFetcher
}

type FreetextInfo struct {
	SettingID string
	UserID    string
}

func NewFreetextSetting(id, title, description, modifyMessage, dependsOn string, store SettingStore, baseURL string, pluginURL string, ftfStore freetext_fetcher.FreetextStore, validate func(string) string, r *mux.Router, posterBot poster.Poster) Setting {
	setting := &freetextSetting{
		title:         title,
		description:   description,
		modifyMessage: modifyMessage,
		id:            id,
		dependsOn:     dependsOn,
		store:         store,
		pluginURL:     pluginURL,
	}
	setting.ftf = freetext_fetcher.NewFreetextFetcher(baseURL, ftfStore, validate, nil, nil, r, posterBot)
	return setting
}

func (s *freetextSetting) Set(userID string, value interface{}) error {
	err := s.store.SetSetting(userID, s.id, value)
	if err != nil {
		return err
	}

	return nil
}

func (s *freetextSetting) Get(userID string) (interface{}, error) {
	value, err := s.store.GetSetting(userID, s.id)
	if err != nil {
		return "", err
	}
	stringValue, ok := value.(string)
	if !ok {
		return "", errors.New("current value is not a string")
	}

	return stringValue, nil
}

func (s *freetextSetting) GetID() string {
	return s.id
}

func (s *freetextSetting) GetTitle() string {
	return s.title
}

func (s *freetextSetting) GetDescription() string {
	return s.description
}

func (s *freetextSetting) GetDependency() string {
	return s.dependsOn
}

func (s *freetextSetting) GetSlackAttachments(userID, settingHandler string, disabled bool) (*model.SlackAttachment, error) {
	title := fmt.Sprintf("Setting: %s", s.title)
	currentValueMessage := "Disabled"

	actions := []*model.PostAction{}
	if !disabled {
		currentValue, err := s.Get(userID)
		if err != nil {
			return nil, err
		}

		currentValueMessage = fmt.Sprintf("Current value: %s", currentValue)

		payload, err := json.Marshal(FreetextInfo{
			UserID:    userID,
			SettingID: s.GetID(),
		})
		if err != nil {
			return nil, err
		}

		actionEdit := model.PostAction{
			Name: "Edit",
			Integration: &model.PostActionIntegration{
				URL: s.pluginURL + s.ftf.URL() + "/new",
				Context: map[string]interface{}{
					freetext_fetcher.ContextNewMessage: s.modifyMessage,
					freetext_fetcher.ContextNewPayload: string(payload),
				},
			},
		}
		actions = []*model.PostAction{&actionEdit}
	}

	text := fmt.Sprintf("%s\n%s", s.description, currentValueMessage)
	sa := model.SlackAttachment{
		Title:   title,
		Text:    text,
		Actions: actions,
	}

	return &sa, nil
}

func (s *freetextSetting) IsDisabled(foreignValue interface{}) bool {
	return foreignValue == "false"
}

func (s *freetextSetting) GetFreetextFetcher() freetext_fetcher.FreetextFetcher {
	return s.ftf
}
