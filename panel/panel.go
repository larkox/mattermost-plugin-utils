package panel

import (
	"errors"

	"github.com/larkox/mattermost-plugin-utils/bot/logger"
	"github.com/larkox/mattermost-plugin-utils/bot/poster"
	"github.com/larkox/mattermost-plugin-utils/common"
	"github.com/larkox/mattermost-plugin-utils/panel/settings"
	"github.com/mattermost/mattermost-server/v5/model"
)

type Panel interface {
	Set(userID, settingID string, value interface{}) error
	Print(userID string)
	ToPost(userID string) (*model.Post, error)
	Clear(userID string) error
	URL() string
	GetSettingIDs() []string
}

type panel struct {
	settings       map[string]settings.Setting
	settingKeys    []string
	poster         poster.Poster
	logger         logger.Logger
	store          PanelStore
	settingHandler string
	pluginURL      string
}

func NewSettingsPanel(settingList []settings.Setting, poster poster.Poster, logger logger.Logger, store PanelStore, settingHandler, pluginURL string) Panel {
	settingsMap := make(map[string]settings.Setting)
	settingKeys := []string{}
	for _, s := range settingList {
		settingsMap[s.GetID()] = s
		settingKeys = append(settingKeys, s.GetID())
	}

	panel := &panel{
		settings:       settingsMap,
		settingKeys:    settingKeys,
		poster:         poster,
		logger:         logger,
		store:          store,
		settingHandler: settingHandler,
		pluginURL:      pluginURL,
	}

	for _, s := range settingsMap {
		ftf := s.GetFreetextFetcher()
		if ftf == nil {
			continue
		}

		ftf.UpdateHooks(nil, panel.ftOnFetch, panel.ftOnCancel)
	}

	return panel
}

func (p *panel) Set(userID, settingID string, value interface{}) error {
	s, ok := p.settings[settingID]
	if !ok {
		return errors.New("cannot find setting " + settingID)
	}

	err := s.Set(userID, value)
	if err != nil {
		return err
	}
	return nil
}

func (p *panel) GetSettingIDs() []string {
	return p.settingKeys
}

func (p *panel) URL() string {
	return p.settingHandler
}

func (p *panel) Print(userID string) {
	err := p.cleanPreviousSettingsPosts(userID)
	if err != nil {
		p.logger.Errorf("could not clean previous setting post, " + err.Error())
	}

	sas := []*model.SlackAttachment{}
	for _, key := range p.settingKeys {
		s := p.settings[key]
		sa, loopErr := s.GetSlackAttachments(userID, p.pluginURL+p.settingHandler, p.isSettingDisabled(userID, s))
		if loopErr != nil {
			p.logger.Errorf("error creating the slack attachment, err=" + loopErr.Error())
			continue
		}
		sas = append(sas, sa)
	}
	postID, err := p.poster.DMWithAttachments(userID, sas...)
	if err != nil {
		p.logger.Errorf("error creating the message, err=", err.Error())
		return
	}

	err = p.store.SetPanelPostID(userID, postID)
	if err != nil {
		p.logger.Errorf("could not set the post IDs, err=", err.Error())
	}
}

func (p *panel) ToPost(userID string) (*model.Post, error) {
	post := &model.Post{}

	sas := []*model.SlackAttachment{}
	for _, key := range p.settingKeys {
		s := p.settings[key]
		sa, err := s.GetSlackAttachments(userID, p.pluginURL+p.settingHandler, p.isSettingDisabled(userID, s))
		if err != nil {
			p.logger.Errorf("error creating the slack attachment for setting %s, err=%s", s.GetID(), err.Error())
			continue
		}
		sas = append(sas, sa)
	}

	model.ParseSlackAttachment(post, sas)
	return post, nil
}

func (p *panel) cleanPreviousSettingsPosts(userID string) error {
	postID, err := p.store.GetPanelPostID(userID)
	if err == common.ErrNotFound {
		return nil
	}

	if err != nil {
		return err
	}

	err = p.poster.DeletePost(postID)
	if err != nil {
		p.logger.Errorf("could not delete setting post, %s", err)
	}

	err = p.store.DeletePanelPostID(userID)
	if err != nil {
		return err
	}

	return nil
}

func (p *panel) Clear(userID string) error {
	return p.cleanPreviousSettingsPosts(userID)
}

func (p *panel) isSettingDisabled(userID string, s settings.Setting) bool {
	dependencyID := s.GetDependency()
	if dependencyID == "" {
		return false
	}
	dependency, ok := p.settings[dependencyID]
	if !ok {
		p.logger.Errorf("settings dependency %s not found", dependencyID)
		return false
	}

	value, err := dependency.Get(userID)
	if err != nil {
		p.logger.Errorf("cannot get dependency %s value", dependencyID)
		return false
	}
	return s.IsDisabled(value)
}
