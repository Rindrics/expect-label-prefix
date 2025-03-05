package application

import (
	"testing"

	"github.com/Rindrics/expect-label-prefix/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewApp(t *testing.T) {
	client := &MockGitHubClient{}
	logger := &MockLogger{}
	config := &Config{
		Owner:        "test-owner",
		Repository:   "test-repo",
		AddLabel:     true,
		DefaultLabel: "bug",
		Comment:      "Label added",
	}
	info := domain.EventInfo{
		Number: 42,
	}

	logger.On("Debug", mock.AnythingOfType("string"), mock.Anything).Return()

	app := New(info, client, *config, logger)

	switch cmd := app.Command.(type) {
	case AddLabelsCommand:
		assert.Equal(t, "test-owner", cmd.Params.RepoInfo.Owner)
		assert.Equal(t, "test-repo", cmd.Params.RepoInfo.Repo)
	case PostCommentCommand:
		assert.Equal(t, "test-owner", cmd.Params.RepoInfo.Owner)
		assert.Equal(t, "test-repo", cmd.Params.RepoInfo.Repo)
	default:
		t.Fatalf("Unexpected command type: %T", app.Command)
	}

	logger.AssertExpectations(t)
}

func TestRun(t *testing.T) {
	t.Run("add label and exit", func(t *testing.T) {
		mockedLabeler := new(MockLabeler)
		mockedLabeler.On("AddLabels", mock.Anything).Return(nil)
		mockedLogger := &MockLogger{}
		mockedLogger.On("Info", "start executing command", mock.Anything).Once()
		mockedLogger.On("Info", "command executed successfully", mock.Anything).Once()

		app := App{
			Command: AddLabelsCommand{
				Labeler:   mockedLabeler,
				Params:    AddLabelsParams{},
				OnSuccess: &ExitAction{},
			},
			Logger: mockedLogger,
		}

		app.Run()

		mockedLabeler.AssertExpectations(t)
		mockedLogger.AssertExpectations(t)
	})
}
