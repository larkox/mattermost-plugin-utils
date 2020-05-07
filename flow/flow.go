package flow

import (
	"github.com/larkox/mattermost-plugin-utils/flow/steps"
)

type Flow interface {
	Step(i int) steps.Step
	URL() string
	Length() int
	StepDone(userID string, step int, value bool)
	FlowDone(userID string)
}

type flow struct {
	steps      []steps.Step
	url        string
	controller FlowController
	onFlowDone func(userID string)
}

func NewFlow(stepList []steps.Step, url string, fc FlowController, onFlowDone func(userID string)) Flow {
	f := &flow{
		url:        "/welcome",
		controller: fc,
		onFlowDone: onFlowDone,
	}
	return f
}

func (f *flow) Step(i int) steps.Step {
	if i < 0 {
		return nil
	}
	if i >= len(f.steps) {
		return nil
	}
	return f.steps[i]
}

func (f *flow) URL() string {
	return f.url
}

func (f *flow) Length() int {
	return len(f.steps)
}

func (f *flow) StepDone(userID string, step int, value bool) {
	f.controller.NextStep(userID, step, value)
}

func (f *flow) FlowDone(userID string) {
	if f.onFlowDone != nil {
		f.onFlowDone(userID)
	}
}
