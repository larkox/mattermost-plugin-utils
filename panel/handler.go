package panel

import (
	"net/http"

	"github.com/larkox/mattermost-utils-flow/common"
	"github.com/larkox/mattermost-utils-flow/panel/settings"
	"github.com/mattermost/mattermost-server/v5/model"
)

type handler struct {
	panel Panel
}

func Init(r common.Router, panel Panel) {
	sh := &handler{
		panel: panel,
	}

	panelRouter := r.PathPrefix("/").Subrouter()
	panelRouter.HandleFunc(panel.URL(), sh.handleAction).Methods(http.MethodPost)
}

func (sh *handler) handleAction(w http.ResponseWriter, r *http.Request) {
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

	id, ok := request.Context[settings.ContextIDKey]
	if !ok {
		common.SlackAttachmentError(w, "Error: missing setting id")
		return
	}

	value, ok := request.Context[settings.ContextButtonValueKey]
	if !ok {
		value, ok = request.Context[settings.ContextOptionValueKey]
		if !ok {
			common.SlackAttachmentError(w, "Error: valid key not found")
			return
		}
	}

	idString := id.(string)
	err := sh.panel.Set(mattermostUserID, idString, value)
	if err != nil {
		common.SlackAttachmentError(w, "Error: cannot set the property, "+err.Error())
		return
	}

	response := model.PostActionIntegrationResponse{}
	post, err := sh.panel.ToPost(mattermostUserID)
	if err == nil {
		response.Update = post
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response.ToJson())
}
