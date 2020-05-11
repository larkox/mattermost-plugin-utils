package flow

import (
	"encoding/json"
	"fmt"

	"github.com/larkox/mattermost-plugin-utils/bot/logger"
	"github.com/larkox/mattermost-plugin-utils/bot/poster"
	"github.com/larkox/mattermost-plugin-utils/flow/steps"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func FreeTextMessageHandler(c *plugin.Context, post *model.Post, api plugin.API, botUserID string, controller FlowController, loggerBot logger.Logger, posterBot poster.Poster) {
	if botUserID == post.UserId {
		return
	}

	ch, appErr := api.GetDirectChannel(botUserID, post.UserId)
	if appErr != nil {
		loggerBot.Errorf("error getting direct channel: %s", appErr.Error())
		return
	}

	if ch.Id != post.ChannelId {
		return
	}

	step, stepIndex, err := controller.GetCurrentStep(post.UserId)
	if err != nil {
		loggerBot.Errorf("error retreiving step: %s", err.Error())
		return
	}

	if stepIndex == 0 {
		return
	}

	if !step.WaitForUserInput() {
		return
	}

	posterBot.DMWithAttachments(post.UserId, freeTextSlackAttachment(controller.GetHandlerURL(), post.Message, step, stepIndex))
}

func freeTextSlackAttachment(flowHandler string, value string, step steps.Step, i int) *model.SlackAttachment {
	stepValue, _ := json.Marshal(i)
	actionConfirm := model.PostAction{
		Name: "Confirm",
		Integration: &model.PostActionIntegration{
			URL: flowHandler,
			Context: map[string]interface{}{
				steps.ContextPropertyKey:    step.GetPropertyName(),
				steps.ContextButtonValueKey: value,
				steps.ContextStepKey:        string(stepValue),
			},
		},
	}

	actionCancel := model.PostAction{
		Name: "Cancel",
		Integration: &model.PostActionIntegration{
			URL: flowHandler,
			Context: map[string]interface{}{
				steps.ContextPropertyKey:    step.GetPropertyName(),
				steps.ContextButtonValueKey: "",
				steps.ContextStepKey:        string(stepValue),
			},
		},
	}

	sa := model.SlackAttachment{
		Title:   "Confirm input",
		Text:    fmt.Sprintf("You have typed `%s`. Is that correct?", value),
		Actions: []*model.PostAction{&actionConfirm, &actionCancel},
	}

	return &sa
}
