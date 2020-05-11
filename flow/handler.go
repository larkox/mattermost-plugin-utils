// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package flow

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larkox/mattermost-plugin-utils/common"
	"github.com/larkox/mattermost-plugin-utils/flow/steps"
	"github.com/mattermost/mattermost-server/v5/model"
)

type fh struct {
	flow  Flow
	store FlowStore
}

func Init(r *mux.Router, flow Flow, store FlowStore) {
	fh := &fh{
		flow:  flow,
		store: store,
	}

	flowRouter := r.PathPrefix("/").Subrouter()
	flowRouter.HandleFunc(flow.URL(), fh.handleFlow).Methods(http.MethodPost)
}

func (fh *fh) handleFlow(w http.ResponseWriter, r *http.Request) {
	mattermostUserID := r.Header.Get("Mattermost-User-ID")
	if mattermostUserID == "" {
		common.SlackAttachmentError(w, "Error: Not authorized")
		return
	}

	request := model.PostActionIntegrationRequestFromJson(r.Body)
	if request == nil {
		common.SlackAttachmentError(w, "Error: invalid request")
		return
	}

	stepNumber, ok := request.Context[steps.ContextStepKey].(int)
	if !ok {
		common.SlackAttachmentError(w, "Error: missing step number")
		return
	}

	step := fh.flow.Step(stepNumber)
	if step == nil {
		common.SlackAttachmentError(w, fmt.Sprintf("Error: There is no step %d.", step))
		return
	}

	property, ok := request.Context[steps.ContextPropertyKey].(string)
	if !ok {
		common.SlackAttachmentError(w, "Error: missing property name")
		return
	}

	value, ok := request.Context[steps.ContextButtonValueKey]
	if !ok {
		common.SlackAttachmentError(w, "Error: missing setting id")
		return
	}

	err := fh.store.SetProperty(mattermostUserID, property, value)
	if err != nil {
		common.SlackAttachmentError(w, "There has been a problem setting the property, err="+err.Error())
		return
	}

	response := model.PostActionIntegrationResponse{}
	post := model.Post{}
	model.ParseSlackAttachment(&post, []*model.SlackAttachment{step.ResponseSlackAttachment(value)})
	response.Update = &post

	w.Header().Set("Content-Type", "application/json")
	w.Write(response.ToJson())

	fh.store.RemovePostID(mattermostUserID, property)
	fh.flow.StepDone(mattermostUserID, stepNumber, value)
}
