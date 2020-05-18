package flow

import (
	"encoding/json"
)

type freetextInfo struct {
	Step     int
	Property string
	UserID   string
}

func (fc *flowController) ftOnFetch(message string, payload string) {
	var ftInfo freetextInfo
	err := json.Unmarshal([]byte(payload), &ftInfo)
	if err != nil {
		// TODO Handle error better
		return
	}

	err = fc.store.SetProperty(ftInfo.UserID, ftInfo.Property, message)
	if err != nil {
		// TODO Handle error better
		return
	}

	fc.store.RemovePostID(ftInfo.UserID, ftInfo.Property)
	fc.flow.StepDone(ftInfo.UserID, ftInfo.Step, message)
}

func (fc *flowController) ftOnCancel(payload string) {
	var ftInfo freetextInfo
	err := json.Unmarshal([]byte(payload), &ftInfo)
	if err != nil {
		// TODO Handle error better
		return
	}

	fc.store.RemovePostID(ftInfo.UserID, ftInfo.Property)
	fc.flow.StepDone(ftInfo.UserID, ftInfo.Step, "")
}
