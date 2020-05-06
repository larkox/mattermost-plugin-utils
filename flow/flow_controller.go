package flow

import (
	"github.com/larkox/mattermost-plugin-utils/common"
	"github.com/larkox/mattermost-plugin-utils/flow/steps"
)

type FlowController interface {
	Start(userID string) error
	NextStep(userID string, from int, value bool) error
	Cancel(userID string) error
}

type flowController struct {
	common.Poster
	common.Logger
	flow      Flow
	store     FlowStore
	pluginURL string
}

func NewFlowController(p common.Poster, l common.Logger, flow Flow, store FlowStore, pluginURL string) FlowController {
	return &flowController{
		Poster:    p,
		Logger:    l,
		flow:      flow,
		store:     store,
		pluginURL: pluginURL,
	}
}

func (fc *flowController) Start(userID string) error {
	err := fc.setFlowStep(userID, 0)
	if err != nil {
		return err
	}
	return fc.processStep(userID, fc.flow.Step(0), 0)
}

func (fc *flowController) NextStep(userID string, from int, value bool) error {
	step, err := fc.getFlowStep(userID)
	if err != nil {
		return err
	}

	if step != from {
		return nil
	}

	skip := fc.flow.Step(step).ShouldSkip(value)
	step += 1 + skip
	if step >= fc.flow.Length() {
		fc.removeFlowStep(userID)
		fc.flow.FlowDone(userID)
		return nil
	}

	err = fc.setFlowStep(userID, step)
	if err != nil {
		return err
	}

	return fc.processStep(userID, fc.flow.Step(step), step)
}

func (fc *flowController) Cancel(userID string) error {
	stepIndex, err := fc.getFlowStep(userID)
	if err != nil {
		return err
	}

	step := fc.flow.Step(stepIndex)
	if step == nil {
		return nil
	}

	postID, err := fc.store.GetPostID(userID, step.GetPropertyName())
	if err != nil {
		return err
	}

	err = fc.DeletePost(postID)
	if err != nil {
		return err
	}

	return nil
}

func (fc *flowController) setFlowStep(userID string, step int) error {
	return fc.store.SetCurrentStep(userID, step)
}

func (fc *flowController) getFlowStep(userID string) (int, error) {
	return fc.store.GetCurrentStep(userID)
}

func (fc *flowController) removeFlowStep(userID string) error {
	return fc.store.DeleteCurrentStep(userID)
}

func (fc *flowController) processStep(userID string, step steps.Step, i int) error {
	if step == nil {
		fc.Errorf("Step nil")
	}

	if fc.flow == nil {
		fc.Errorf("Flow nil")
	}

	if fc.store == nil {
		fc.Errorf("Store nil")
	}
	postID, err := fc.DMWithAttachments(userID, step.PostSlackAttachment(fc.pluginURL+fc.flow.URL(), i))
	if err != nil {
		return err
	}

	if step.IsEmpty() {
		return fc.NextStep(userID, i, false)
	}

	err = fc.store.SetPostID(userID, step.GetPropertyName(), postID)
	if err != nil {
		return err
	}

	return nil
}
