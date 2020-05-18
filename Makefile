GO ?= $(shell command -v go 2> /dev/null)
GO_TEST_FLAGS ?= -race

## Generates mock golang interfaces for testing
mock:
	go install github.com/golang/mock/mockgen
	mockgen -destination panel/mocks/mock_panel.go -package mock_panel github.com/larkox/mattermost-plugin-utils/panel Panel
	mockgen -destination panel/mocks/mock_panelStore.go -package mock_panel github.com/larkox/mattermost-plugin-utils/panel PanelStore
	mockgen -destination panel/mocks/mock_setting.go -package mock_panel github.com/larkox/mattermost-plugin-utils/panel/settings Setting
	mockgen -destination flow/mocks/mock_flow.go -package mock_flow github.com/larkox/mattermost-plugin-utils/flow Flow
	mockgen -destination flow/mocks/mock_controller.go -package mock_flow github.com/larkox/mattermost-plugin-utils/flow FlowController
	mockgen -destination flow/mocks/mock_store.go -package mock_flow github.com/larkox/mattermost-plugin-utils/flow FlowStore
	mockgen -destination flow/mocks/mock_step.go -package mock_flow github.com/larkox/mattermost-plugin-utils/flow/steps Step
	mockgen -destination bot/mocks/mock_bot.go -package mock_bot github.com/larkox/mattermost-plugin-utils/bot Bot
	mockgen -destination bot/mocks/mock_admin.go -package mock_bot github.com/larkox/mattermost-plugin-utils/bot Admin
	mockgen -destination bot/mocks/mock_logger.go -package mock_bot github.com/larkox/mattermost-plugin-utils/bot/logger Logger
	mockgen -destination bot/mocks/mock_poster.go -package mock_bot github.com/larkox/mattermost-plugin-utils/bot/poster Poster
	mockgen -destination freetext_fetcher/mocks/mock_fetcher.go -package mock_freetext_fetcher github.com/larkox/mattermost-plugin-utils/freetext_fetcher FreetextFetcher
	mockgen -destination freetext_fetcher/mocks/mock_manager.go -package mock_freetext_fetcher github.com/larkox/mattermost-plugin-utils/freetext_fetcher Manager
	mockgen -destination freetext_fetcher/mocks/mock_store.go -package mock_freetext_fetcher github.com/larkox/mattermost-plugin-utils/freetext_fetcher FreetextStore


test:
	$(GO) test -race -v ./...