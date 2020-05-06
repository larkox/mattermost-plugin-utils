package steps

import "github.com/mattermost/mattermost-server/v5/model"

type Step interface {
	PostSlackAttachment(flowHandler string, i int) *model.SlackAttachment
	ResponseSlackAttachment(value bool) *model.SlackAttachment
	GetPropertyName() string
	ShouldSkip(value bool) int
	IsEmpty() bool
}
