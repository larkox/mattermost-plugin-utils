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

type FlowStore interface {
	SetProperty(userID, propertyName string, value bool) error
	SetPostID(userID, propertyName, postID string) error
	GetPostID(userID, propertyName string) (string, error)
	RemovePostID(userID, propertyName string) error
	GetCurrentStep(userID string) (int, error)
	SetCurrentStep(userID string, step int) error
	DeleteCurrentStep(userID string) error
}
